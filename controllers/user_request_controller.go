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
	if data , readerr := io.ReadAll(r.Body); readerr != nil {
		http.Error(w , "Error While Reading The Data" , http.StatusBadRequest)
		return
	}else{
		var user models.UserModel
		if err := json.Unmarshal(data , &user); err != nil {
			http.Error(w , "Error While converting json to Structs" , http.StatusBadRequest)
			return
		}
		if  err := usr.UserRequestRepositoryInstance.AddUserDetailsToDatabase(user); err != nil{
			http.Error(w , "Error While writing data to Database" , http.StatusConflict)
			return
		} else{
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(models.SuccessPayload{Message: "Registration Successfull"})
		}
	}
	defer r.Body.Close()
}

func (usr *UserControllerInstance) LoginUserHttpRequest(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	data , _ := io.ReadAll(r.Body)
	var user models.UserModel
	json.Unmarshal(data , &user)
	if err := usr.UserRequestRepositoryInstance.LoginUser(user); err != nil {
		http.Error(w , err.Error() , http.StatusBadRequest)
	}
}

