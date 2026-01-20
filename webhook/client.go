package webhook

import (
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/wahyujatirestu/simple-procurement-system/dto"
)

type Client interface {
	SendPurchaseCreated(payload dto.PurchaseWebhookPayload)
}

type client struct {
	http *resty.Client
	url  string
}

func NewClient() Client {
	return &client{
		http: resty.New(),
		url:  os.Getenv("WEBHOOK_URL"),
	}
}

func (c *client) SendPurchaseCreated(payload dto.PurchaseWebhookPayload) {
	if c.url == "" {
		log.Println("WEBHOOK_URL not set")
		return
	}

	_, err := c.http.R().
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post(c.url)

	if err != nil {
		log.Println("failed to send webhook:", err)
	}
}
