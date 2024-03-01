package binance

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	apiKey             string
	apiSecret          string
	apiEndpoint        string
	futuresAPIEndpoint string
	httpClient         *http.Client
}

// Endpoints
const (
	baseAPIMainURL    = "https://api.binance.com"
	futuresAPIMainURL = "https://fapi.binance.com"

	timestampKey  = "timestamp"
	signatureKey  = "signature"
	recvWindowKey = "recvWindow"

	apiKeyHeader = "X-MBX-APIKEY"
)

// NewClient initialize an API client instance with API key and secret key.
// You should always call this function before using this SDK.
// Services will be created by the form client.NewXXXService().
func NewClient(apiKey, apiSecret string) *Client {
	return &Client{
		apiKey:             apiKey,
		apiSecret:          apiSecret,
		apiEndpoint:        baseAPIMainURL,
		futuresAPIEndpoint: futuresAPIMainURL,
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
			return nil, &APIError{
				Code:    int64(-1),
				Message: string(e.Error()),
			}
		}
		return nil, apiErr
	}
	return data, nil
}

func (c *Client) getHttpRequest(ctx context.Context, r *Request) (*http.Request, error) {
	var (
		u      string
		req    *http.Request
		query  url.Values  = url.Values{}
		header http.Header = http.Header{}
		body   string
	)
	// Build the URL
	if r.apiType == ApiTypeSpot {
		u = fmt.Sprintf("%s%s", c.apiEndpoint, r.endpoint)
	} else {
		u = fmt.Sprintf("%s%s", c.futuresAPIEndpoint, r.endpoint)
	}
	// Set the request parameters
	if r.recvWindow > 0 {
		r.params[recvWindowKey] = r.recvWindow
	}
	if r.secType == SecTypeSigned {
		r.params[timestampKey] = currentTimestamp()
	}
	if r.method == http.MethodGet {
		// Add the parameters to the query string
		for k, v := range r.params {
			query.Add(k, fmt.Sprintf("%v", v))
		}
	} else {
		// Add the parameters to the request body
		form := url.Values{}
		for k, v := range r.params {
			form.Add(k, fmt.Sprintf("%v", v))
		}
		body = form.Encode()
		if body != "" {
			header.Set("Content-Type", "application/x-www-form-urlencoded")
		}
	}
	if r.secType == SecTypeSigned {
		// Add the signature to the query string
		query.Set(signatureKey, sign(c.apiSecret, query.Encode()))
	}
	if r.secType == SecTypeAPIKey || r.secType == SecTypeSigned {
		// Add the API key to the header
		header.Set(apiKeyHeader, c.apiKey)
	}
	if len(query) > 0 {
		// add the query string to the url
		u = fmt.Sprintf("%s?%s", u, query.Encode())
	}
	// create the request
	req, err := http.NewRequestWithContext(ctx, r.method, u, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, err
	}
	req.Header = header
	return req, nil
}

func sign(secret, message string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	return fmt.Sprintf("%x", mac.Sum(nil))
}
