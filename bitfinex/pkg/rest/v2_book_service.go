package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dasbd72/go-exchange-sdk/bitfinex/pkg/models"
)

func (c *Client) GetTradingPairBooks(ctx context.Context, req *models.GetBooksRequest, opts ...RequestOption) (models.TradingPairBooks, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("/book/%s/%s", req.Symbol, req.Precision),
		Version:  Version2,
		SecType:  SecTypePublic,
		Params:   req.Params(),
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := models.TradingPairBooks{}
	err = data.FromRaw(res)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) GetFundingCurrencyBooks(ctx context.Context, req *models.GetBooksRequest, opts ...RequestOption) (models.FundingCurrencyBooks, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("/book/%s/%s", req.Symbol, req.Precision),
		Version:  Version2,
		SecType:  SecTypePublic,
		Params:   req.Params(),
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := models.FundingCurrencyBooks{}
	err = data.FromRaw(res)
	if err != nil {
		return nil, err
	}
	return data, nil
}
