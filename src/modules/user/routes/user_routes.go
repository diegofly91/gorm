package routes

import (
	"gorm/src/modules/user/controller"
	"gorm/src/modules/user/middlawares"

	"github.com/gorilla/mux"
)

func UserRoutes(mux *mux.Router, controller *controller.UserController) {

	//mux.HandleFunc("/user/", controllers.GetUsers).Methods("GET")
	//mux.HandleFunc("/user/{id:[0-9]+}", middlawares.ExitUser(controllers.GetUser)).Methods("GET")
	mux.HandleFunc("/user/", middlawares.DecodeUserMiddlaware(controller.CreateUser)).Methods("POST")
	//mux.HandleFunc("/user/{id:[0-9]+}", middlawares.ExitUser(middlawares.DecodeUserMiddlaware(controllers.UpdateUser))).Methods("PUT")
	//mux.HandleFunc("/user/{id:[0-9]+}", middlawares.ExitUser(controllers.DeleteUser)).Methods("DELETE")
}
