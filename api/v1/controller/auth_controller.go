package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"tweeter/db/repository"
)

var AuthController = authController{}

type authController struct{}

func (a authController) UserLogin(w http.ResponseWriter, r *http.Request) {

}

type loginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (a authController) UserRegistration(w http.ResponseWriter, r *http.Request) {
	loginDetails := loginPayload{}
	err := json.NewDecoder(r.Body).Decode(&loginDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validation of input
	// TODO: move to seperate function
	if len(strings.TrimSpace(loginDetails.Email)) == 0 {
		http.Error(w, "empty email", http.StatusBadRequest)
		return
	}
	if len(strings.TrimSpace(loginDetails.Password)) == 0 {
		http.Error(w, "empty password", http.StatusBadRequest)
		return
	}

	fmt.Println(loginDetails)
	err = repository.UserRepository.CreateUser(loginDetails.Username, loginDetails.Email, loginDetails.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := json.Marshal(loginDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
