package repository

import (
	"database/sql"
	"encoding/json"
	"io"
	"sync"

	"github.com/Shubhangcs/go-water-dispenser/models"
)

type TransactionRepository struct{
	db *sql.DB
	mut *sync.Mutex
}

func NewTransactionRepositoryInstance(db *sql.DB , mut *sync.Mutex) *TransactionRepository {
	return &TransactionRepository{
		db: db,
		mut: mut,
	}
}

func (tr *TransactionRepository) ConfirmTransaction(inp *io.ReadCloser) error {
	data , readErr := io.ReadAll(*inp)
	if readErr != nil {
		return readErr
	}
	var transaction models.Transaction
	convErr := json.Unmarshal(data , &transaction)
	if convErr != nil{
		return convErr
	}
	tr.mut.Lock()
	_ , dbErr := tr.db.Exec("INSERT INTO transactions(qrid , quantity , amount , reciptno) VALUES($1 , $2 , $3 , $4)" , transaction.QrId , transaction.NoOfLiters , transaction.Ammount , transaction.ReciptNo)
	tr.mut.Unlock()
	if dbErr != nil {
		return dbErr
	}
	return nil
}