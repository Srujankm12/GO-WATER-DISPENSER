package repository

import (
	"database/sql"

	"github.com/Shubhangcs/go-water-dispenser/models"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (nr *UserRepository) AddUserDetailsToDatabase(usr models.UserModel) (*sql.Rows, error) {
	res, err := nr.Db.Query("INSERT INTO transaction(userid,name,phone,qrid) VALUES(?,?,?,?)", usr.UserId, usr.Name, usr.Phone, usr.QrId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
