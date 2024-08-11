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

	GetFundingBillsRequest struct {
		params map[string]interface{}
	}

	GetFundingBillsResponse struct {
		BasicResponse
		Bills []struct {
			BillId   cast.JSONInt64   `json:"billId"`
			Ccy      string           `json:"ccy"`
			ClientId string           `json:"clientId"`
			BalChg   cast.JSONFloat64 `json:"balChg"`
			Bal      cast.JSONFloat64 `json:"bal"`
			Type     cast.JSONInt64   `json:"type"`
			Ts       cast.JSONTime    `json:"ts"`
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

func NewGetFundingBillsRequest() *GetFundingBillsRequest {
	return &GetFundingBillsRequest{
		params: make(map[string]interface{}),
	}
}

func (r *GetFundingBillsRequest) Ccy(ccy string) *GetFundingBillsRequest {
	r.params["ccy"] = ccy
	return r
}

func (r *GetFundingBillsRequest) Type(t string) *GetFundingBillsRequest {
	r.params["type"] = t
	return r
}

func (r *GetFundingBillsRequest) ClientId(clientId string) *GetFundingBillsRequest {
	r.params["clientId"] = clientId
	return r
}

func (r *GetFundingBillsRequest) After(after string) *GetFundingBillsRequest {
	r.params["after"] = after
	return r
}

func (r *GetFundingBillsRequest) Before(before string) *GetFundingBillsRequest {
	r.params["before"] = before
	return r
}

func (r *GetFundingBillsRequest) Limit(limit string) *GetFundingBillsRequest {
	r.params["limit"] = limit
	return r
}

func (r *GetFundingBillsRequest) Params() map[string]interface{} {
	return r.params
}
