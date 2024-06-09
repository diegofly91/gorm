package auth

import (
	"gorm/src/modules/auth/controller"
	"gorm/src/modules/auth/routes"
	authS "gorm/src/modules/auth/service"
	"gorm/src/modules/user/repository"
	"gorm/src/modules/user/service"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitializeAuthModule(router *mux.Router, db *gorm.DB) {
	// Crear la tabla de usuario
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	authService := authS.NewAuthService(userService)
	// Inicializar el módulo de usuario
	authController := controller.NewAuthController(authService)
	routes.AuthRoutes(router, authController, userService)
}
