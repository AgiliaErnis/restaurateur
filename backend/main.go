package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type responseJSON struct {
	Status int
	Msg    string
	Data   []RestaurantDB
}

var allowedEndpoints = [...]string{"/restaurants"}

func logRequest(r *http.Request, handlerName string) {
	method := r.Method
	endpoint := r.URL
	clientAddr := r.RemoteAddr
	headers, _ := json.Marshal(r.Header)
	log.Printf("Incoming request\n"+
		"Method: %q\n"+
		"Client's address: %q\n"+
		"Request URL: %q\n"+
		"Request headers: %s\n"+
		"Handler: %q\n",
		method, clientAddr, endpoint, headers, handlerName)
}

func catchAllHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	logRequest(r, "catchAllHandler")
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

func pcRestaurantsHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "pcRestaurantsHandler")
	pcLat := 50.0785714
	pcLon := 14.4400922
	v := r.URL.Query()
	radiusParam := v.Get("radius")
	radius, errRad := strconv.ParseFloat(radiusParam, 64)
	if errRad != nil {
		radius = 300
	}
	conn, _ := dbInitialise()
	// Null is sometimes "null" sometimes null
	loadedRestaurants, _ := loadRestaurants(conn)
	filteredRestaurants := getRestaurantsInRadius(loadedRestaurants, pcLat, pcLon, radius)
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

func restaurantsHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "restaurantsHandler")
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
		conn, _ := dbInitialise()
		loadedRestaurants, _ := loadRestaurants(conn)
		filteredRestaurants := getRestaurantsInRadius(loadedRestaurants, lat, lon, radius)
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
	r.HandleFunc("/prague-college/restaurants", pcRestaurantsHandler).Methods(http.MethodGet)
	r.HandleFunc("/restaurants", restaurantsHandler).Methods(http.MethodGet)
	r.PathPrefix("/").HandlerFunc(catchAllHandler)
	log.Println("Starting server on", port)
	log.Fatal(http.ListenAndServe(port, r))
}
