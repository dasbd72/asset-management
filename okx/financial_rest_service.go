package okx

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	GetEarnOffersRequest struct {
		ProductId    string `json:"productId"`
		ProtocolType string `json:"protocolType"`
		Ccy          string `json:"ccy"`
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
		ProductId    string    `json:"productId"`
		ProtocolType string    `json:"protocolType"`
		Ccy          string    `json:"ccy"`
		State        JSONInt64 `json:"state"`
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

	GetETHStakingBalanceRequest struct {
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
		Ccy string `json:"ccy,omitempty"`
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
		Ccy string `json:"ccy,omitempty"`
		// Pagination of data to return records earlier than the requested ts,
		// Unix timestamp format in milliseconds, e.g. 1597026383085
		After string `json:"after,omitempty"`
		// Pagination of data to return records newer than the requested ts,
		// Unix timestamp format in milliseconds, e.g. 1597026383085
		Before string `json:"before,omitempty"`
		// Number of results per request.
		// The maximum is 100. The default is 100.
		Limit JSONInt64 `json:"limit,omitempty"`
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

// GetEarnOffers get the available earn offers
func (c *Client) GetEarnOffers(ctx context.Context, req *GetEarnOffersRequest, opts ...RequestOption) (*GetEarnOffersResponse, error) {
	params := map[string]interface{}{}
	if req.ProductId != "" {
		params["productId"] = req.ProductId
	}
	if req.ProtocolType != "" {
		params["protocolType"] = req.ProtocolType
	}
	if req.Ccy != "" {
		params["ccy"] = req.Ccy
	}
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/finance/staking-defi/offers",
		SecType:  SecTypePrivate,
		Params:   params,
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

// GetActiveEarnOrders get the active earn orders
func (c *Client) GetActiveEarnOrders(ctx context.Context, req *GetActiveEarnOrdersRequest, opts ...RequestOption) (*GetActiveEarnOrdersResponse, error) {
	params := map[string]interface{}{}
	if req.ProductId != "" {
		params["productId"] = req.ProductId
	}
	if req.ProtocolType != "" {
		params["protocolType"] = req.ProtocolType
	}
	if req.Ccy != "" {
		params["ccy"] = req.Ccy
	}
	if req.State != 0 {
		params["state"] = req.State
	}
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/finance/staking-defi/orders-active",
		SecType:  SecTypePrivate,
		Params:   params,
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
func (c *Client) GetETHStakingBalance(ctx context.Context, req *GetETHStakingBalanceRequest, opts ...RequestOption) (*GetETHStakingBalanceResponse, error) {
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

// GetSavingBalance get balances in the saving account
func (c *Client) GetSavingBalance(ctx context.Context, req *GetSavingBalanceRequest, opts ...RequestOption) (*GetSavingBalanceResponse, error) {
	params := map[string]interface{}{}
	if req.Ccy != "" {
		params["ccy"] = req.Ccy
	}
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/finance/savings/balance",
		SecType:  SecTypePrivate,
		Params:   params,
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

// GetLendingHistory get the lending history
func (c *Client) GetLendingHistory(ctx context.Context, req *GetLendingHistoryRequest, opts ...RequestOption) (*GetLendingHistoryResponse, error) {
	params := map[string]interface{}{}
	if req.Ccy != "" {
		params["ccy"] = req.Ccy
	}
	if req.After != "" {
		params["after"] = req.After
	}
	if req.Before != "" {
		params["before"] = req.Before
	}
	if req.Limit != 0 {
		params["limit"] = req.Limit
	}
	res, err := c.CallAPI(ctx, Request_builder{
		Method:   http.MethodGet,
		Endpoint: "/api/v5/finance/savings/lending-history",
		SecType:  SecTypePrivate,
		Params:   params,
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
