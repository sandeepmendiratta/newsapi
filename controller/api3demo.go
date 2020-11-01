package controller

import (
	"net/http"
)

func GetApi3(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to Basicauth Demo!"))
}
