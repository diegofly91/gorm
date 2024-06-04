package user

import (
	"gorm/src/modules/user/controller"
	"gorm/src/modules/user/models"
	"gorm/src/modules/user/repository"
	"gorm/src/modules/user/routes"
	"gorm/src/modules/user/service"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitializeUserModule(router *mux.Router, db *gorm.DB) {
	// Crear la tabla de usuario
	db.AutoMigrate(models.User{})
	// Inicializar el módulo de usuario
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	// Configurar las rutas
	routes.UserRoutes(router, userController, userService)
}
