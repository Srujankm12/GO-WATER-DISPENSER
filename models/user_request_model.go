package models

import "database/sql"

type UserModel struct{
	UserId string `json:"userid"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	QrId string `json:"qrid"`
}

type UserRepositoryInterface interface{
	AddUserDetailsToDatabase(UserModel)(*sql.Rows , error)
}