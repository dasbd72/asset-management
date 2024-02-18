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
			UTime       JSONTime    `json:"uTime"`
			TotalEq     JSONFloat64 `json:"totalEq"`
			IsoEq       JSONFloat64 `json:"isoEq,omitempty"`
			AdjEq       JSONFloat64 `json:"adjEq,omitempty"`
			OrdFroz     JSONFloat64 `json:"ordFroz,omitempty"`
			Imr         JSONFloat64 `json:"imr,omitempty"`
			Mmr         JSONFloat64 `json:"mmr,omitempty"`
			BorrowFroz  JSONFloat64 `json:"borrowFroz,omitempty"`
			MgnRatio    JSONFloat64 `json:"mgnRatio,omitempty"`
			NotionalUsd JSONFloat64 `json:"notionalUsd,omitempty"`
			Upl         JSONFloat64 `json:"upl,omitempty"`
			Details     []*struct {
				Ccy           string      `json:"ccy"`
				Eq            JSONFloat64 `json:"eq"`
				CashBal       JSONFloat64 `json:"cashBal"`
				UTime         JSONTime    `json:"uTime"`
				IsoEq         JSONFloat64 `json:"isoEq,omitempty"`
				AvailEq       JSONFloat64 `json:"availEq,omitempty"`
				DisEq         JSONFloat64 `json:"disEq"`
				FixedBal      JSONFloat64 `json:"fixedBal,omitempty"`
				AvailBal      JSONFloat64 `json:"availBal"`
				FrozenBal     JSONFloat64 `json:"frozenBal"`
				OrdFrozen     JSONFloat64 `json:"ordFrozen,omitempty"`
				Liab          JSONFloat64 `json:"liab,omitempty"`
				Upl           JSONFloat64 `json:"upl,omitempty"`
				UplLiab       JSONFloat64 `json:"uplLiab,omitempty"`
				CrossLiab     JSONFloat64 `json:"crossLiab,omitempty"`
				IsoLiab       JSONFloat64 `json:"isoLiab,omitempty"`
				MgnRatio      JSONFloat64 `json:"mgnRatio,omitempty"`
				Interest      JSONFloat64 `json:"interest,omitempty"`
				Twap          JSONFloat64 `json:"twap,omitempty"`
				MaxLoan       JSONFloat64 `json:"maxLoan,omitempty"`
				EqUsd         JSONFloat64 `json:"eqUsd"`
				BorrowFroz    JSONFloat64 `json:"borrowFroz,omitempty"`
				NotionalLever JSONFloat64 `json:"notionalLever,omitempty"`
				StgyEq        JSONFloat64 `json:"stgyEq"`
				IsoUpl        JSONFloat64 `json:"isoUpl,omitempty"`
				SpotInUseAmt  JSONFloat64 `json:"spotInUseAmt,omitempty"`
				SpotIsoBal    JSONFloat64 `json:"spotIsoBal,omitempty"`
				Imr           JSONFloat64 `json:"imr,omitempty"`
				Mmr           JSONFloat64 `json:"mmr,omitempty"`
			} `json:"details,omitempty"`
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
