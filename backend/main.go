package main

import (
	"encoding/json"
	"fmt"
	"github.com/AgiliaErnis/restaurateur/backend/scraper"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

func getDBRestaurants(params url.Values) ([]*RestaurantDB, error) {
	var andParams = [...]string{"vegetarian", "vegan", "gluten-free", "takeaway"}
	var queries []string
	pgQuery := "SELECT * from restaurants"
	paramCtr := 1
	var values []interface{}
	for _, param := range andParams {
		_, ok := params[param]
		if ok {
			param = strings.Replace(param, "-", "_", -1)
			pgParam := fmt.Sprintf("%s=$%d", param, paramCtr)
			queries = append(queries, pgParam)
			value := params.Get(param)
			if value == "" {
				values = append(values, true)
			} else {
				values = append(values, value)
			}
			paramCtr++
		}
	}
	if len(queries) > 0 {
		pgQuery += " WHERE "
	}
	pgQuery += strings.Join(queries, " AND ")
	var restaurants []*RestaurantDB
	conn, err := dbInitialise()
	if err != nil {
		return restaurants, err
	}
	err = conn.Select(&restaurants, pgQuery, values...)
	if err != nil {
		return restaurants, err
	}
	return restaurants, nil
}

func filterRestaurants(restaurants []*RestaurantDB, params url.Values, lat, lon float64) []*RestaurantDB {
	var filteredRestaurants []*RestaurantDB
	radiusParam := params.Get("radius")
	cuisineParam := params.Get("cuisine")
	priceRangeParam := params.Get("price-range")
	districtParam := params.Get("district")
	radius, errRad := strconv.ParseFloat(radiusParam, 64)
	if errRad != nil {
		// default value
		radius = 1000
	}
	for _, r := range restaurants {
		if radiusParam == "ignore" || r.isInRadius(lat, lon, radius) {
			if r.hasCuisines(cuisineParam) && r.isInPriceRange(priceRangeParam) && r.isInDistrict(districtParam) {
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
	params := r.URL.Query()
	pcLat := 50.0785714
	pcLon := 14.4400922
	// Null is sometimes "null" sometimes null
	loadedRestaurants, err := getDBRestaurants(params)
	if err != nil {
		log.Println("Couldn't load restaurants from db")
		log.Println(err)
		prepareResponse(w, http.StatusInternalServerError, responseJSON{})
		return
	}
	filteredRestaurants := filterRestaurants(loadedRestaurants, params, pcLat, pcLon)
	res := responseJSON{
		Msg:  "Success",
		Data: filteredRestaurants,
	}
	prepareResponse(w, http.StatusOK, res)
}

func getCoordinates(params url.Values) (float64, float64, error) {
	addressParam := params.Get("address")
	radiusParam := params.Get("radius")
	if radiusParam == "ignore" {
		return 0, 0, nil
	} else if addressParam != "" {
		var lat float64
		var lon float64
		i := strings.Index(addressParam, "Praha")
		var address string
		var city string
		if i > -1 {
			address = addressParam[:i]
			city = addressParam[i:]
		} else {
			address = addressParam
			city = "Praha"
		}
		nominatim, err := scraper.GetNominatimJSON(address, city)
		if len(nominatim) == 0 || err != nil {
			return 0, 0, fmt.Errorf("Couldn't get coordinates for %q", addressParam)
		}
		lat, errLat := strconv.ParseFloat(nominatim[0].Lat, 64)
		lon, errLon := strconv.ParseFloat(nominatim[0].Lon, 64)
		if errLat != nil || errLon != nil {
			return 0, 0, fmt.Errorf("Coordinates for %s not found", addressParam)
		}
		return lat, lon, nil
	}
	latParam := params.Get("lat")
	lonParam := params.Get("lon")
	lat, errLat := strconv.ParseFloat(latParam, 64)
	lon, errLon := strconv.ParseFloat(lonParam, 64)
	if errLat != nil || errLon != nil {
		return 0, 0, fmt.Errorf("Invalid coordinates(Lat: %s, Lon: %s)", latParam, lonParam)
	}
	return lat, lon, nil
}

func restaurantsHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "restaurantsHandler")
	params := r.URL.Query()
	lat, lon, err := getCoordinates(params)
	res := responseJSON{}
	if err != nil {
		res.Msg = fmt.Sprintf("%s", err)
		prepareResponse(w, http.StatusBadRequest, res)
		return
	}
	loadedRestaurants, err := getDBRestaurants(params)
	if err != nil {
		log.Println("Couldn't load restaurants from db")
		prepareResponse(w, http.StatusInternalServerError, responseJSON{})
		return
	}
	filteredRestaurants := filterRestaurants(loadedRestaurants, params, lat, lon)
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
