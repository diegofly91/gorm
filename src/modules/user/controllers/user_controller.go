package controllers

import (
	"encoding/json"
	"fmt"
	"gorm/src/common/handlers"
	"gorm/src/modules/user/models"
	"gorm/src/modules/user/services"
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

	fmt.Println(" ============================= ")
	fmt.Println(user)
	newUser := services.Create(user)
	fmt.Println(newUser)
	handlers.SendData(rw, newUser, http.StatusCreated)

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
