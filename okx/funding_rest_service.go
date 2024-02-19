package okx

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	GetFundingBalancesRequest struct {
		params map[string]interface{}
	}

	GetFundingBalancesResponse struct {
		BasicResponse
		Balances []struct {
			Ccy       string      `json:"ccy"`
			Bal       JSONFloat64 `json:"bal"`
			FrozenBal JSONFloat64 `json:"frozenBal"`
			AvailBal  JSONFloat64 `json:"availBal"`
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

// GetFundingBalances get balances in the funding account
func (c *Client) GetFundingBalances(ctx context.Context, req *GetFundingBalancesRequest, opts ...RequestOption) (*GetFundingBalancesResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/asset/balances",
		SecType:  SecTypePrivate,
		Params:   req.params,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &GetFundingBalancesResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
