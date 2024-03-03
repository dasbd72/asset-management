package rest

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Client struct {
	apiKey             string
	apiSecret          string
	publicApiEndpoint  string
	privateApiEndpoint string
	httpClient         *http.Client
}

const (
	basePublicApiURL  = "https://api-pub.bitfinex.com"
	basePrivateApiURL = "https://api.bitfinex.com"

	apiKeyHeader    = "bfx-apikey"
	nonceHeader     = "bfx-nonce"
	signatureHeader = "bfx-signature"
)

// NewClient initialize an API client instance with API key and secret key.
// You should always call this function before using this SDK.
// Services will be created by the form client.NewXXXService().
func NewClient(apiKey, apiSecret string) *Client {
	return &Client{
		apiKey:             apiKey,
		apiSecret:          apiSecret,
		publicApiEndpoint:  basePublicApiURL,
		privateApiEndpoint: basePrivateApiURL,
		httpClient:         http.DefaultClient,
	}
}

func (c *Client) CallAPI(ctx context.Context, r *Request, opts ...RequestOption) (data []byte, err error) {
	// set request options from user
	for _, opt := range opts {
		opt(r)
	}

	req, err := c.getHttpRequest(ctx, r)
	if err != nil {
		return []byte{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		// Only overwrite the retured error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()

	if res.StatusCode >= http.StatusBadRequest {
		apiErr := &APIError{}
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			fmt.Fprintf(os.Stderr, "failed to unmarshal json: %s from %s\n", e, string(data))
		}
		return nil, apiErr
	}
	return data, nil
}

func (c *Client) getHttpRequest(ctx context.Context, r *Request) (*http.Request, error) {
	apiEndpoint := c.publicApiEndpoint
	path := r.endpoint
	query := url.Values{}
	header := http.Header{}
	body := ""
	if r.method == http.MethodGet {
		for k, v := range r.params {
			query.Add(k, fmt.Sprintf("%v", v))
		}
		if len(query) > 0 {
			path += "?" + query.Encode()
		}
	} else {
		b, err := json.Marshal(r.params)
		if err != nil {
			return nil, err
		}
		body = string(b)
		if body == "{}" {
			body = ""
		}
		header.Add("Content-Type", "application/json")
		header.Add("accept", "application/json")
	}
	if r.secType == SecTypePrivate {
		apiEndpoint = c.privateApiEndpoint
		nonce := currentTimestamp()
		header.Set(apiKeyHeader, c.apiKey)
		header.Set(nonceHeader, nonce)
		header.Set(signatureHeader, sign(c.apiSecret, fmt.Sprintf("/api%s%s%s%s", r.version, path, nonce, body)))
	}
	// create request
	req, err := http.NewRequestWithContext(ctx, r.method, fmt.Sprintf("%s%s%s", apiEndpoint, r.version, path), bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, err
	}
	req.Header = header
	return req, nil
}

func sign(secret, message string) string {
	mac := hmac.New(sha512.New384, []byte(secret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}

func currentTimestamp() string {
	return fmt.Sprintf("%d", formatTimestamp(time.Now()))
}

// formatTimestamp formats a time into Unix timestamp in milliseconds, as requested by Binance.
func formatTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
