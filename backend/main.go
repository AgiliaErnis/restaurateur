package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

// MockRestaurant placeholder for real struct
// TODO: replace with Restaurant struct from scraper after it's merged
type MockRestaurant struct {
	Name string
}

type responseJSON struct {
	Status int
	Msg    string
	Data   []*MockRestaurant
}

var allowedEndpoints = [...]string{"/restaurants"}

// TODO: change this func to fetch actual data from database
// after scraper and db is merged to main
func getFakeRestaurants(lat, lon float64, radius int) []*MockRestaurant {
	log.Printf("Lat: %v, Lon: %v, Rad: %v", lat, lon, radius)
	return []*MockRestaurant{{"Restaurant1"}, {"Restaurant2"}}
}

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
		radius, errRad := strconv.Atoi(radiusParam)
		if errRad != nil {
			radius = 1000
		}
		fakeRestaurants := getFakeRestaurants(lat, lon, radius)
		res.Status = resStatus
		res.Msg = "Success"
		res.Data = fakeRestaurants
	}
	response, err := json.Marshal(res)
	if err != nil {
		resStatus = http.StatusInternalServerError
		response = []byte("Internal server error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(resStatus)
	w.Write(response)
}

func main() {
	r := mux.NewRouter()
	port := ":8080"
	r.HandleFunc("/restaurants", restaurants).Methods(http.MethodGet)
	r.PathPrefix("/").HandlerFunc(catchAllHandler)
	log.Println("Starting server on", port)
	log.Fatal(http.ListenAndServe(port, r))
}
