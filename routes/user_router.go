package routes

import (
	"database/sql"

	"github.com/Shubhangcs/go-water-dispenser/controllers"
	"github.com/Shubhangcs/go-water-dispenser/repository"
	"github.com/gorilla/mux"
)


func UserRouter(s *sql.DB , router *mux.Router){
	usr := repository.NewUserRepository(s)
	cont := controllers.NewUserControllerInstance(usr)

	router.HandleFunc("/user" , cont.AddUserHttpRequest).Methods("POST")
}