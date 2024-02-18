package okx

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	GetBalanceRequest struct {
		Ccy string `json:"ccy,omitempty"`
	}

	GetBalanceResponse struct {
		BasicResponse
		Balances []*struct {
			TotalEq     JSONFloat64 `json:"totalEq"`
			IsoEq       JSONFloat64 `json:"isoEq"`
			AdjEq       JSONFloat64 `json:"adjEq,omitempty"`
			OrdFroz     JSONFloat64 `json:"ordFroz,omitempty"`
			Imr         JSONFloat64 `json:"imr,omitempty"`
			Mmr         JSONFloat64 `json:"mmr,omitempty"`
			MgnRatio    JSONFloat64 `json:"mgnRatio,omitempty"`
			NotionalUsd JSONFloat64 `json:"notionalUsd,omitempty"`
			Details     []*struct {
				Ccy           string      `json:"ccy"`
				Eq            JSONFloat64 `json:"eq"`
				CashBal       JSONFloat64 `json:"cashBal"`
				IsoEq         JSONFloat64 `json:"isoEq,omitempty"`
				AvailEq       JSONFloat64 `json:"availEq,omitempty"`
				DisEq         JSONFloat64 `json:"disEq"`
				AvailBal      JSONFloat64 `json:"availBal"`
				FrozenBal     JSONFloat64 `json:"frozenBal"`
				OrdFrozen     JSONFloat64 `json:"ordFrozen"`
				Liab          JSONFloat64 `json:"liab,omitempty"`
				Upl           JSONFloat64 `json:"upl,omitempty"`
				UplLib        JSONFloat64 `json:"uplLib,omitempty"`
				CrossLiab     JSONFloat64 `json:"crossLiab,omitempty"`
				IsoLiab       JSONFloat64 `json:"isoLiab,omitempty"`
				MgnRatio      JSONFloat64 `json:"mgnRatio,omitempty"`
				Interest      JSONFloat64 `json:"interest,omitempty"`
				Twap          JSONFloat64 `json:"twap,omitempty"`
				MaxLoan       JSONFloat64 `json:"maxLoan,omitempty"`
				EqUsd         JSONFloat64 `json:"eqUsd"`
				NotionalLever JSONFloat64 `json:"notionalLever,omitempty"`
				StgyEq        JSONFloat64 `json:"stgyEq"`
				IsoUpl        JSONFloat64 `json:"isoUpl,omitempty"`
				UTime         JSONTime    `json:"uTime"`
			} `json:"details,omitempty"`
			UTime JSONTime `json:"uTime"`
		} `json:"data,omitempty"`
	}
)

// GetBalance get account balance
func (c *Client) GetBalance(ctx context.Context, req *GetBalanceRequest, opts ...RequestOption) (*GetBalanceResponse, error) {
	params := map[string]interface{}{}
	if req != nil && req.Ccy != "" {
		params["ccy"] = req.Ccy
	}
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/account/balance",
		SecType:  SecTypePrivate,
		Params:   params,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &GetBalanceResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
