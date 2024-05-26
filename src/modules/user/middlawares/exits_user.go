package middlawares

import (
	"gorm/src/common/handlers"
	"gorm/src/modules/user/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ExitUser(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		if user := services.FindById(id); user.Id == 0 {
			handlers.SendError(rw, http.StatusNotFound)
			return
		}
		next.ServeHTTP(rw, r)
	}
}
