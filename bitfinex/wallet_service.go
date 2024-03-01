package bitfinex

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	GetWalletsResponse [][]interface{}
)

func (c *Client) GetWalletStatus(ctx context.Context, opts ...RequestOption) (*GetWalletsResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodPost,
		Endpoint: "/auth/r/wallets",
		SecType:  SecTypePrivate,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &GetWalletsResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
