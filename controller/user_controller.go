package controller

import (
	"rest-api/service"

	"github.com/gorilla/mux"
)

type UserController struct{}

func (t UserController) RegistreRoutes(router *mux.Router) {

	router.HandleFunc("/user/{id}", service.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{id}", service.GetUser).Methods("GET")
	router.HandleFunc("/user", service.CreatUser).Methods("POST")

}
