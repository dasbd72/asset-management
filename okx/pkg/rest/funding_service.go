package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dasbd72/go-exchange-sdk/okx/pkg/models"
)

// GetFundingBalances get balances in the funding account
func (c *Client) GetFundingBalances(ctx context.Context, req *models.GetFundingBalancesRequest, opts ...RequestOption) (*models.GetFundingBalancesResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/asset/balances",
		SecType:  SecTypePrivate,
		Params:   req.Params(),
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &models.GetFundingBalancesResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
