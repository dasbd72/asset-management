package ws

func (c *Client) OrderBook(params map[string]interface{}) error {
	return c.Send(Request_builder{
		Method:  "depth",
		SecType: SecTypeNone,
		Params:  params,
	}.Build())
}

func (c *Client) RecentTrades(params map[string]interface{}) error {
	return c.Send(Request_builder{
		Method:  "trades.recent",
		SecType: SecTypeNone,
		Params:  params,
	}.Build())
}

func (c *Client) KLines(params map[string]interface{}) error {
	return c.Send(Request_builder{
		Method:  "klines",
		SecType: SecTypeNone,
		Params:  params,
	}.Build())
}
