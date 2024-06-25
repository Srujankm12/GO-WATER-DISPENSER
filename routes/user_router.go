package routes

import (
	"database/sql"
	"sync"

	"github.com/Shubhangcs/go-water-dispenser/controllers"
	"github.com/Shubhangcs/go-water-dispenser/repository"
	"github.com/gorilla/mux"
)


func UserRouter(s *sql.DB , router *mux.Router , mut *sync.Mutex){
	usr := repository.NewUserRepository(s , mut)
	cont := controllers.NewUserControllerInstance(usr)

	router.HandleFunc("/user" , cont.AddUserHttpRequest).Methods("POST")
	router.HandleFunc("/login" , cont.LoginUserHttpRequest).Methods("POST")
}