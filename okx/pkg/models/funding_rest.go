package models

import "github.com/dasbd72/go-exchange-sdk/okx/pkg/cast"

type (
	GetFundingBalancesRequest struct {
		params map[string]interface{}
	}

	GetFundingBalancesResponse struct {
		BasicResponse
		Balances []struct {
			Ccy       string           `json:"ccy"`
			Bal       cast.JSONFloat64 `json:"bal"`
			FrozenBal cast.JSONFloat64 `json:"frozenBal"`
			AvailBal  cast.JSONFloat64 `json:"availBal"`
		} `json:"data,omitempty"`
	}
)

func NewGetFundingBalancesRequest() *GetFundingBalancesRequest {
	return &GetFundingBalancesRequest{
		params: make(map[string]interface{}),
	}
}

func (r *GetFundingBalancesRequest) Ccy(ccy string) *GetFundingBalancesRequest {
	r.params["ccy"] = ccy
	return r
}

func (r *GetFundingBalancesRequest) Params() map[string]interface{} {
	return r.params
}
