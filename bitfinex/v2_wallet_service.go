package bitfinex

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dasbd72/go-exchange-sdk/bitfinex/pkg/cast"
)

type (
	GetWalletsResponse struct {
		Wallet []Wallet `json:"wallet"`
	}
	Wallet struct {
		Type               cast.NilOrString
		Currency           cast.NilOrString
		Balance            cast.NilOrFloat64
		UnsettledInterest  cast.NilOrFloat64
		AvailableBalance   cast.NilOrFloat64
		LastChange         cast.NilOrString
		LastChangeMetadata interface{}
	}
)

func (data *GetWalletsResponse) FromRaw(raw []byte) error {
	container := [][]interface{}{}
	err := json.Unmarshal(raw, &container)
	if err != nil {
		return err
	}
	for _, v := range container {
		data.Wallet = append(data.Wallet, Wallet{
			Type:               cast.IfToNilOrString(v[0]),
			Currency:           cast.IfToNilOrString(v[1]),
			Balance:            cast.IfToNilOrFloat64(v[2]),
			UnsettledInterest:  cast.IfToNilOrFloat64(v[3]),
			AvailableBalance:   cast.IfToNilOrFloat64(v[4]),
			LastChange:         cast.IfToNilOrString(v[5]),
			LastChangeMetadata: v[6],
		})
	}
	return nil
}

func (c *Client) GetWallets(ctx context.Context, opts ...RequestOption) (*GetWalletsResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodPost,
		Endpoint: "/auth/r/wallets",
		Version:  Version2,
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
