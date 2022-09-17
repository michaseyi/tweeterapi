package routes

import (
	"net/http"
	"tweeter/controller"

	"github.com/gorilla/mux"
)

func AuthenticationRouter(router *mux.Router) {
	router.HandleFunc("/login/user", controller.AuthController.UserLogin).Methods(http.MethodPost)
	router.HandleFunc("/register/user", controller.AuthController.UserRegistration).Methods(http.MethodPost)
}
