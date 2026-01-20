package dto

import "time"

type PurchaseWebhookPayload struct {
	PurchaseID uint                   `json:"purchase_id"`
	Supplier   SupplierMiniResponse   `json:"supplier"`
	User       UserMiniResponse       `json:"user"`
	GrandTotal float64                `json:"grand_total"`
	Items      []PurchaseWebhookItem  `json:"items"`
	CreatedAt  time.Time              `json:"created_at"`
}

type PurchaseWebhookItem struct {
	ItemID   uint    `json:"item_id"`
	ItemName string  `json:"item_name"`
	Qty      int     `json:"qty"`
	SubTotal float64 `json:"sub_total"`
}
