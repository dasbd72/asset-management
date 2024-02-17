package okx

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	baseAPIRestURL = "https://www.okx.com"
	baseAPIDemoURL = "https://www.okx.com"
)

type Client struct {
	apiKey      string
	apiSecret   string
	passphrase  string
	apiEndpoint string
	httpClient  *http.Client
	logger      *log.Logger
	debugging   bool
}

const (
	apiKeyHeader     = "OK-ACCESS-KEY"
	signHeader       = "OK-ACCESS-SIGN"
	timestampHeader  = "OK-ACCESS-TIMESTAMP"
	passphraseHeader = "OK-ACCESS-PASSPHRASE"
	simulateHeader   = "x-simulated-trading"
)

// UseDemo switch all the API endpoints from production to the testnet
var UseDemo = false

// getAPIEndpoint return the base endpoint of the Rest API according the UseDemo flag
func getAPIEndpoint() string {
	if UseDemo {
		return baseAPIDemoURL
	}
	return baseAPIRestURL
}

// NewClient initialize an API client instance with API key and secret key.
func NewClient(apiKey, apiSecret, passphrase string) *Client {
	return &Client{
		apiKey:      apiKey,
		apiSecret:   apiSecret,
		passphrase:  passphrase,
		apiEndpoint: getAPIEndpoint(),
		httpClient:  http.DefaultClient,
		logger:      log.New(os.Stderr, "OKX", log.LstdFlags),
		debugging:   false,
	}
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.debugging {
		c.logger.Printf(format, v...)
	}
}

func (c *Client) CallAPI(ctx context.Context, r *Request, opts ...RequestOption) (data []byte, err error) {
	// set request options from user
	for _, opt := range opts {
		opt(r)
	}

	u := fmt.Sprintf("%s%s", c.apiEndpoint, r.endpoint)
	path := r.endpoint
	var (
		req  *http.Request
		body string
	)
	if r.method == http.MethodGet {
		query := url.Values{}
		for k, v := range r.params {
			query.Add(k, fmt.Sprintf("%v", v))
		}
		if len(query) > 0 {
			u += "?" + query.Encode()
			path += "?" + query.Encode()
		}
		req, err = http.NewRequest(r.method, u, nil)
	} else {
		b, err := json.Marshal(r.params)
		if err != nil {
			return nil, err
		}
		body = string(b)
		if body == "{}" {
			body = ""
		}
		req, err = http.NewRequest(r.method, u, bytes.NewBuffer(b))
		if err != nil {
			return []byte{}, err
		}
		req.Header.Add("Content-Type", "application/json")
	}
	c.debug("full url: %s, body: %s", u, body)
	if r.secType == SecTypePrivate {
		ts := currentTimestamp()
		mac := hmac.New(sha256.New, []byte(c.apiSecret))
		_, err = mac.Write([]byte(ts + r.method + path + body))
		if err != nil {
			return nil, err
		}

		req.Header.Set(apiKeyHeader, c.apiKey)
		req.Header.Set(passphraseHeader, c.passphrase)
		req.Header.Set(timestampHeader, ts)
		req.Header.Set(signHeader, base64.StdEncoding.EncodeToString(mac.Sum(nil)))
	}
	if UseDemo {
		req.Header.Set(simulateHeader, "1")
	}
	req = req.WithContext(ctx)
	c.debug("request: %#v", req)
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
	c.debug("response: %#v", res)
	c.debug("response body: %s", string(data))
	c.debug("response status code: %d", res.StatusCode)

	if res.StatusCode >= http.StatusBadRequest {
		apiErr := &APIError{}
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s", e)
		}
		return nil, apiErr
	}
	return data, nil
}
