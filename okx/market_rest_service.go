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
		InstType   string `json:"instType"`
		Uly        string `json:"uly,omitempty"`
		InstFamily string `json:"instFamily,omitempty"`
	}

	GetTickersResponse struct {
		BasicResponse
		Tickers []Tickers `json:"data,omitempty"`
	}

	GetTickerRequest struct {
		InstID string `json:"instId"`
	}

	GetTickerResponse struct {
		BasicResponse
		Tickers []Tickers `json:"data,omitempty"`
	}
)

// GetTickers get tickers of instruments
func (c *Client) GetTickers(ctx context.Context, req *GetTickersRequest, opts ...RequestOption) (*GetTickersResponse, error) {
	params := map[string]interface{}{"instType": req.InstType}
	if req.Uly != "" {
		params["uly"] = req.Uly
	}
	if req.InstFamily != "" {
		params["instFamily"] = req.InstFamily
	}
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/market/tickers",
		SecType:  SecTypePublic,
		Params:   params,
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

// GetTicker get ticker of instrument
func (c *Client) GetTicker(ctx context.Context, req *GetTickerRequest, opts ...RequestOption) (*GetTickerResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/market/ticker",
		SecType:  SecTypePublic,
		Params:   map[string]interface{}{"instId": req.InstID},
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
