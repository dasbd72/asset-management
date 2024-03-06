package ws

func (c *Client) StartUserDataStream(params map[string]interface{}) error {
	return c.Send(Request_builder{
		Method:  "userDataStream.start",
		SecType: SecTypeAPIKey,
		Params:  params,
	}.Build())
}
