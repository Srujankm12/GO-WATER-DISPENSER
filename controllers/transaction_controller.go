package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Shubhangcs/go-water-dispenser/models"
)


type TransactionController struct{
	transactionInterface models.TransactionInterface
}

func NewTransactionControllerInstance(ti models.TransactionInterface) *TransactionController{
	return &TransactionController{
		transactionInterface: ti,
	}
}

func (tc *TransactionController) ConfirmTransactionImpl(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	if err := tc.transactionInterface.ConfirmTransaction(&r.Body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorPayload{Message: "Unable to Complete the Transaction..."})
	}
	json.NewEncoder(w).Encode(models.SuccessPayload{Message: "Transaction Successfull Please Wait..."})
}