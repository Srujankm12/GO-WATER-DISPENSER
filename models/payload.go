package models

type ErrorPayload struct {
	Message string `json:"message"`
}

type SuccessPayload struct {
	Message string `json:"message"`
}
type QuantityResponse struct {
	Quantity int `json:"quantity"`
}
