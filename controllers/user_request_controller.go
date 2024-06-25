package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/Shubhangcs/go-water-dispenser/models"
)

type UserControllerInstance struct{
	UserRepositoryInstance models.UserModelInterface
}

func NewUserControllerInstance(ins models.UserModelInterface) *UserControllerInstance {
	return &UserControllerInstance{
		UserRepositoryInstance: ins,
	}
}

func (usr *UserControllerInstance) RegisterUserHttpRequest(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	if err := usr.UserRepositoryInstance.RegisterUser(&r.Body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorPayload{Message: err.Error()})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.SuccessPayload{Message: "Registration Successfull"})
	defer r.Body.Close()
}

// func (usr *UserControllerInstance) LoginUserHttpRequest(w http.ResponseWriter , r *http.Request){
// 	w.Header().Set("Content-Type" , "application/json")
// 	if jsonData , readerr := io.ReadAll(r.Body); readerr != nil{
// 		http.Error(w , "Unable To Read The Data",http.StatusBadRequest)
// 	}else{
// 		var userModel
// 	}
	
// }

