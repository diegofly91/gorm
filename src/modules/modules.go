package modules

import (
	"gorm/src/modules/user"

	"github.com/gorilla/mux"
)

func SetupModules(router *mux.Router) {
	// Inicializar los módulos
	user.InitializeUserModule(router)
	// Agrega más inicializaciones de módulos aquí si es necesario
}
