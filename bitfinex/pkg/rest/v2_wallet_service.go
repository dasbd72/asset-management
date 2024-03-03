package rest

import (
	"context"
	"net/http"

	"github.com/dasbd72/go-exchange-sdk/bitfinex/pkg/models"
)

func (c *Client) GetWallets(ctx context.Context, opts ...RequestOption) (*models.Wallets, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodPost,
		Endpoint: "/auth/r/wallets",
		Version:  Version2,
		SecType:  SecTypePrivate,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &models.Wallets{}
	err = data.FromRaw(res)
	if err != nil {
		return nil, err
	}
	return data, nil
}
