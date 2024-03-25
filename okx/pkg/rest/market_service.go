package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dasbd72/go-exchange-sdk/okx/pkg/models"
)

// GetTickers get tickers of instruments
func (c *Client) GetTickers(ctx context.Context, req *models.GetTickersRequest, opts ...RequestOption) (*models.GetTickersResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/market/tickers",
		SecType:  SecTypePublic,
		Params:   req.Params(),
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &models.GetTickersResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetTicker get ticker of instrument
func (c *Client) GetTicker(ctx context.Context, req *models.GetTickerRequest, opts ...RequestOption) (*models.GetTickerResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/market/ticker",
		SecType:  SecTypePublic,
		Params:   req.Params(),
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &models.GetTickerResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
