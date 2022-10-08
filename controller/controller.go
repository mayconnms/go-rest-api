package controller

import "github.com/gorilla/mux"

type Controller interface {
	RegistreRoutes(router *mux.Router)
}
