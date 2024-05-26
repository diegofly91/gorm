package user

import (
	"gorm/src/modules/user/routes"

	"gorm/src/modules/user/models"

	"github.com/gorilla/mux"
)

func InitializeUserModule(router *mux.Router) {
	models.MigrarUser()
	// Configurar las rutas
	routes.UserRoutes(router)
}
