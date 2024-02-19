package okx

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	GetEarnOffersRequest struct {
		params map[string]interface{}
	}

	GetEarnOffersResponse struct {
		BasicResponse
		Offers []struct {
			Ccy          string      `json:"ccy"`
			ProductId    JSONInt64   `json:"productId"`
			Protocol     string      `json:"protocol"`
			ProtocolType string      `json:"protocolType"`
			Term         JSONInt64   `json:"term"`
			Apy          JSONFloat64 `json:"apy"`
			EarlyRedeem  bool        `json:"earlyRedeem"`
			InvestData   []struct {
				Ccy    string      `json:"ccy"`
				Bal    JSONFloat64 `json:"bal"`
				MinAmt JSONFloat64 `json:"minAmt"`
				MaxAmt JSONFloat64 `json:"maxAmt"`
			} `json:"investData"`
			EarningData []struct {
				Ccy         string    `json:"ccy"`
				EarningType JSONInt64 `json:"earningType"`
			} `json:"earningData"`
			State string `json:"state"`
		} `json:"data,omitempty"`
	}

	GetActiveEarnOrdersRequest struct {
		params map[string]interface{}
	}

	GetActiveEarnOrdersResponse struct {
		BasicResponse
		Orders []struct {
			Ccy          string      `json:"ccy"`
			OrdId        JSONInt64   `json:"ordId"`
			ProductId    JSONInt64   `json:"productId"`
			State        string      `json:"state"`
			Protocol     string      `json:"protocol"`
			ProtocolType string      `json:"protocolType"`
			Term         JSONInt64   `json:"term"`
			Apy          JSONFloat64 `json:"apy"`
			InvestData   []struct {
				Ccy string      `json:"ccy"`
				Amt JSONFloat64 `json:"amt"`
			} `json:"investData"`
			EarningData []struct {
				Ccy         string      `json:"ccy"`
				EarningType JSONInt64   `json:"earningType"`
				Earnings    JSONFloat64 `json:"earnings"`
			} `json:"earningData"`
			PurchasedTime            JSONTime `json:"purchasedTime"`
			EstSettlementTime        JSONTime `json:"estSettlementTime,omitempty"`
			CancelRedemptionDeadline JSONTime `json:"cancelRedemptionDeadline,omitempty"`
			Tag                      string   `json:"tag,omitempty"`
		} `json:"data,omitempty"`
	}

	GetETHStakingBalanceResponse struct {
		BasicResponse
		Balances []struct {
			Ccy                   string      `json:"ccy"`
			Amt                   JSONFloat64 `json:"amt"`
			LatestInterestAccrual JSONFloat64 `json:"latestInterestAccrual"`
			TotalInterestAccrual  JSONFloat64 `json:"totalInterestAccrual"`
			Ts                    JSONTime    `json:"ts"`
		} `json:"data,omitempty"`
	}

	GetSavingBalanceRequest struct {
		params map[string]interface{}
	}

	GetSavingBalanceResponse struct {
		BasicResponse
		Balances []struct {
			Ccy        string      `json:"ccy"`
			Amt        JSONFloat64 `json:"amt"`
			Earnings   JSONFloat64 `json:"earnings"`
			Rate       JSONFloat64 `json:"rate"`
			LoanAmt    JSONFloat64 `json:"loanAmt"`
			PendingAmt JSONFloat64 `json:"pendingAmt"`
			RedemptAmt JSONFloat64 `json:"redemptAmt,omitempty"` // Deprecated
		} `json:"data,omitempty"`
	}

	GetLendingHistoryRequest struct {
		params map[string]interface{}
	}

	GetLendingHistoryResponse struct {
		BasicResponse
		Records []struct {
			Ccy      string      `json:"ccy"`
			Amt      JSONFloat64 `json:"amt"`
			Earnings JSONFloat64 `json:"earnings"`
			Rate     JSONFloat64 `json:"rate"`
			Ts       JSONTime    `json:"ts"`
		} `json:"data,omitempty"`
	}
)

func NewGetEarnOffersRequest() *GetEarnOffersRequest {
	return &GetEarnOffersRequest{
		params: map[string]interface{}{},
	}
}

func (r *GetEarnOffersRequest) ProductId(productId string) *GetEarnOffersRequest {
	r.params["productId"] = productId
	return r
}

