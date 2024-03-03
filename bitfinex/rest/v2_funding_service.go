package bitfinexRest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dasbd72/go-exchange-sdk/bitfinex/pkg/models"
)

func (c *Client) GetFundingStats(ctx context.Context, symbol string, opts ...RequestOption) (*models.FundingStats, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("/funding/stats/%s/hist", symbol),
		Version:  Version2,
		SecType:  SecTypePublic,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &models.FundingStats{}
	err = data.FromRaw(res)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) GetActiveFundingOffers(ctx context.Context, symbol string, opts ...RequestOption) (*models.FundingOffers, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodPost,
		Endpoint: fmt.Sprintf("/auth/r/funding/offers/%s", symbol),
		Version:  Version2,
		SecType:  SecTypePrivate,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &models.FundingOffers{}
	err = data.FromRaw(res)
	if err != nil {
		return nil, err
	}
	return data, nil
}
