package middlawares

import (
	"encoding/json"
	"gorm/src/common/handlers"
	"gorm/src/modules/user/models"
	"net/http"
)

func DecodeUser(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		user := models.User{}
		decode := json.NewDecoder(r.Body)
		if err := decode.Decode(&user); err != nil {
			handlers.SendError(rw, http.StatusUnprocessableEntity)
		}
		next.ServeHTTP(rw, r)
	}
}
