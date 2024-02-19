package okx

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	Tickers struct {
		InstID    string      `json:"instId"`
		InstType  string      `json:"instType"`
		Last      JSONFloat64 `json:"last"`
		LastSz    JSONFloat64 `json:"lastSz"`
		AskPx     JSONFloat64 `json:"askPx"`
		AskSz     JSONFloat64 `json:"askSz"`
		BidPx     JSONFloat64 `json:"bidPx"`
		BidSz     JSONFloat64 `json:"bidSz"`
		Open24h   JSONFloat64 `json:"open24h"`
		High24h   JSONFloat64 `json:"high24h"`
		Low24h    JSONFloat64 `json:"low24h"`
		VolCcy24h JSONFloat64 `json:"volCcy24h"`
		Vol24h    JSONFloat64 `json:"vol24h"`
		SodUtc0   JSONFloat64 `json:"sodUtc0"`
		SodUtc8   JSONFloat64 `json:"sodUtc8"`
		TS        JSONTime    `json:"ts"`
	}

	GetTickersRequest struct {
		params map[string]interface{}
	}

	GetTickersResponse struct {
		BasicResponse
		Tickers []Tickers `json:"data,omitempty"`
	}

	GetTickerRequest struct {
		params map[string]interface{}
	}

	GetTickerResponse struct {
		BasicResponse
		Tickers []Tickers `json:"data,omitempty"`
	}
)

func NewGetTickersRequest(instType string) *GetTickersRequest {
	return &GetTickersRequest{
		params: map[string]interface{}{
			"instType": instType,
		},
	}
}

func (r *GetTickersRequest) Uly(uly string) *GetTickersRequest {
	r.params["uly"] = uly
	return r
}

func (r *GetTickersRequest) InstFamily(instFamily string) *GetTickersRequest {
	r.params["instFamily"] = instFamily
	return r
}

// GetTickers get tickers of instruments
func (c *Client) GetTickers(ctx context.Context, req *GetTickersRequest, opts ...RequestOption) (*GetTickersResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/market/tickers",
		SecType:  SecTypePublic,
		Params:   req.params,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &GetTickersResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewGetTickerRequest(instID string) *GetTickerRequest {
	return &GetTickerRequest{
		params: map[string]interface{}{
			"instId": instID,
		},
	}
}

// GetTicker get ticker of instrument
func (c *Client) GetTicker(ctx context.Context, req *GetTickerRequest, opts ...RequestOption) (*GetTickerResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/market/ticker",
		SecType:  SecTypePublic,
		Params:   req.params,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &GetTickerResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
