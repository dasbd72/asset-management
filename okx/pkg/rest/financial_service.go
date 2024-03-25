package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dasbd72/go-exchange-sdk/okx/pkg/models"
)

// GetEarnOffers get the available earn offers
func (c *Client) GetEarnOffers(ctx context.Context, req *models.GetEarnOffersRequest, opts ...RequestOption) (*models.GetEarnOffersResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/finance/staking-defi/offers",
		SecType:  SecTypePrivate,
		Params:   req.Params(),
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &models.GetEarnOffersResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetActiveEarnOrders get the active earn orders
func (c *Client) GetActiveEarnOrders(ctx context.Context, req *models.GetActiveEarnOrdersRequest, opts ...RequestOption) (*models.GetActiveEarnOrdersResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/finance/staking-defi/orders-active",
		SecType:  SecTypePrivate,
		Params:   req.Params(),
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &models.GetActiveEarnOrdersResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetETHStakingBalance get balances in the ETH staking account
func (c *Client) GetETHStakingBalance(ctx context.Context, opts ...RequestOption) (*models.GetETHStakingBalanceResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/finance/staking-defi/eth/balance",
		SecType:  SecTypePrivate,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &models.GetETHStakingBalanceResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetSavingBalance get balances in the saving account
func (c *Client) GetSavingBalance(ctx context.Context, req *models.GetSavingBalanceRequest, opts ...RequestOption) (*models.GetSavingBalanceResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/finance/savings/balance",
		SecType:  SecTypePrivate,
		Params:   req.Params(),
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &models.GetSavingBalanceResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetLendingHistory get the lending history
func (c *Client) GetLendingHistory(ctx context.Context, req *models.GetLendingHistoryRequest, opts ...RequestOption) (*models.GetLendingHistoryResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/finance/savings/lending-history",
		SecType:  SecTypePrivate,
		Params:   req.Params(),
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &models.GetLendingHistoryResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
