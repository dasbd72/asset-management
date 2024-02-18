package max

import (
	"encoding/json"
	"net/http"
)

type (
	Ticker struct {
		At       JSONTime    `json:"at"`
		Buy      JSONFloat64 `json:"buy"`
		BuyVol   JSONFloat64 `json:"buy_vol"`
		Sell     JSONFloat64 `json:"sell"`
		SellVol  JSONFloat64 `json:"sell_vol"`
		Open     JSONFloat64 `json:"open"`
		Low      JSONFloat64 `json:"low"`
		High     JSONFloat64 `json:"high"`
		Last     JSONFloat64 `json:"last"`
		Vol      JSONFloat64 `json:"vol"`
		VolInBTC JSONFloat64 `json:"vol_in_btc"`
	}

	Tickers map[string]*Ticker
)

func GetUsdtToTWD() (float64, error) {
	ticker, err := GetTicker("usdttwd")
	if err != nil {
		return 0, err
	}
	return ticker.Last.Float64(), nil
}

func (c *Client) GetTicker(symbol string) (*Ticker, error) {
	return GetTicker(symbol)
}

func GetTicker(symbol string) (*Ticker, error) {
	res, err := http.Get(baseAPImainURL + "/api/v2/tickers/" + symbol)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data := &Ticker{}
	err = json.NewDecoder(res.Body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) GetTickers() (*Tickers, error) {
	return GetTickers()
}

func GetTickers() (*Tickers, error) {
	res, err := http.Get(baseAPImainURL + "/api/v2/tickers")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data := &Tickers{}
	err = json.NewDecoder(res.Body).Decode(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