func (r *GetEarnOffersRequest) ProtocolType(protocolType string) *GetEarnOffersRequest {
	r.params["protocolType"] = protocolType
	return r
}

func (r *GetEarnOffersRequest) Ccy(ccy string) *GetEarnOffersRequest {
	r.params["ccy"] = ccy
	return r
}

// GetEarnOffers get the available earn offers
func (c *Client) GetEarnOffers(ctx context.Context, req *GetEarnOffersRequest, opts ...RequestOption) (*GetEarnOffersResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/finance/staking-defi/offers",
		SecType:  SecTypePrivate,
		Params:   req.params,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &GetEarnOffersResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewGetActiveEarnOrdersRequest() *GetActiveEarnOrdersRequest {
	return &GetActiveEarnOrdersRequest{
		params: map[string]interface{}{},
	}
}

func (r *GetActiveEarnOrdersRequest) ProductId(productId string) *GetActiveEarnOrdersRequest {
	r.params["productId"] = productId
	return r
}

func (r *GetActiveEarnOrdersRequest) ProtocolType(protocolType string) *GetActiveEarnOrdersRequest {
	r.params["protocolType"] = protocolType
	return r
}

func (r *GetActiveEarnOrdersRequest) Ccy(ccy string) *GetActiveEarnOrdersRequest {
	r.params["ccy"] = ccy
	return r
}

func (r *GetActiveEarnOrdersRequest) State(state JSONInt64) *GetActiveEarnOrdersRequest {
	r.params["state"] = state
	return r
}

// GetActiveEarnOrders get the active earn orders
func (c *Client) GetActiveEarnOrders(ctx context.Context, req *GetActiveEarnOrdersRequest, opts ...RequestOption) (*GetActiveEarnOrdersResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/finance/staking-defi/orders-active",
		SecType:  SecTypePrivate,
		Params:   req.params,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &GetActiveEarnOrdersResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetETHStakingBalance get balances in the ETH staking account
func (c *Client) GetETHStakingBalance(ctx context.Context, opts ...RequestOption) (*GetETHStakingBalanceResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/finance/staking-defi/eth/balance",
		SecType:  SecTypePrivate,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &GetETHStakingBalanceResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewGetSavingBalanceRequest() *GetSavingBalanceRequest {
	return &GetSavingBalanceRequest{
		params: map[string]interface{}{},
	}
}

func (r *GetSavingBalanceRequest) Ccy(ccy string) *GetSavingBalanceRequest {
	r.params["ccy"] = ccy
	return r
}

// GetSavingBalance get balances in the saving account
func (c *Client) GetSavingBalance(ctx context.Context, req *GetSavingBalanceRequest, opts ...RequestOption) (*GetSavingBalanceResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/finance/savings/balance",
		SecType:  SecTypePrivate,
		Params:   req.params,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &GetSavingBalanceResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func NewGetLendingHistoryRequest() *GetLendingHistoryRequest {
	return &GetLendingHistoryRequest{
		params: map[string]interface{}{},
	}
}

func (r *GetLendingHistoryRequest) Ccy(ccy string) *GetLendingHistoryRequest {
	r.params["ccy"] = ccy
	return r
}

// After sets pagination of data to return records earlier than the requested ts,
// Unix timestamp format in milliseconds, e.g. 1597026383085
func (r *GetLendingHistoryRequest) After(after int64) *GetLendingHistoryRequest {
	r.params["after"] = after
	return r
}

// Before sets pagination of data to return records newer than the requested ts,
// Unix timestamp format in milliseconds, e.g. 1597026383085
func (r *GetLendingHistoryRequest) Before(before int64) *GetLendingHistoryRequest {
	r.params["before"] = before
	return r
}

// Limit sets number of results per request. The maximum is 100. The default is 100.
func (r *GetLendingHistoryRequest) Limit(limit int64) *GetLendingHistoryRequest {
	r.params["limit"] = limit
	return r
}

// GetLendingHistory get the lending history
func (c *Client) GetLendingHistory(ctx context.Context, req *GetLendingHistoryRequest, opts ...RequestOption) (*GetLendingHistoryResponse, error) {
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/finance/savings/lending-history",
		SecType:  SecTypePrivate,
		Params:   req.params,
	}.Build(), opts...)
	if err != nil {
		return nil, err
	}
	data := &GetLendingHistoryResponse{}
	err = json.Unmarshal(res, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
