package spot

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dasbd72/go-exchange-sdk/binance/pkg/models"
)

func (c *Client) GetPing(ctx context.Context, opts ...RequestOption) (*models.GetPingResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v3/ping",
		SecType:  SecTypeNone,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &models.GetPingResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) GetServerTime(ctx context.Context, opts ...RequestOption) (*models.GetServerTimeResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v3/time",
		SecType:  SecTypeNone,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &models.GetServerTimeResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) GetExchangeInfo(ctx context.Context, opts ...RequestOption) (*models.GetExchangeInfoResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v3/exchangeInfo",
		SecType:  SecTypeNone,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &models.GetExchangeInfoResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) GetOrderBook(ctx context.Context, req *models.GetOrderBookRequest, opts ...RequestOption) (*models.GetOrderBookResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v3/depth",
		SecType:  SecTypeNone,
		Params:   req.Params(),
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &models.GetOrderBookResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) GetAveragePrice(ctx context.Context, req *models.GetAveragePriceRequest, opts ...RequestOption) (*models.GetAveragePriceResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v3/avgPrice",
		SecType:  SecTypeNone,
		Params:   req.Params(),
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &models.GetAveragePriceResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
