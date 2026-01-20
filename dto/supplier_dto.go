package dto

type CreateSupplierRequest struct {
	Name    string `json:"name" validate:"required,min=3"`
	Email   string `json:"email" validate:"required,email"`
	Address string `json:"address" validate:"required"`
}

type UpdateSupplierRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type SupplierResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}