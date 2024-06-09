package routes

import (
	"gorm/src/modules/auth/controller"
	"gorm/src/modules/auth/middlewares"
	"gorm/src/modules/user/service"
	"net/http"

	"github.com/gorilla/mux"
)

func AuthRoutes(router *mux.Router, controller *controller.AuthController, service service.UserService) {
	router.Handle("/auth/login", middlewares.AuthUserMiddleware(service)(http.HandlerFunc(controller.Login))).Methods("POST")
	router.Handle("/auth/me", middlewares.JWTMiddleware(controller.GetCurrentUser)).Methods("GET")
}
