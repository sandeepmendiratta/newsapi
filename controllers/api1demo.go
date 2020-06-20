package controllers

import (
"net/http"
)


func GetApi1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to DemoAPI!"))
}