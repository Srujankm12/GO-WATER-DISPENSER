package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Shubhangcs/go-water-dispenser/models"
)

type UserControllerInstance struct{
	UserRequestRepositoryInstance models.UserRepositoryInterface
}

func NewUserControllerInstance(ins models.UserRepositoryInterface) *UserControllerInstance {
	return &UserControllerInstance{
		UserRequestRepositoryInstance: ins,
	}
}

func (usr *UserControllerInstance) AddUserHttpRequest(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	data , redrr := io.ReadAll(r.Body)
	if redrr != nil{
		json.NewEncoder(w).Encode(models.ErrorPayload{Message: redrr.Error()})
		return
	}
	var user models.UserModel
	json.Unmarshal(data , &user)
	_ , err := usr.UserRequestRepositoryInstance.AddUserDetailsToDatabase(user)
	if err != nil{
		json.NewEncoder(w).Encode(models.ErrorPayload{Message: err.Error()})
	}
	json.NewEncoder(w).Encode(models.SuccessPayload{Message: "Query Executed Succefully"})
}

