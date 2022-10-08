package main

import (
	"net/http"
	"rest-api/controller"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	registerRoutes(router)
	http.ListenAndServe(":8080", router)

}

func registerRoutes(router *mux.Router) {
	registerControllerRoutes(controller.UserController{}, router)
}

func registerControllerRoutes(controller controller.Controller, router *mux.Router) {
	controller.RegistreRoutes(router)
}
