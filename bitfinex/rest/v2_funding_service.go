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

func (c *Client) SubmitFundingOffer(ctx context.Context, req *models.SubmitFundingOfferRequest, opts ...RequestOption) (*models.FundingOffer, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodPost,
		Endpoint: "/auth/w/funding/offer/submit",
		Version:  Version2,
		SecType:  SecTypePrivate,
		Params:   req.Params(),
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &models.FundingOffer{}
	err = data.FromRaw(res)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) GetFundingOffersHistory(ctx context.Context, symbol string, opts ...RequestOption) (*models.FundingOffers, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodPost,
		Endpoint: fmt.Sprintf("/auth/r/funding/offers/%s/hist", symbol),
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

func (c *Client) GetFundingInfo(ctx context.Context, symbol string, opts ...RequestOption) (*models.FundingInfo, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodPost,
		Endpoint: fmt.Sprintf("/auth/r/info/funding/%s", symbol),
		Version:  Version2,
		SecType:  SecTypePrivate,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(res))
	data := &models.FundingInfo{}
	err = data.FromRaw(res)
	if err != nil {
		return nil, err
	}
	return data, nil
}
