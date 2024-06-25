package repository

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/Shubhangcs/go-water-dispenser/models"
)

type UserRepository struct {
	Db *sql.DB
	mut *sync.Mutex
}

func NewUserRepository(db *sql.DB , mut *sync.Mutex) *UserRepository {
	return &UserRepository{
		Db: db,
		mut: mut,
	}
}

func (nr *UserRepository) AddUserDetailsToDatabase(usr models.UserModel)  error{
	nr.mut.Lock()
	_ , err := nr.Db.Exec("INSERT INTO transaction(userid,name,phone,qrid) VALUES($1,$2,$3,$4)", usr.UserId, usr.Name, usr.Phone, usr.QrId)
	nr.mut.Unlock()
	if err != nil {
		return  err
	}

	return  nil
}

func (nr *UserRepository) LoginUser(usr models.UserModel)  error {
	var ctx context.Context
	var user models.UserModel
	rows , err := nr.Db.QueryContext(ctx , "Select $1 , $2 FROM transaction" , usr.Name , usr.Phone)
	if err != nil {
		fmt.Println(err.Error())
	}
	rows.Close()
	name := make([]models.UserModel , 0)
	for rows.Next() {
		if err := rows.Scan(&user); err != nil {
			panic(err)
		}
		name = append(name, user)
	}
	 fmt.Println(name)
	return nil
}
