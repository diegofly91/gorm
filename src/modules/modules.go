package modules

import (
	db "gorm/src/config"
	"gorm/src/modules/auth"
	"gorm/src/modules/user"

	"github.com/gorilla/mux"
)

func SetupModules(router *mux.Router) {
	auth.InitializeAuthModule(router, db.Database)
	user.InitializeUserModule(router, db.Database)
}
