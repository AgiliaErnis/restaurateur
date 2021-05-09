package api

import (
	"flag"
	"fmt"
	_ "github.com/AgiliaErnis/restaurateur/backend/docs"
	"github.com/AgiliaErnis/restaurateur/backend/internal/db"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title Restaurateur API
// @version 0.2.0
// @description Provides info about restaurants in Prague
// @host localhost:8080
// @BasePath /

// autocompleteHandler godoc
// @Summary Autocomplete backend
// @Description Provides restaurant candidates for autocompletion based on provided input
// @Tags autocomplete
// @Param name query string false "name of searched restaurant"
// @Param address query string false "address of searched restaurant"
// @Accept  json
// @Produce  json
// @Success 200 {object} responseAutocompleteJSON
// @Failure 500 {string} []byte
// @Router /autocomplete [get]

// pcRestaurantsHandler godoc
// @Summary Returns restaurants around Prague College
// @Tags PC restaurants
// @Accept  json
// @Produce  json
// @Param radius query string false "Radius (in meters) of the area around a provided or pre-selected starting point. Restaurants in this area will be returned. Radius can be ignored when specified with radius=ignore and lat and lon parameters will no longer be required. When no radius is provided, a default value of 1000 meters is used."
// @Param cuisine query string false "Filters restaurants based on a list of cuisines, separated by commas -> cuisine=Czech,English. A restaurant will be returned only if it satisfies all provided cuisines.Available cuisines: American, Italian, Asian, Indian, Japanese, Vietnamese, Spanish, Mediterranean, French, Thai, Mexican, International, Czech, English, Balkan, Brazil, Russian, Chinese, Greek, Arabic, Korean."
// @Param price-range query string false "Filters restaurants based on a list of price ranges, separated by commas -> price-range=0-300,600-. A restaurant will be returned if it satisfies at least one provided price range. Available price ranges: 0-300,300-600,600-"
// @Param vegetarian query bool false "Filters out all non vegetarian restaurants."
// @Param vegan query bool false "Filters out all non vegan restaurants."
// @Param gluten-free query bool false "Filters out all non gluten free restaurants."
// @Param takeaway query bool false "Filters out all restaurants that don't have a takeaway option."
// @Param delivery-options query bool false "Filters out all restaurants that don't have a delivery option."
// @Success 200 {object} responseFullJSON
// @Failure 405 {object} responseSimpleJSON
// @Router /prague-college/restaurants [get]

// restaurantHandler godoc
// @Summary Provides info about a specific restaurant
// @Tags restaurant
// @Param restaurant-id path int true "Restaurant ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} responseFullJSON
// @Success 400 {object} responseSimpleJSON
// @Failure 500 {string} []byte
// @Router /restaurant/{restaurant-id} [get]

// restaurantsHandler godoc
// @Summary Returns restaurants based on queries
// @Tags restaurants
// @Accept  json
// @Produce  json
// @Param radius query string false "Radius (in meters) of the area around a provided or pre-selected starting point. Restaurants in this area will be returned. Radius can be ignored when specified with radius=ignore and lat and lon parameters will no longer be required. When no radius is provided, a default value of 1000 meters is used."
// @Param address query string false "Starting point for a search in a given radius."
// @Param lat query float64 false "Latitude in degrees. Lat is required if radius is not set to ignore."
// @Param lon query float64 false "Longitude in degrees. Lon is required if radius is not set to ignore."
// @Param cuisine query string false "Filters restaurants based on a list of cuisines, separated by commas -> cuisine=Czech,English. A restaurant will be returned only if it satisfies all provided cuisines.Available cuisines: American, Italian, Asian, Indian, Japanese, Vietnamese, Spanish, Mediterranean, French, Thai, Mexican, International, Czech, English, Balkan, Brazil, Russian, Chinese, Greek, Arabic, Korean."
// @Param price-range query string false "Filters restaurants based on a list of price ranges, separated by commas -> price-range=0-300,600-. A restaurant will be returned if it satisfies at least one provided price range. Available price ranges: 0-300,300-600,600-"
// @Param district query string false "Filters restaurants based on a list districts, separated by commas. A restaurant will be returned if it is in one of the provided districts. Example: district=Praha 1,Praha2"
// @Param vegetarian query bool false "Filters out all non vegetarian restaurants."
// @Param vegan query bool false "Filters out all non vegan restaurants."
// @Param gluten-free query bool false "Filters out all non gluten free restaurants."
// @Param takeaway query bool false "Filters out all restaurants that don't have a takeaway option."
// @Param delivery-options query bool false "Filters out all restaurants that don't have a delivery option."
// @Success 200 {object} responseFullJSON
// @Failure 400 {object} responseSimpleJSON
// @Failure 500 {string} []byte
// @Router /restaurants [get]

// registerHandler godoc
// @Summary Registers a user
// @Tags Register user
// @Accept  json
// @Produce  json
// @Param user body db.User true "Create a new user"
// @Success 200 {object} responseUserJSON
// @Success 400 {object} responseSimpleJSON
// @Failure 500 {string} []byte
// @Router /register [post]

// userGetHandler godoc
// @Summary Get info about a user
// @Description Returns a JSON with user info if the request headers contain an authenticated cookie.
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} responseUserJSON
// @Success 400 {object} responseSimpleJSON
// @Failure 500 {string} []byte
// @Router /user [get]

// userDeleteHandler godoc
// @Summary Deletes a user
// @Description Deletes a user if the request headers contain an authenticated cookie.
// @Tags Delete user
// @Accept  json
// @Produce  json
// @Success 200 {object} responseSimpleJSON
// @Success 400 {object} responseSimpleJSON
// @Failure 500 {string} []byte
// @Router /user [delete]
// loginHandler godoc

// userPatchHandler godoc
// @Summary Updates a user's password or username
// @Description Updates user's password or username based on the provided JSON. Only 1 field can be updated at a time. For password you need to provide "oldPassword" and "newPassword" fields, omitting the "newUsername" field and vice versa if you'd like to update the username
// @Tags Patch user
// @Param updateJSON body userUpdate true "Create a new user"
// @Accept  json
// @Produce  json
// @Success 200 {object} responseSimpleJSON
// @Success 400 {object} responseSimpleJSON
// @Failure 500 {string} []byte
// @Router /user [patch]

// @Summary Logs in a user
// @Description Logs in a user if the request headers contain an authenticated cookie.
// @Tags login
// @Accept  json
// @Produce  json
// @Param user body db.User true "Logs in a new user"
// @Success 200 {object} responseUserJSON
// @Success 400 {object} responseSimpleJSON
// @Failure 500 {string} []byte
// @Router /login [post]

var allowedEndpoints = [...]string{"/restaurants", "/prague-college/restaurants", "/autocomplete", "/user", "/login", "/register"}

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
	r.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	r.PathPrefix("/").HandlerFunc(catchAllHandler)
	log.Println("Starting server on", port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
