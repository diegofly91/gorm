package routes

import (
	"gorm/src/modules/user/controller"
	"gorm/src/modules/user/middlawares"
	"gorm/src/modules/user/service"
	"net/http"

	"github.com/gorilla/mux"
)

func UserRoutes(mux *mux.Router, controller *controller.UserController, service service.UserService) {

	mux.HandleFunc("/user/", controller.FindAll).Methods("GET")
	mux.Handle("/user/{id:[0-9]+}", middlawares.UserExistsMiddleware(service)(http.HandlerFunc(controller.GetUserById))).Methods("GET")
	mux.HandleFunc("/user/", middlawares.DecodeUserMiddlaware(controller.CreateUser)).Methods("POST")
	mux.Handle("/user/{id:[0-9]+}",
		middlawares.UserExistsMiddleware(service)(
			middlawares.DecodeUserMiddlaware(http.HandlerFunc(controller.UpdateUser)))).Methods("PUT")
	//mux.HandleFunc("/user/{id:[0-9]+}", middlawares.ExitUser(controllers.DeleteUser)).Methods("DELETE")
}
