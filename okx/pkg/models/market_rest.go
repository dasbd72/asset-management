package models

import "github.com/dasbd72/go-exchange-sdk/okx/pkg/cast"

type (
	Tickers struct {
		InstID    string           `json:"instId"`
		InstType  string           `json:"instType"`
		Last      cast.JSONFloat64 `json:"last"`
		LastSz    cast.JSONFloat64 `json:"lastSz"`
		AskPx     cast.JSONFloat64 `json:"askPx"`
		AskSz     cast.JSONFloat64 `json:"askSz"`
		BidPx     cast.JSONFloat64 `json:"bidPx"`
		BidSz     cast.JSONFloat64 `json:"bidSz"`
		Open24h   cast.JSONFloat64 `json:"open24h"`
		High24h   cast.JSONFloat64 `json:"high24h"`
		Low24h    cast.JSONFloat64 `json:"low24h"`
		VolCcy24h cast.JSONFloat64 `json:"volCcy24h"`
		Vol24h    cast.JSONFloat64 `json:"vol24h"`
		SodUtc0   cast.JSONFloat64 `json:"sodUtc0"`
		SodUtc8   cast.JSONFloat64 `json:"sodUtc8"`
		TS        cast.JSONTime    `json:"ts"`
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

func (r *GetTickersRequest) Params() map[string]interface{} {
	return r.params
}

func NewGetTickerRequest(instID string) *GetTickerRequest {
	return &GetTickerRequest{
		params: map[string]interface{}{
			"instId": instID,
		},
	}
}

func (r *GetTickerRequest) Params() map[string]interface{} {
	return r.params
}
