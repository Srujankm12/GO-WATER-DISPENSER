package models

import "io"

type Transaction struct {
	QrId       string `json:"qrid"`
	NoOfLiters string `json:"liters"`
	Ammount    string `json:"ammount"`
	ReciptNo   string `json:"reciptno"`
}

type TransactionInterface interface {
	ConfirmTransaction(inp *io.ReadCloser) error
	GetQuantity() (int, error)
}
