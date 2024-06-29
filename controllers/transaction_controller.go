package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Shubhangcs/go-water-dispenser/models"
)

type TransactionController struct {
	transactionInterface models.TransactionInterface
}

func NewTransactionControllerInstance(ti models.TransactionInterface) *TransactionController {
	return &TransactionController{
		transactionInterface: ti,
	}
}

func (tc *TransactionController) ConfirmTransactionImpl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := tc.transactionInterface.ConfirmTransaction(&r.Body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorPayload{Message: "Unable to Complete the Transaction..."})
		return
	}
	json.NewEncoder(w).Encode(models.SuccessPayload{Message: "Transaction Successful. Please Wait..."})
}

func (tc *TransactionController) GetQuantityImpl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	quantity, err := tc.transactionInterface.GetQuantity()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorPayload{Message: "Unable to fetch quantity"})
		return
	}
	json.NewEncoder(w).Encode(models.QuantityResponse{Quantity: quantity})
}
