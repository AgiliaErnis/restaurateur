// @title Restaurateur API
// @version 0.2.0
// @description Provides info about restaurants in Prague
// @host localhost:8080
// @BasePath /
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/AgiliaErnis/restaurateur/backend/docs"
	"github.com/AgiliaErnis/restaurateur/backend/scraper"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type response interface {
	WriteResponse()
}

type responseJSON struct {
	Status int             `json:"Status" example:"200"`
	Msg    string          `json:"Msg" example:"Success"`
	Data   []*RestaurantDB `json:"Data"`
}

type responseErrorJSON struct {
	Status int      `json:"Status"`
	Msg    string   `json:"Msg" example:"Error message"`
	Data   struct{} `json:"Data"`
}

type responseAutocompleteJSON struct {
	Status int                       `json:"Status" example:"200"`
	Msg    string                    `json:"Msg" example:"Success"`
	Data   []*restaurantAutocomplete `json:"Data"`
}

func (r *responseAutocompleteJSON) WriteResponse(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	r.Status = status
	res, err := json.Marshal(r)
	if err != nil {
		log.Println("Error while marshalling JSON response")
	}
	if status == http.StatusInternalServerError || err != nil {
		w.Write([]byte("Internal server error"))
	} else {
		w.Write(res)
	}
}

func (r *responseJSON) WriteResponse(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	r.Status = status
	res, err := json.Marshal(r)
	if err != nil {
		log.Println("Error while marshalling JSON response")
	}
	if status == http.StatusInternalServerError || err != nil {
		w.Write([]byte("Internal server error"))
	} else {
		w.Write(res)
	}
}

func (r *responseErrorJSON) WriteResponse(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	r.Status = status
	res, err := json.Marshal(r)
	if err != nil {
		log.Println("Error while marshalling JSON response")
	}
	if status == http.StatusInternalServerError || err != nil {
		w.Write([]byte("Internal server error"))
	} else {
		w.Write(res)
	}
}

type restaurantAutocomplete struct {
	ID       int    `json:"ID" example:"1"`
	Name     string `json:"Name" example:"Steakhouse"`
	District string `json:"District" example:"Praha 1"`
}

var allowedEndpoints = [...]string{"/restaurants", "/prague-college/restaurants", "/autocomplete"}

func logRequest(r *http.Request, handlerName string) {
	method := r.Method
	endpoint := r.URL
	clientAddr := r.RemoteAddr
	log.Printf("Incoming request\n"+
		"Method: %q\n"+
		"Client's address: %q\n"+
		"Request URL: %q\n"+
		"Handler: %q\n",
		method, clientAddr, endpoint, handlerName)
}

func getAutocompleteCandidates(input string) ([]*restaurantAutocomplete, error) {
	pgQuery := "SELECT id, name, district FROM restaurants WHERE " +
		"(unaccent(name) % unaccent($1))" +
		" ORDER BY SIMILARITY(unaccent(name), unaccent($1)) DESC"
	var restaurants []*restaurantAutocomplete
	conn, err := dbGetConn()
	defer conn.Close()
	if err != nil {
		return restaurants, err
	}
	err = conn.Select(&restaurants, pgQuery, input)
	if err != nil {
		return restaurants, err
	}
	if len(restaurants) > 10 {
		return restaurants[:10], nil
	}
	return restaurants, nil
}

