package dto

import "time"


type CreatePurchasingRequest struct {
	SupplierID uint                     `json:"supplier_id"`
	Items      []PurchasingItemRequest  `json:"items"`
}

type PurchasingItemRequest struct {
	ItemID uint `json:"item_id"`
	Qty    int  `json:"qty"`
}


type PurchasingResponse struct {
	ID         uint                    `json:"id"`
	Date       time.Time               `json:"date"`
	GrandTotal float64                 `json:"grand_total"`
	Supplier   SupplierMiniResponse    `json:"supplier"`
	User       UserMiniResponse        `json:"user"`
	Details    []PurchasingDetailResponse `json:"items"`
}

type PurchasingDetailResponse struct {
	ItemID   uint    `json:"item_id"`
	ItemName string  `json:"item_name"`
	Qty      int     `json:"qty"`
	Price    float64 `json:"price"`
	SubTotal float64 `json:"sub_total"`
}

type SupplierMiniResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserMiniResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}
