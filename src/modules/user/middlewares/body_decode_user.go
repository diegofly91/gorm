package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gorm/src/common/handlers"
	"gorm/src/modules/user/models"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func DecodeUserMiddlaware(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Leer el cuerpo de la solicitud original
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			handlers.SendError(rw, http.StatusInternalServerError)
			return
		}
		// Restablecer el cuerpo de la solicitud para que pueda ser leído nuevamente
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		// Decodificar el cuerpo de la solicitud JSON en una estructura de usuario
		var user models.User
		if err := json.Unmarshal(body, &user); err != nil {
			fmt.Println(err)
			handlers.SendError(rw, http.StatusBadRequest)
			return
		}
		validate := validator.New()
		if err := validate.Struct(user); err != nil {
			fmt.Println(err)
			rw.WriteHeader(http.StatusUnprocessableEntity)
			fmt.Fprintln(rw, err.Error())
			return
		}
		next(rw, r)
	}
}
