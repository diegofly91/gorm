package middlewares

import (
	"net/http"

	"gorm/src/common/handlers"
	"gorm/src/modules/user/service"

	"github.com/gorilla/mux"
)

func UserExistsMiddleware(service service.UserService) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			id := handlers.GetId(r)
			if user, err := service.FindById(id); user.Id == 0 || err != nil {
				handlers.SendError(rw, http.StatusNotFound)
				return
			}
			next.ServeHTTP(rw, r)
		})
	}
}
