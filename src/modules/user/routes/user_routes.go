package routes

import (
	authMiddlewares "gorm/src/modules/auth/middlewares"
	"gorm/src/modules/user/controller"
	"gorm/src/modules/user/middlewares"
	"gorm/src/modules/user/service"
	"net/http"

	"github.com/gorilla/mux"
)

func UserRoutes(mux *mux.Router, controller *controller.UserController, service service.UserService) {

	mux.HandleFunc("/user/", controller.FindAll).Methods("GET")
	mux.Handle("/user/{id:[0-9]+}", middlewares.UserExistsMiddleware(service)(http.HandlerFunc(controller.GetUserById))).Methods("GET")
	mux.HandleFunc("/user/", authMiddlewares.JWTMiddleware(middlewares.DecodeUserMiddlaware(controller.CreateUser))).Methods("POST")
	mux.Handle("/user/{id:[0-9]+}",
		middlewares.UserExistsMiddleware(service)(
			middlewares.DecodeUserMiddlaware(http.HandlerFunc(controller.UpdateUser)))).Methods("PUT")
	mux.Handle("/user/{id:[0-9]+}", middlewares.UserExistsMiddleware(service)(http.HandlerFunc(controller.DeleteUser))).Methods("DELETE")
	//mux.HandleFunc("/user/{id:[0-9]+}", middlewares.ExitUser(controllers.DeleteUser)).Methods("DELETE")
}
