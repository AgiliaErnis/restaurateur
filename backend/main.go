package main

import (
	"encoding/json"
	"fmt"
	"github.com/AgiliaErnis/restaurateur/backend/scraper"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type responseJSON struct {
	Status int
	Msg    string
	Data   []*scraper.Restaurant
}

var allowedEndpoints = [...]string{"/restaurants"}

// Placeholder data before data from the database is available
var vinohradyRestaurants, restaurantErr = scraper.GetRestaurants("vinohrady")

func catchAllHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	res := responseJSON{}
	resStatus := http.StatusNotFound
	found := false
	for _, endpoint := range allowedEndpoints {
		if endpoint == path {
			found = true
			break
		}
	}
	if found {
		resStatus = http.StatusMethodNotAllowed
		res.Status = resStatus
		res.Msg = fmt.Sprintf("Wrong method: %v", r.Method)
	} else {
		res.Status = resStatus
		res.Msg = fmt.Sprintf("Invalid endpoint: %v", path)
	}
	response, err := json.Marshal(res)
	if err != nil {
		resStatus = http.StatusInternalServerError
		response = []byte("Internal server error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resStatus)
	w.Write(response)
}

func pcRestaurants(w http.ResponseWriter, r *http.Request) {
	pcLat := 50.0785714
	pcLon := 14.4400922
	v := r.URL.Query()
	radiusParam := v.Get("radius")
	radius, errRad := strconv.ParseFloat(radiusParam, 64)
	if errRad != nil {
		radius = 5000
	}
	filteredRestaurants := getRestaurantsInRadius(vinohradyRestaurants, pcLat, pcLon, radius)
	resStatus := http.StatusOK
	res := responseJSON{
		Status: http.StatusOK,
		Msg:    "Success",
		Data:   filteredRestaurants,
	}
	response, err := json.Marshal(res)
	if err != nil {
		resStatus = http.StatusInternalServerError
		response = []byte("Internal server error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resStatus)
	w.Write(response)
}

func restaurants(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	latParam := v.Get("lat")
	lonParam := v.Get("lon")
	lat, errLat := strconv.ParseFloat(latParam, 64)
	lon, errLon := strconv.ParseFloat(lonParam, 64)
	resStatus := http.StatusOK
	res := responseJSON{}
	if errLat != nil || errLon != nil {
		resStatus = http.StatusBadRequest
		res.Status = resStatus
		res.Msg = fmt.Sprintf("Invalid coordinates(Lat: %s, Lon: %s)", latParam, lonParam)
	} else {
		radiusParam := v.Get("radius")
		radius, errRad := strconv.ParseFloat(radiusParam, 64)
		if errRad != nil {
			radius = 1000
		}
		// TODO: change this func to fetch actual data from database
		filteredRestaurants := getRestaurantsInRadius(vinohradyRestaurants, lat, lon, radius)
		res.Status = resStatus
		res.Msg = "Success"
		res.Data = filteredRestaurants
	}
	response, err := json.Marshal(res)
	if err != nil {
		resStatus = http.StatusInternalServerError
		response = []byte("Internal server error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resStatus)
	w.Write(response)
}

func main() {
	r := mux.NewRouter()
	port := ":8080"
	r.HandleFunc("/prague-college/restaurants", pcRestaurants).Methods(http.MethodGet)
	r.HandleFunc("/restaurants", restaurants).Methods(http.MethodGet)
	r.PathPrefix("/").HandlerFunc(catchAllHandler)
	log.Println("Starting server on", port)
	log.Fatal(http.ListenAndServe(port, r))
}
