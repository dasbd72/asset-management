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
	"strings"
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
		apiKey:             apiKey,
		apiSecret:          apiSecret,
		apiEndpoint:        getAPIEndpoint(),
		futuresAPIEndpoint: getFuturesAPIEndpoint(),
		httpClient:         http.DefaultClient,
	}
}

func (c *Client) CallAPI(ctx context.Context, r *Request, opts ...RequestOption) (data []byte, err error) {
	// set request options from user
	for _, opt := range opts {
		opt(r)
	}

	var (
		u        string
		req      *http.Request
		query    url.Values = url.Values{}
		body     string
		queryStr string
	)
	// choose the endpoint
	if strings.HasPrefix(r.endpoint, "/fapi") {
		u = fmt.Sprintf("%s%s", c.futuresAPIEndpoint, r.endpoint)
	} else {
		u = fmt.Sprintf("%s%s", c.apiEndpoint, r.endpoint)
	}
	// set the parameters
	if r.recvWindow > 0 {
		r.params[recvWindowKey] = r.recvWindow
	}
	if r.secType == SecTypeSigned {
		r.params[timestampKey] = currentTimestamp()
	}
	if r.method == http.MethodGet {
		for k, v := range r.params {
			query.Add(k, fmt.Sprintf("%v", v))
		}
	} else {
		form := url.Values{}
		for k, v := range r.params {
			form.Add(k, fmt.Sprintf("%v", v))
		}
		body = form.Encode()
	}
	if len(query) > 0 {
		queryStr = query.Encode()
	}
	if r.secType == SecTypeSigned {
		// sign the request
		raw := fmt.Sprintf("%s%s", queryStr, body)
		mac := hmac.New(sha256.New, []byte(c.apiSecret))
		_, err = mac.Write([]byte(raw))
		if err != nil {
			return []byte{}, err
		}
		v := url.Values{}
		v.Set(signatureKey, fmt.Sprintf("%x", (mac.Sum(nil))))
		if queryStr == "" {
			queryStr = v.Encode()
		} else {
			queryStr = fmt.Sprintf("%s&%s", queryStr, v.Encode())
		}
	}
	if queryStr != "" {
		// add the query string to the url
		u = fmt.Sprintf("%s?%s", u, queryStr)
	}
	// create the request
	req, err = http.NewRequest(r.method, u, bytes.NewBuffer([]byte(body)))
	if body != "" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	if r.secType == SecTypeAPIKey || r.secType == SecTypeSigned {
		req.Header.Set(apiKeyHeader, c.apiKey)
	}
	req = req.WithContext(ctx)

	// do the request
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
