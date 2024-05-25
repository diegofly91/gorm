package handlers

import (
	"encoding/json"
	"gorm/db"
	"gorm/models"
	"strconv"

	"net/http"

	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	users := models.Users{}
	db.Database.Find(&users)
	sendData(rw, users, http.StatusOK)

}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	if user, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		sendData(rw, user, http.StatusOK)
	}
}
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&user); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database.Save(&user)
		sendData(rw, user, http.StatusCreated)
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {

	var userId int64
	if user_act, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		userId = user_act.Id
		user := models.User{}
		decode := json.NewDecoder(r.Body)
		if err := decode.Decode(&user); err != nil {
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			user.Id = userId
			db.Database.Save(&user)
			sendData(rw, user, http.StatusOK)
		}
	}

}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	if user, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		db.Database.Delete(&user)
		sendData(rw, user, http.StatusOK)
	}
}
func getUserById(r *http.Request) (models.User, *gorm.DB) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user := models.User{}
	if err := db.Database.First(&user, id); err.Error != nil {
		return user, err
	} else {
		return user, nil
	}
}
