package routes

import (
	"gorm/src/modules/user/controllers"
	"gorm/src/modules/user/middlawares"

	"github.com/gorilla/mux"
)

func UserRoutes(mux *mux.Router) {
	mux.HandleFunc("/user/", controllers.GetUsers).Methods("GET")
	mux.HandleFunc("/user/{id:[0-9]+}", middlawares.ExitUser(controllers.GetUser)).Methods("GET")
	mux.HandleFunc("/user/", middlawares.DecodeUserMiddlaware(controllers.CreateUser)).Methods("POST")
	mux.HandleFunc("/user/{id:[0-9]+}", middlawares.ExitUser(middlawares.DecodeUserMiddlaware(controllers.UpdateUser))).Methods("PUT")
	mux.HandleFunc("/user/{id:[0-9]+}", middlawares.ExitUser(controllers.DeleteUser)).Methods("DELETE")
}
