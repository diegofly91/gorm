package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetId(r *http.Request) int {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	return id
}
