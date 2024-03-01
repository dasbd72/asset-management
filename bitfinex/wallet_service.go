package bitfinex

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	GetWalletsResponse struct {
		Wallet []Wallet `json:"wallet"`
	}
	Wallet struct {
		Type               string      `json:"type"`
		Currency           string      `json:"currency"`
		Balance            JSONFloat64 `json:"balance"`
		UnsettledInterest  JSONFloat64 `json:"unsettled_interest"`
		AvailableBalance   JSONFloat64 `json:"available_balance"`
		LastChange         string      `json:"last_change"`
		LastChangeMetadata interface{} `json:"last_change_metadata"`
	}
)

func (data *GetWalletsResponse) FromRaw(raw []byte) error {
	container := [][]interface{}{}
	err := json.Unmarshal(raw, &container)
	if err != nil {
		return err
	}
	for _, v := range container {
		for i, vv := range v {
			if vv == nil {
				v[i] = ""
			}
		}
		data.Wallet = append(data.Wallet, Wallet{
			Type:               v[0].(string),
			Currency:           v[1].(string),
			Balance:            JSONFloat64(v[2].(float64)),
			UnsettledInterest:  JSONFloat64(v[3].(float64)),
			AvailableBalance:   JSONFloat64(v[4].(float64)),
			LastChange:         v[5].(string),
			LastChangeMetadata: v[6],
		})
	}
	return nil
}

func (c *Client) GetWallets(ctx context.Context, opts ...RequestOption) (*GetWalletsResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodPost,
		Endpoint: "/auth/r/wallets",
		SecType:  SecTypePrivate,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &GetWalletsResponse{}
	err = data.FromRaw(res)
	if err != nil {
		return nil, err
	}
	return data, nil
}
