package pionex

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const (
	baseAPIMainURL = "https://api.pionex.com"

	timestampKey = "timestamp"

	apiKeyHeader    = "PIONEX-KEY"
	signatureHeader = "PIONEX-SIGNATURE"
)

type Client struct {
	apiKey      string
	apiSecret   string
	apiEndpoint string
	httpClient  *http.Client
}

type Client_builder struct {
	APIKey      string
	APISecret   string
	APIEndpoint string
	HTTPClient  *http.Client
}

func (b Client_builder) Build() *Client {
	if b.APIKey == "" {
		b.APIKey = os.Getenv("PIONEX_API_KEY")
	}
	if b.APISecret == "" {
		b.APISecret = os.Getenv("PIONEX_API_SECRET")
	}
	if b.APIKey == "" || b.APISecret == "" {
		panic("API key and secret are required")
	}
	if b.APIEndpoint == "" {
		b.APIEndpoint = baseAPIMainURL
	}
	if b.HTTPClient == nil {
		b.HTTPClient = http.DefaultClient
	}
	return &Client{
		apiKey:      b.APIKey,
		apiSecret:   b.APISecret,
		apiEndpoint: b.APIEndpoint,
		httpClient:  b.HTTPClient,
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
			fmt.Fprintf(os.Stderr, "failed to unmarshal json: %s", e)
		}
		return nil, apiErr
	}
	return data, nil
}

func (c *Client) getHttpRequest(ctx context.Context, r *Request) (*http.Request, error) {
	var (
		path   string      = r.endpoint
		query  url.Values  = url.Values{}
		header http.Header = http.Header{}
		body   string
	)
	if r.method == http.MethodGet {
		for k, v := range r.params {
			query.Add(k, fmt.Sprintf("%v", v))
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
	}
	if r.secType == SecTypePrivate {
		query.Add(timestampKey, fmt.Sprintf("%d", currentTimestamp()))
	}
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	if r.secType == SecTypePrivate {
		header.Set(apiKeyHeader, c.apiKey)
		header.Set(signatureHeader, sign(c.apiSecret, r.method+path+body))
	}
	// Create the request
	req, err := http.NewRequestWithContext(ctx, r.method, fmt.Sprintf("%s%s", c.apiEndpoint, path), bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, err
	}
	req.Header = header
	return req, nil
}

func sign(secret, message string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}
