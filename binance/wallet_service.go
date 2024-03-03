package binance

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dasbd72/go-exchange-sdk/binance/pkg/models"
)

func (c *Client) GetWalletStatus(ctx context.Context, opts ...RequestOption) (*models.GetWalletStatusResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/sapi/v1/system/status",
		SecType:  SecTypeNone,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &models.GetWalletStatusResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) GetUserWalletBalance(ctx context.Context, opts ...RequestOption) (*models.GetUserWalletBalanceResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/sapi/v1/asset/wallet/balance",
		SecType:  SecTypeSigned,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &models.GetUserWalletBalanceResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
