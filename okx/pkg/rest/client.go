package rest

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	baseAPIRestURL = "https://www.okx.com"
)

type Client struct {
	apiKey      string
	apiSecret   string
	passphrase  string
	apiEndpoint string
	httpClient  *http.Client
}

const (
	apiKeyHeader     = "OK-ACCESS-KEY"
	signHeader       = "OK-ACCESS-SIGN"
	timestampHeader  = "OK-ACCESS-TIMESTAMP"
	passphraseHeader = "OK-ACCESS-PASSPHRASE"
	simulateHeader   = "x-simulated-trading"
)

// NewClient initialize an API client instance with API key and secret key.
func NewClient(apiKey, apiSecret, passphrase string) *Client {
	return &Client{
		apiKey:      apiKey,
		apiSecret:   apiSecret,
		passphrase:  passphrase,
		apiEndpoint: baseAPIRestURL,
		httpClient:  http.DefaultClient,
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
				Message: fmt.Sprintf("error unmarshalling response %s: %s", data, e.Error()),
			}
		}
		return nil, apiErr
	}
	return data, nil
}

func (c *Client) getHttpRequest(ctx context.Context, r *Request) (*http.Request, error) {
	var (
		path   string      = r.endpoint
		header http.Header = http.Header{}
		body   string
	)
	if r.method == http.MethodGet {
		query := url.Values{}
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
	}
	if r.secType == SecTypePrivate {
		ts := currentTimestamp()
		header.Set(apiKeyHeader, c.apiKey)
		header.Set(passphraseHeader, c.passphrase)
		header.Set(timestampHeader, ts)
		header.Set(signHeader, sign(c.apiSecret, fmt.Sprintf("%s%s%s%s", ts, r.method, path, body)))
	}
	// Create the request
	req, err := http.NewRequestWithContext(ctx, r.method, fmt.Sprintf("%s%s", c.apiEndpoint, path), nil)
	if err != nil {
		return nil, err
	}
	req.Header = header
	return req, nil
}

func sign(secret, message string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func currentTimestamp() string {
	return formatTimestamp(time.Now().UTC())
}

// formatTimestamp formats a time into string.
func formatTimestamp(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.999Z07:00")
}
