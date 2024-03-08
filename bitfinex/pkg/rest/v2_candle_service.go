package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dasbd72/go-exchange-sdk/bitfinex/pkg/models"
)

func (c *Client) GetCandlesRequest(ctx context.Context, req *models.GetCandlesRequest, opts ...RequestOption) (models.Candles, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("/candles/%s/%s", req.Candle, req.Section),
		Version:  Version2,
		SecType:  SecTypePublic,
		Params:   req.Params(),
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := models.Candles{}
	err = data.FromRaw(res)
	if err != nil {
		return nil, err
	}
	return data, nil
}
