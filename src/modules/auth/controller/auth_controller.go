package controller

import (
	"encoding/json"
	"net/http"

	"gorm/src/common/handlers"
	"gorm/src/modules/auth/models"
	"gorm/src/modules/auth/service"
)

type AuthController struct {
	service service.AuthService
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{service: service}
}

func (c *AuthController) Login(rw http.ResponseWriter, r *http.Request) {
	var credentials models.LoginUser
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		handlers.SendError(rw, http.StatusBadRequest)
		return
	}
	token, err := c.service.Login(credentials.Username, credentials.Password)
	if err != nil {
		handlers.SendError(rw, http.StatusUnauthorized)
		return
	}
	handlers.SendData(rw, map[string]string{"token": token}, http.StatusOK)
}

func (c *AuthController) GetCurrentUser(rw http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(models.UserPayload)
	handlers.SendData(rw, user, http.StatusOK)
}
