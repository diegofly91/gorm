package controller

import (
	"encoding/json"
	"gorm/src/common/handlers"
	"gorm/src/modules/user/models"
	"gorm/src/modules/user/service"

	"net/http"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service}
}

func (c *UserController) CreateUser(rw http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		handlers.SendError(rw, http.StatusBadRequest)
		return
	}

	newUser, err := c.service.CreateUser(user)
	if err != nil {
		handlers.SendError(rw, http.StatusInternalServerError)
		return
	}

	handlers.SendData(rw, newUser, http.StatusCreated)
}

/*
func GetUsers(rw http.ResponseWriter, r *http.Request) {
	users := services.FindAll()
	handlers.SendData(rw, users, http.StatusOK)

}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user := services.FindById(id)
	handlers.SendData(rw, user, http.StatusOK)
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user := models.User{}
	decode := json.NewDecoder(r.Body)
	decode.Decode(&user)
	user.Id = int64(id)
	services.Update(user)
	handlers.SendData(rw, user, http.StatusOK)

}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user := services.Delete(id)
	handlers.SendData(rw, user, http.StatusOK)
}
*/
