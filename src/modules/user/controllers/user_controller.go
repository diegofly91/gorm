package controllers

import (
	"encoding/json"
	"gorm/src/modules/user/models"
	"gorm/src/modules/user/services"

	"gorm/src/common/handlers"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)

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
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decode := json.NewDecoder(r.Body)
	decode.Decode(&user)
	newuser := services.Create(user)
	handlers.SendData(rw, newuser, http.StatusCreated)

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
