package controllers

import (
	"net/http"
)

func GetApi2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to tollbooth Demo!"))
}
