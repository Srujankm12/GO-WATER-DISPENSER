package models

type UserModel struct{
	UserId string `json:"userid"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	QrId string `json:"qrid"`
}

type UserRepositoryInterface interface{
	AddUserDetailsToDatabase(UserModel) error
	LoginUser(UserModel) error
}