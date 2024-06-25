package repository

import (
	"database/sql"
	"encoding/json"
	// "fmt"
	"io"
	// "net/http"
	"sync"

	"github.com/Shubhangcs/go-water-dispenser/models"
)

type UserRepository struct {
	db *sql.DB
	mut *sync.Mutex
}

func NewUserRepository(db *sql.DB , mut *sync.Mutex) *UserRepository {
	return &UserRepository{
		db: db,
		mut: mut,
	}
}

func (repo *UserRepository) RegisterUser(req *io.ReadCloser)  error {
	var user models.UserModel
	data , readErr := io.ReadAll(*req)
	if readErr != nil {
		return readErr
	}
	unmarErr := json.Unmarshal(data , &user)
	if unmarErr != nil {
		return unmarErr
	}
	_ , queryErr := repo.db.Exec("INSERT INTO users(userid,name,phone,qrid) VALUES($1,$2,$3,$4)", user.UserId, user.Name, user.Phone)
	if queryErr != nil{
		return queryErr
	}
	return nil
}

// func (repo *UserRepository) LoginUser(user *models.UserModel)  (bool , error) {
// 	var userModel models.UserModel
// 	rows := repo.db.QueryRow("Select * FROM transaction WHERE name=$1 AND phone=$2" , user.Name , user.Phone)
// 	if err := rows.Scan( &userModel.UserId,&userModel.Name , &userModel.Phone); err != nil{
// 		fmt.Println(err.Error())
// 		return false , err
// 	}
// 	return true , nil
// }
