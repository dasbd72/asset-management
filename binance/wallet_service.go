package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	GetWalletStatusResponse struct {
		Status int    `json:"status"`
		Msg    string `json:"msg"`
	}

	GetUserWalletBalanceResponse []UserWalletBalance

	UserWalletBalance struct {
		Activate   bool        `json:"activate"`
		Balance    JSONFloat64 `json:"balance"` // in BTC
		WalletName string      `json:"walletName"`
	}
)

func (c *Client) GetWalletStatus(ctx context.Context, opts ...RequestOption) (*GetWalletStatusResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/sapi/v1/system/status",
		SecType:  SecTypeNone,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &GetWalletStatusResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) GetUserWalletBalance(ctx context.Context, opts ...RequestOption) (*GetUserWalletBalanceResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/sapi/v1/asset/wallet/balance",
		SecType:  SecTypeSigned,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &GetUserWalletBalanceResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
