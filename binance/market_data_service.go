package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	GetPingResponse struct {
	}

	GetServerTimeResponse struct {
		ServerTime int64 `json:"serverTime"`
	}

	GetExchangeInfoResponse struct {
		Timezone        string `json:"timezone"`
		ServerTime      int64  `json:"serverTime"`
		RateLimits      []interface{}
		ExchangeFilters []interface{}
		Symbols         []struct {
			Symbol                          string        `json:"symbol"`
			Status                          string        `json:"status"`
			BaseAsset                       string        `json:"baseAsset"`
			BaseAssetPrecision              int           `json:"baseAssetPrecision"`
			QuoteAsset                      string        `json:"quoteAsset"`
			QuotePrecision                  int           `json:"quotePrecision"`
			QuoteAssetPrecision             int           `json:"quoteAssetPrecision"`
			OrderTypes                      []string      `json:"orderTypes"`
			IcebergAllowed                  bool          `json:"icebergAllowed"`
			OcoAllowed                      bool          `json:"ocoAllowed"`
			QuoteOrderQtyMarketAllowed      bool          `json:"quoteOrderQtyMarketAllowed"`
			AllowTrailingStop               bool          `json:"allowTrailingStop"`
			CancelReplaceAllowed            bool          `json:"cancelReplaceAllowed"`
			IsSpotTradingAllowed            bool          `json:"isSpotTradingAllowed"`
			IsMarginTradingAllowed          bool          `json:"isMarginTradingAllowed"`
			Filters                         []interface{} `json:"filters"`
			Permissions                     []string      `json:"permissions"`
			DefaultSelfTradePreventionMode  string        `json:"defaultSelfTradePreventionMode"`
			AllowedSelfTradePreventionModes []string      `json:"allowedSelfTradePreventionModes"`
		}
	}

	GetOrderBookRequest struct {
		params map[string]interface{}
	}

	GetOrderBookResponse struct {
		LastUpdateID int64           `json:"lastUpdateId"`
		Bids         [][]JSONFloat64 `json:"bids"`
		Asks         [][]JSONFloat64 `json:"asks"`
	}

	GetAveragePriceRequest struct {
		params map[string]interface{}
	}

	GetAveragePriceResponse struct {
		Mins      int         `json:"mins"`
		Price     JSONFloat64 `json:"price"`
		CloseTime int64       `json:"closeTime"`
	}
)

func (c *Client) GetPing(ctx context.Context, opts ...RequestOption) (*GetPingResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v3/ping",
		SecType:  SecTypeNone,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &GetPingResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) GetServerTime(ctx context.Context, opts ...RequestOption) (*GetServerTimeResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v3/time",
		SecType:  SecTypeNone,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &GetServerTimeResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) GetExchangeInfo(ctx context.Context, opts ...RequestOption) (*GetExchangeInfoResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v3/exchangeInfo",
		SecType:  SecTypeNone,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &GetExchangeInfoResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewGetOrderBookRequest(symbol string) *GetOrderBookRequest {
	return &GetOrderBookRequest{
		params: map[string]interface{}{
			"symbol": symbol,
		},
	}
}

func (c *Client) GetOrderBook(ctx context.Context, req *GetOrderBookRequest, opts ...RequestOption) (*GetOrderBookResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v3/depth",
		SecType:  SecTypeNone,
		Params:   req.params,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &GetOrderBookResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewGetAveragePriceRequest(symbol string) *GetAveragePriceRequest {
	return &GetAveragePriceRequest{
		params: map[string]interface{}{
			"symbol": symbol,
		},
	}
}

func (c *Client) GetAveragePrice(ctx context.Context, req *GetAveragePriceRequest, opts ...RequestOption) (*GetAveragePriceResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v3/avgPrice",
		SecType:  SecTypeNone,
		Params:   req.params,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &GetAveragePriceResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
