package models

import (
	"io"
)

type UserModel struct{
	UserId string `json:"userid"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type UserModelInterface interface{
	RegisterUser(*io.ReadCloser) error
	// LoginUser(UserModel) (bool , error)
}