package modules

import (
	db "gorm/src/config"
	"gorm/src/modules/user"

	"github.com/gorilla/mux"
)

func SetupModules(router *mux.Router) {
	user.InitializeUserModule(router, db.Database)
}
