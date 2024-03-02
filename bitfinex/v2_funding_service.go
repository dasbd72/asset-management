package bitfinex

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	GetFundingStats struct {
		FundingStatArray []FundingStat `json:"funding_stat_array"`
	}
	FundingStat struct {
		MTS                         JSONInt64   `json:"mts"`
		FRR                         JSONFloat64 `json:"frr"`
		AvgPeriod                   JSONFloat64 `json:"avg_period"`
		FundingAmount               JSONFloat64 `json:"funding_amount"`
		FundingAmountUsed           JSONFloat64 `json:"funding_amount_used"`
		FundingAmountBelowThreshold JSONFloat64 `json:"funding_amount_below_threshold"`
	}

	GetActiveFundingOffers struct {
		FundingOfferArray []FundingOffer `json:"funding_offer_array"`
	}
	FundingOffer struct {
		ID         JSONInt64   `json:"id"`
		Symbol     string      `json:"symbol"`
		MTSCreated JSONInt64   `json:"mts_create"`
		MTSUpdated JSONInt64   `json:"mts_update"`
		Amount     JSONFloat64 `json:"amount"`
		AmountOrig JSONFloat64 `json:"amount_orig"`
		Type       string      `json:"type"`
		Flags      interface{} `json:"flags"`
		Status     string      `json:"status"`
		Rate       JSONFloat64 `json:"rate"`
		Period     JSONInt64   `json:"period"`
		Notify     bool        `json:"notify"`
		Hidden     bool        `json:"hidden"`
		Renew      bool        `json:"renew"`
	}
)

func (data *GetFundingStats) FromRaw(raw []byte) error {
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
		data.FundingStatArray = append(data.FundingStatArray, FundingStat{
			MTS:                         JSONInt64(v[0].(float64)),
			FRR:                         JSONFloat64(v[3].(float64)),
			AvgPeriod:                   JSONFloat64(v[4].(float64)),
			FundingAmount:               JSONFloat64(v[7].(float64)),
			FundingAmountUsed:           JSONFloat64(v[8].(float64)),
			FundingAmountBelowThreshold: JSONFloat64(v[11].(float64)),
		})
	}
	return nil
}

func (c *Client) GetFundingStats(ctx context.Context, symbol string, opts ...RequestOption) (*GetFundingStats, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: fmt.Sprintf("/funding/stats/%s/hist", symbol),
		SecType:  SecTypePublic,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &GetFundingStats{}
	err = data.FromRaw(res)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (data *GetActiveFundingOffers) FromRaw(raw []byte) error {
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
		data.FundingOfferArray = append(data.FundingOfferArray, FundingOffer{
			ID:         JSONInt64(v[0].(float64)),
			Symbol:     v[1].(string),
			MTSCreated: JSONInt64(v[2].(float64)),
			MTSUpdated: JSONInt64(v[3].(float64)),
			Amount:     JSONFloat64(v[4].(float64)),
			AmountOrig: JSONFloat64(v[5].(float64)),
			Type:       v[6].(string),
			Flags:      v[7],
			Status:     v[8].(string),
			Rate:       JSONFloat64(v[9].(float64)),
			Period:     JSONInt64(v[10].(float64)),
			Notify:     v[11].(bool),
			Hidden:     v[12].(bool),
			Renew:      v[13].(bool),
		})
	}
	return nil
}

func (c *Client) GetActiveFundingOffers(ctx context.Context, symbol string, opts ...RequestOption) (*GetActiveFundingOffers, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodPost,
		Endpoint: fmt.Sprintf("/auth/r/funding/offers/%s", symbol),
		SecType:  SecTypePrivate,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(res))
	data := &GetActiveFundingOffers{}
	err = data.FromRaw(res)
	if err != nil {
		return nil, err
	}
	return data, nil
}
