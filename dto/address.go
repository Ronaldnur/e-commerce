package dto

type NewAddressRequest struct {
	Street     string `json:"street" valid:"required"`
	City       string `json:"city" valid:"required"`
	Province   string `json:"province" valid:"required"`
	Country    string `json:"country" valid:"required"`
	PostalCode string `json:"postal_code" valid:"required"`
}
