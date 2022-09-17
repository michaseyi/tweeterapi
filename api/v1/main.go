package main

import (
	"net/http"
	"tweeter/db"
	"tweeter/routes"

	"github.com/gorilla/mux"
)

func main() {

	db.ConnectDB()

	server := mux.NewRouter()
	app := server.PathPrefix("/api/v1").Subrouter()

	// authentication route
	routes.AuthenticationRouter(app.PathPrefix("/auth").Subrouter())

	http.ListenAndServe(":4000", server)
}
