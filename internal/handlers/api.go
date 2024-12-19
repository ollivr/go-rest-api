package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"restful-api/internal/tools"
)

func HandleRequests() http.Handler {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", tools.HandleHomePage)
	myRouter.HandleFunc("/users", tools.AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}", tools.HandleDeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", tools.HandleUpdateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}/{email}", tools.HandleNewUser).Methods("POST")

	return myRouter
}
