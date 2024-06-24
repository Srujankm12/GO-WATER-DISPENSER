package models

type UserRequestModel struct{
	Name string `json:"name"`
	Phone string `json:"phone"`
	QrId string `json:"qrid"`
}

type UserRequestRepository interface{
	AddUserToDatabase(UserRequestModel) error
}