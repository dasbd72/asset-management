package models

import "github.com/dasbd72/go-exchange-sdk/binance/pkg/cast"

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
		LastUpdateID int64                `json:"lastUpdateId"`
		Bids         [][]cast.JSONFloat64 `json:"bids"`
		Asks         [][]cast.JSONFloat64 `json:"asks"`
	}

	GetAveragePriceRequest struct {
		params map[string]interface{}
	}

	GetAveragePriceResponse struct {
		Mins      int              `json:"mins"`
		Price     cast.JSONFloat64 `json:"price"`
		CloseTime int64            `json:"closeTime"`
	}
)

func NewGetOrderBookRequest(symbol string) *GetOrderBookRequest {
	return &GetOrderBookRequest{
		params: map[string]interface{}{
			"symbol": symbol,
		},
	}
}

func (data *GetOrderBookRequest) Params() map[string]interface{} {
	return data.params
}

func NewGetAveragePriceRequest(symbol string) *GetAveragePriceRequest {
	return &GetAveragePriceRequest{
		params: map[string]interface{}{
			"symbol": symbol,
		},
	}
}

func (data *GetAveragePriceRequest) Params() map[string]interface{} {
	return data.params
}