func getDBRestaurants(params url.Values) ([]*RestaurantDB, error) {
	var andParams = [...]string{"vegetarian", "vegan", "gluten-free", "takeaway"}
	var nullParams = [...]string{"delivery-options"}
	var queries []string
	var orderBy = ""
	pgQuery := "SELECT * from restaurants"
	paramCtr := 1
	var values []interface{}
	for _, param := range andParams {
		_, ok := params[param]
		if ok && params.Get(param) != "false" {
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
	_, ok := params["search"]
	if ok {
		searchString := params.Get("search")
		pgParam := fmt.Sprintf("(unaccent(name) %% unaccent($%d))", paramCtr)
		queries = append(queries, pgParam)
		values = append(values, searchString)
		orderBy = fmt.Sprintf(" ORDER BY SIMILARITY(unaccent(name), unaccent($%d)) DESC", paramCtr)
		paramCtr++
	}
	for _, param := range nullParams {
		_, ok := params[param]
		if ok {
			param = strings.Replace(param, "-", "_", -1)
			pgParam := fmt.Sprintf("%s IS NOT NULL", param)
			queries = append(queries, pgParam)
		}
	}
	if len(queries) > 0 {
		pgQuery += " WHERE "
	}
	pgQuery += strings.Join(queries, " AND ") + orderBy
	var restaurants []*RestaurantDB
	conn, err := dbGetConn()
	defer conn.Close()
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
func autocompleteHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "autocompleteHandler")
	params := r.URL.Query()
	autocompletedRestaurants, err := getAutocompleteCandidates(params.Get("input"))
	if err != nil {
		log.Println("Couldn't load restaurants from db")
		log.Println(err)
		res := responseAutocompleteJSON{}
		res.WriteResponse(w, http.StatusInternalServerError)
		return
	}
	res := responseAutocompleteJSON{
		Msg:  "Success",
		Data: autocompletedRestaurants,
	}
	res.WriteResponse(w, http.StatusOK)
}

func catchAllHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "catchAllHandler")
	path := r.URL.Path
	res := responseErrorJSON{}
	found := false
	for _, endpoint := range allowedEndpoints {
		if endpoint == path {
			found = true
			break
		}
	}
	if found {
		res.Msg = fmt.Sprintf("Wrong method: %v", r.Method)
		res.WriteResponse(w, http.StatusMethodNotAllowed)
		return
	}
	res.Msg = fmt.Sprintf("Invalid endpoint: %v", path)
	res.WriteResponse(w, http.StatusBadRequest)
}

// pcRestaurantsHandler godoc
// @Summary Returns restaurants around Prague College
// @Tags PC restaurants
// @Accept  json
// @Produce  json
// @Success 200 {object} responseJSON
// @Failure 405 {object} responseErrorJSON
// @Router /prague-college/restaurants [get]
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
		res := responseJSON{}
		res.WriteResponse(w, http.StatusInternalServerError)
		return
	}
	filteredRestaurants := filterRestaurants(loadedRestaurants, params, pcLat, pcLon)
	res := responseJSON{
		Msg:  "Success",
		Data: filteredRestaurants,
	}
	res.WriteResponse(w, http.StatusOK)
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
// @Success 200 {object} responseJSON
// @Failure 400 {object} responseErrorJSON
// @Failure 500 {string} []byte
// @Router /restaurants [get]
func restaurantsHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "restaurantsHandler")
	params := r.URL.Query()
	lat, lon, err := getCoordinates(params)
	res := responseJSON{}
	if err != nil {
		res.Msg = fmt.Sprintf("%s", err)
		res.WriteResponse(w, http.StatusBadRequest)
		return
	}
	loadedRestaurants, err := getDBRestaurants(params)
	if err != nil {
		log.Println("Couldn't load restaurants from db")
		res := responseJSON{}
		res.WriteResponse(w, http.StatusInternalServerError)
		return
	}
	filteredRestaurants := filterRestaurants(loadedRestaurants, params, lat, lon)
	res.Msg = "Success"
	res.Data = filteredRestaurants
	if err != nil {
		log.Println("Database not initialized")
		res := responseJSON{}
		res.WriteResponse(w, http.StatusInternalServerError)
		return
	}
	res.WriteResponse(w, http.StatusOK)
}

func main() {
	var initialize bool
	flag.BoolVar(&initialize, "initialize", false, "Initializes the database and downloads restaurant data")
	var portNum int
	flag.IntVar(&portNum, "p", 8080, "Port number")
	flag.Parse()
	if initialize {
		dbInit()
	}
	if portNum < 1024 || portNum > 65535 {
		log.Fatal("Invalid port number, use a number from 1024-65535")
	}
	port := fmt.Sprintf(":%d", portNum)
	r := mux.NewRouter()
	r.HandleFunc("/prague-college/restaurants", pcRestaurantsHandler).Methods(http.MethodGet)
	r.HandleFunc("/restaurants", restaurantsHandler).Methods(http.MethodGet)
	r.HandleFunc("/autocomplete", autocompleteHandler).Methods(http.MethodGet)
	r.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	r.PathPrefix("/").HandlerFunc(catchAllHandler)
	log.Println("Starting server on", port)
	log.Fatal(http.ListenAndServe(port, r))
}
