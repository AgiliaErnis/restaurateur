package api

// @title Restaurateur API
// @version 0.2.0
// @description Provides info about restaurants in Prague
// @host localhost:8080
// @BasePath /

import (
	"flag"
	"fmt"
	_ "github.com/AgiliaErnis/restaurateur/backend/docs" // Swagger docs
	"github.com/AgiliaErnis/restaurateur/backend/internal/db"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

var allowedEndpoints = [...]string{"/restaurants", "/prague-college/restaurants", "/autocomplete", "/user", "/login", "/register"}

// Run starts the server on the specified port
func Run() {
	var portNum int
	flag.IntVar(&portNum, "p", 8080, "Port number")
	flag.Parse()
	db.CheckDB()
	if portNum < 1024 || portNum > 65535 {
		log.Fatal("Invalid port number, use a number from 1024-65535")
	}
	port := fmt.Sprintf(":%d", portNum)
	r := mux.NewRouter()
	r.HandleFunc("/prague-college/restaurants", pcRestaurantsHandler).Methods(http.MethodGet)
	r.HandleFunc("/restaurants", restaurantsHandler).Methods(http.MethodGet)
	r.HandleFunc("/restaurant/{id:[0-9]+}", restaurantHandler).Methods(http.MethodGet)
	r.HandleFunc("/autocomplete", autocompleteHandler).Methods(http.MethodGet)
	r.HandleFunc("/register", registerHandler).Methods(http.MethodPost)
	r.HandleFunc("/user", userGetHandler).Methods(http.MethodGet)
	r.HandleFunc("/user", userDeleteHandler).Methods(http.MethodDelete)
	r.HandleFunc("/user", userPatchHandler).Methods(http.MethodPatch)
	r.HandleFunc("/login", loginHandler).Methods(http.MethodPost)
	r.HandleFunc("/logout", logoutHandler).Methods(http.MethodGet)
	r.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	r.PathPrefix("/").HandlerFunc(catchAllHandler)
	log.Println("Starting server on", port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk, credentialsOk)(r)))
}
