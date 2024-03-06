package ws

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	apiKey     string
	apiSecret  string
	wsEndpoint string
	dialer     *websocket.Dialer
	conn       *websocket.Conn

	interrupt chan os.Signal
	done      chan struct{}
	ticker    *time.Ticker
}

const (
	baseWSEndpoint = "wss://ws-api.binance.com:443/ws-api/v3"

	apiKey       = "apiKey"
	signatureKey = "signature"
	timestampKey = "timestamp"
)

func NewClient(apiKey, apiSecret string) *Client {
	return &Client{
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		wsEndpoint: baseWSEndpoint,
		dialer: &websocket.Dialer{
			Proxy:             http.ProxyFromEnvironment,
			HandshakeTimeout:  45 * time.Second,
			EnableCompression: false,
		},

		interrupt: make(chan os.Signal, 1),
		done:      make(chan struct{}),
		ticker:    time.NewTicker(1 * time.Second),
	}
}

func (c *Client) Connect() (err error) {
	signal.Notify(c.interrupt, os.Interrupt)

	c.conn, _, err = c.dialer.Dial(c.wsEndpoint, nil)
	if err != nil {
		return err
	}

	// Read messages from the websocket connection
	go func() {
		defer close(c.done)
		for {
			mt, message, err := c.conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s, type: %v", message, mt)
		}
	}()
	return nil
}

func (c *Client) KeepAlive() error {
	for {
		select {
		case <-c.done:
			return nil
		case <-c.interrupt:
			log.Println("interrupt")
			err := c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				return err
			}
			select {
			case <-c.done:
			case <-time.After(1 * time.Second):
			}
			return nil
		}
	}
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) Send(r *Request, opts ...RequestOption) error {
	for _, opt := range opts {
		opt(r)
	}
	id := uuid.New().String()
	method := r.method
	params := r.params
	if r.secType == SecTypeEd25519 {
		return fmt.Errorf("not implemented")
	} else if r.secType == SecTypeAPIKey {
		params[apiKey] = c.apiKey
	} else if r.secType == SecTypeSigned {
		params[apiKey] = c.apiKey
		params[timestampKey] = currentTimestamp()
		params[signatureKey] = sign(c.apiSecret, params)
	}
	return c.SendJSON(map[string]interface{}{
		"method": method,
		"id":     id,
		"params": params,
	})
}

func (c *Client) SendJSON(v interface{}) error {
	err := c.conn.WriteJSON(v)
	if err != nil {
		return err
	}
	return nil
}

// Login send login request
//
// This method is not implemented yet
func (c *Client) Login() error {
	return c.Send(Request_builder{
		Method:  "session.logon",
		SecType: SecTypeEd25519,
	}.Build())
}

func sign(secret string, v map[string]interface{}) string {
	query := url.Values{}
	for k, v := range v {
		query.Add(k, fmt.Sprintf("%v", v))
	}
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(query.Encode()))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func currentTimestamp() int64 {
	return formatTimestamp(time.Now())
}

// formatTimestamp formats a time into Unix timestamp in milliseconds, as requested by Binance.
func formatTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}
