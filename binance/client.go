package binance

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Client struct {
	APIKey             string
	APISecret          string
	APIEndpoint        string
	FuturesAPIEndpoint string
	UserAgent          string
	HTTPClient         *http.Client
	Logger             *log.Logger
	TimeOffset         int64
	Debug              bool
}

// Endpoints
const (
	baseAPIMainURL       = "https://api.binance.com"
	baseAPITestnetURL    = "https://testnet.binance.vision"
	futuresAPIMainURL    = "https://fapi.binance.com"
	futuresAPITestnetURL = "https://testnet.binancefuture.com"

	timestampKey  = "timestamp"
	signatureKey  = "signature"
	recvWindowKey = "recvWindow"

	apiKeyHeader = "X-MBX-APIKEY"
)

// UseTestnet switch all the API endpoints from production to the testnet
var UseTestnet = false

// getAPIEndpoint return the base endpoint of the Rest API according the UseTestnet flag
func getAPIEndpoint() string {
	if UseTestnet {
		return baseAPITestnetURL
	}
	return baseAPIMainURL
}

// getFuturesAPIEndpoint return the base endpoint of the Futures Rest API according the UseTestnet flag
func getFuturesAPIEndpoint() string {
	if UseTestnet {
		return futuresAPITestnetURL
	}
	return futuresAPIMainURL
}

// NewClient initialize an API client instance with API key and secret key.
// You should always call this function before using this SDK.
// Services will be created by the form client.NewXXXService().
func NewClient(apiKey, apiSecret string) *Client {
	return &Client{
		APIKey:             apiKey,
		APISecret:          apiSecret,
		APIEndpoint:        getAPIEndpoint(),
		FuturesAPIEndpoint: getFuturesAPIEndpoint(),
		UserAgent:          "golang/binance/v1",
		HTTPClient:         http.DefaultClient,
		Logger:             log.New(os.Stderr, "Binance", log.LstdFlags),
	}
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

func (c *Client) parseRequest(r *Request, opts ...RequestOption) (err error) {
	// set request options from user
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.APIEndpoint, r.Endpoint)
	if strings.HasPrefix(r.Endpoint, "/fapi") {
		fullURL = fmt.Sprintf("%s%s", c.FuturesAPIEndpoint, r.Endpoint)
	}
	if r.RecvWindow > 0 {
		r.SetParam(recvWindowKey, r.RecvWindow)
	}
	if r.SecType == SecTypeSigned {
		r.SetParam(timestampKey, currentTimestamp()-c.TimeOffset)
	}
	queryString := r.Query.Encode()
	body := &bytes.Buffer{}
	bodyString := r.Form.Encode()
	header := http.Header{}
	if r.Header != nil {
		header = r.Header.Clone()
	}
	if bodyString != "" {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		body = bytes.NewBufferString(bodyString)
	}
	if r.SecType == SecTypeAPIKey || r.SecType == SecTypeSigned {
		header.Set(apiKeyHeader, c.APIKey)
	}

	if r.SecType == SecTypeSigned {
		raw := fmt.Sprintf("%s%s", queryString, bodyString)
		mac := hmac.New(sha256.New, []byte(c.APISecret))
		_, err = mac.Write([]byte(raw))
		if err != nil {
			return err
		}
		v := url.Values{}
		v.Set(signatureKey, fmt.Sprintf("%x", (mac.Sum(nil))))
		if queryString == "" {
			queryString = v.Encode()
		} else {
			queryString = fmt.Sprintf("%s&%s", queryString, v.Encode())
		}
	}
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("full url: %s, body: %s", fullURL, bodyString)

	r.FullURL = fullURL
	r.Header = header
	r.Body = body
	return nil
}

func (c *Client) CallAPI(ctx context.Context, r *Request, opts ...RequestOption) ([]byte, error) {
	err := c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(r.Method, r.FullURL, r.Body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.Header
	c.debug("request: %#v", req)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	data, err := io.ReadAll(res.Body)
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
