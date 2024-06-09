package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"gorm/src/common/handlers"
	"gorm/src/modules/auth/models"
	"gorm/src/modules/user/service"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func AuthUserMiddleware(service service.UserService) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

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
			var login models.LoginUser
			if err := json.Unmarshal(body, &login); err != nil {
				fmt.Println(err)
				handlers.SendError(rw, http.StatusBadRequest)
				return
			}
			validate := validator.New()
			// Validar la estructura del usuario
			if err := validate.Struct(login); err != nil {
				fmt.Println(err)
				rw.WriteHeader(http.StatusUnprocessableEntity)
				fmt.Fprintln(rw, err.Error())
				return
			}

			// Buscar el usuario por nombre de usuario
			if user, err := service.FindUserByUsername(login.Username); err != nil {
				handlers.SendErrorWithMessage(rw, http.StatusInternalServerError, "El usuario o la contraseña son incorrectos")
				return
			} else if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
				handlers.SendErrorWithMessage(rw, http.StatusInternalServerError, "El usuario o la contraseña son incorrectos")
				return
			}
			// Comparar la contraseña ingresada con la contraseña almacenada en la base de datos
			next.ServeHTTP(rw, r)
		})
	}
}
