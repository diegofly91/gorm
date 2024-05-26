package routes

import (
	"gorm/src/modules/user/controllers"
	"gorm/src/modules/user/middlawares"

	"github.com/gorilla/mux"
)

func UserRoutes(mux *mux.Router) {
	mux.HandleFunc("/api/user/", controllers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", middlawares.ExitUser(controllers.GetUser)).Methods("GET")
	mux.HandleFunc("/api/user/", middlawares.DecodeUser(controllers.CreateUser)).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", middlawares.ExitUser(middlawares.DecodeUser(controllers.UpdateUser))).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", middlawares.ExitUser(controllers.DeleteUser)).Methods("DELETE")
}
