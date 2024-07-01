package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/Shubhangcs/go-water-dispenser/models"
)

type TransactionRepository struct {
	db  *sql.DB
	mut *sync.Mutex
}

func NewTransactionRepositoryInstance(db *sql.DB, mut *sync.Mutex) *TransactionRepository {
	return &TransactionRepository{
		db:  db,
		mut: mut,
	}
}

func (tr *TransactionRepository) ConfirmTransaction(inp *io.ReadCloser) error {
	data, readErr := io.ReadAll(*inp)
	if readErr != nil {
		return readErr
	}
	var transaction models.Transaction
	convErr := json.Unmarshal(data, &transaction)
	if convErr != nil {
		return convErr
	}
	tr.mut.Lock()
	_, dbErr := tr.db.Exec("INSERT INTO transactions(qrid, quantity, amount, reciptno) VALUES($1, $2, $3, $4)", transaction.QrId, transaction.NoOfLiters, transaction.Ammount, transaction.ReciptNo)
	tr.mut.Unlock()
	if dbErr != nil {
		return dbErr
	}
	return nil
}
func (tr *TransactionRepository) GetQuantity() (string, error) {
	var quantities []int
	rows, err := tr.db.Query("SELECT quantity FROM transactions")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		var quantity int
		if err := rows.Scan(&quantity); err != nil {
			return "", err
		}
		quantities = append(quantities, quantity)
	}

	var strQuantities []string
	for _, quantity := range quantities {
		strQuantities = append(strQuantities, fmt.Sprintf("%d", quantity))
	}
	result := strings.Join(strQuantities, ",")
	return result, nil
}
