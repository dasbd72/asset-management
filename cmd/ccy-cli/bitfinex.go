package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	bitfinexRest "github.com/dasbd72/go-exchange-sdk/bitfinex/rest"
	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

func Bitfinex(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	// Create a new bitfinex client
	c := bitfinexRest.NewClient(
		os.Getenv("BFX_API_KEY"),
		os.Getenv("BFX_API_SECRET"),
	)

	// Raw(ctx, c)

	SDK(ctx, c)
}

func Raw(ctx context.Context, c *bitfinexRest.Client) {
	{
		data, err := c.CallAPI(ctx, bitfinexRest.Request_builder{
			Method:   http.MethodPost,
			Endpoint: "/auth/r/funding/offers/Symbol",
			Version:  bitfinexRest.Version2,
			SecType:  bitfinexRest.SecTypePrivate,
			Params:   map[string]interface{}{},
		}.Build())
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(data))
	}
	{
		data, err := c.CallAPI(ctx, bitfinexRest.Request_builder{
			Method:   http.MethodPost,
			Endpoint: "/auth/r/funding/loans/Symbol",
			Version:  bitfinexRest.Version2,
			SecType:  bitfinexRest.SecTypePrivate,
			Params:   map[string]interface{}{},
		}.Build())
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(data))
	}
	{
		data, err := c.CallAPI(ctx, bitfinexRest.Request_builder{
			Method:   http.MethodPost,
			Endpoint: "/auth/r/funding/offers/fUSD/hist",
			Version:  bitfinexRest.Version2,
			SecType:  bitfinexRest.SecTypePrivate,
			Params:   map[string]interface{}{},
		}.Build())
		if err != nil {
			log.Fatal(err)
		}
		// log.Println(string(data))
		container := [][]interface{}{}
		err = json.Unmarshal(data, &container)
		if err != nil {
			log.Fatal(err)
		}
		b, err := json.MarshalIndent(container, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(b))
	}
	// {
	// 	data, err := c.CallAPI(ctx, bitfinexRest.Request_builder{
	// 		Method:   http.MethodPost,
	// 		Endpoint: "/auth/w/funding/offer/submit",
	// 		Version:  bitfinexRest.Version2,
	// 		SecType:  bitfinexRest.SecTypePrivate,
	// 		Params:   map[string]interface{}{},
	// 	}.Build())
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(string(data))
	// }
}

func SDK(ctx context.Context, c *bitfinexRest.Client) {
	{
		// dummy
		data, err := c.GetWallets(ctx)
		if err != nil {
			log.Fatal(err)
		}
		_, err = json.MarshalIndent(data, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
	}
	{
		data, err := c.GetWallets(ctx)
		if err != nil {
			log.Fatal(err)
		}
		b, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(b))
	}
	{
		data, err := c.GetAllActiveFundingOffers(ctx)
		// data, err := c.GetActiveFundingOffers(ctx, "fUSD")
		if err != nil {
			log.Fatal(err)
		}
		b, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(b))
	}
	// {
	// 	data, err := c.SubmitFundingOffer(ctx, models.NewSubmitFundingOfferRequest(models.FRRLIMIT, "fUSD", "150", "0.01", 2))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	b, err := json.MarshalIndent(data, "", "  ")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(string(b))
	// }
}

func WS() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	wsEndpoint := "wss://api-pub.bitfinex.com/ws/2"
	log.Printf("connecting to %s", wsEndpoint)

	c, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	// Read messages from the websocket connection
	go func() {
		type event struct {
			Event string `json:"event"`
		}
		type lst []interface{}
		defer close(done)
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			// log.Printf("recv: %s, type: %v", message, mt)
			var e event
			var l lst
			if err := json.Unmarshal(message, &e); err == nil {
				log.Printf("[Event] recv: %s, type: %v", message, mt)
			} else if err := json.Unmarshal(message, &l); err == nil {
				switch l[1].(string) {
				case "hb":
				case "ws":
					log.Printf("[WalletSnapshot] recv: %s, type: %v", message, mt)
				case "wu":
					log.Printf("[WalletUpdate] recv: %s, type: %v", message, mt)
				default:
					log.Printf("[Default] recv: %s, type: %v", message, mt)
				}
			} else {
				log.Printf("[Unknown] recv: %s, type: %v", message, mt)
			}
		}
	}()

	// Write messages to the websocket connection
	nonce := strconv.FormatInt(time.Now().Unix(), 10)
	payload := "AUTH" + nonce
	err = c.WriteJSON(struct {
		APIKey      string   `json:"apiKey"`
		AuthSig     string   `json:"authSig"`
		AuthNonce   string   `json:"authNonce"`
		AuthPayload string   `json:"authPayload"`
		Event       string   `json:"event"`
		Filter      []string `json:"filter"`
	}{
		APIKey:      os.Getenv("BFX_API_KEY"),
		AuthSig:     sign(os.Getenv("BFX_API_SECRET"), payload),
		AuthNonce:   nonce,
		AuthPayload: payload,
		Event:       "auth",
		Filter:      []string{"funding"},
	})
	if err != nil {
		log.Println("write:", err)
	}

	// Termination handler
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			// Wait for the server to close the connection.
			case <-done:
			// Or force close the connection after a timeout.
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func sign(secret, message string) string {
	mac := hmac.New(sha512.New384, []byte(secret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}
