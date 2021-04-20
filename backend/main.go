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
	Data   []*RestaurantDB
}

var allowedEndpoints = [...]string{"/restaurants", "/prague-college/restaurants"}

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

func filterRestaurants(restaurants []*RestaurantDB, req *http.Request, lat, lon float64) []*RestaurantDB {
	var filteredRestaurants []*RestaurantDB
	params := req.URL.Query()
	radiusParam := params.Get("radius")
	cuisinesParam := params.Get("cuisines")
	radius, errRad := strconv.ParseFloat(radiusParam, 64)
	if errRad != nil {
		// default value
		radius = 500
	}
	_, vegan := params["vegan"]
	_, vegetarian := params["vegetarian"]
	_, glutenFree := params["glutenfree"]
	_, takeaway := params["takeaway"]
	for _, r := range restaurants {
		if radiusParam == "all" || r.isInRadius(lat, lon, radius) {
			if (vegan && !r.Vegan) || (vegetarian && !r.Vegetarian) ||
				(glutenFree && !r.GlutenFree) || (takeaway && !r.Takeaway) || !r.hasCuisines(cuisinesParam) {
				continue
			} else {
				filteredRestaurants = append(filteredRestaurants, r)
			}
		}
	}
	return filteredRestaurants
}

func prepareResponse(w http.ResponseWriter, status int, response responseJSON) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response.Status = status
	res, err := json.Marshal(response)
	if err != nil {
		log.Println("Error while marshalling JSON response")
	}
	if status == http.StatusInternalServerError || err != nil {
		w.Write([]byte("Internal server error"))
	} else {
		w.Write(res)
	}
}

func catchAllHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "catchAllHandler")
	path := r.URL.Path
	res := responseJSON{}
	found := false
	for _, endpoint := range allowedEndpoints {
		if endpoint == path {
			found = true
			break
		}
	}
	if found {
		res.Msg = fmt.Sprintf("Wrong method: %v", r.Method)
		prepareResponse(w, http.StatusMethodNotAllowed, res)
		return
	}
	res.Msg = fmt.Sprintf("Invalid endpoint: %v", path)
	prepareResponse(w, http.StatusBadRequest, res)
}

func pcRestaurantsHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "pcRestaurantsHandler")
	pcLat := 50.0785714
	pcLon := 14.4400922
	conn, err := dbInitialise()
	if err != nil {
		log.Println("Database not initialized")
		prepareResponse(w, http.StatusInternalServerError, responseJSON{})
		return
	}
	// Null is sometimes "null" sometimes null
	loadedRestaurants, err := loadRestaurants(conn)
	if err != nil {
		log.Println("Couldn't load restaurants from db")
		prepareResponse(w, http.StatusInternalServerError, responseJSON{})
		return
	}
	filteredRestaurants := filterRestaurants(loadedRestaurants, r, pcLat, pcLon)
	res := responseJSON{
		Msg:  "Success",
		Data: filteredRestaurants,
	}
	prepareResponse(w, http.StatusOK, res)
}

func restaurantsHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "restaurantsHandler")
	v := r.URL.Query()
	latParam := v.Get("lat")
	lonParam := v.Get("lon")
	radiusParam := v.Get("radius")
	lat, errLat := strconv.ParseFloat(latParam, 64)
	lon, errLon := strconv.ParseFloat(lonParam, 64)
	res := responseJSON{}
	if radiusParam != "all" && (errLat != nil || errLon != nil) {
		res.Msg = fmt.Sprintf("Invalid coordinates(Lat: %s, Lon: %s)", latParam, lonParam)
		prepareResponse(w, http.StatusBadRequest, res)
		return
	}
	conn, err := dbInitialise()
	if err != nil {
		log.Println("Database not initialized")
		prepareResponse(w, http.StatusInternalServerError, responseJSON{})
		return
	}
	loadedRestaurants, err := loadRestaurants(conn)
	if err != nil {
		log.Println("Couldn't load restaurants from db")
		prepareResponse(w, http.StatusInternalServerError, responseJSON{})
		return
	}
	filteredRestaurants := filterRestaurants(loadedRestaurants, r, lat, lon)
	res.Msg = "Success"
	res.Data = filteredRestaurants
	if err != nil {
		log.Println("Database not initialized")
		prepareResponse(w, http.StatusInternalServerError, responseJSON{})
		return
	}
	prepareResponse(w, http.StatusOK, res)
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
