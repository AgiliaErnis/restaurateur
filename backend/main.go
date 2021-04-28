package main

import (
	"encoding/json"
	"fmt"
	"github.com/AgiliaErnis/restaurateur/backend/scraper"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type responseJSON struct {
	Status int
	Msg    string
	Data   []interface{}
}

type restaurantAutocomplete struct {
	ID       int
	Name     string
	Address  string
	District string
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

func getRestaurantDBInterfaces(restaurants []*RestaurantDB) []interface{} {
	interfaces := make([]interface{}, len(restaurants))
	for i, v := range restaurants {
		interfaces[i] = v
	}
	return interfaces
}

func getAutocompleteInterfaces(restaurants []*restaurantAutocomplete) []interface{} {
	interfaces := make([]interface{}, len(restaurants))
	for i, v := range restaurants {
		interfaces[i] = v
	}
	return interfaces
}

func getAutocompleteCandidates(params url.Values) ([]*restaurantAutocomplete, error) {
	pgQuery := ""
	input := ""
	_, name := params["name"]
	_, address := params["address"]
	if name {
		pgQuery = "SELECT id, name, address, district FROM restaurants WHERE " +
			"(unaccent(name) % unaccent($1))" +
			" ORDER BY SIMILARITY(unaccent(name), unaccent($1)) DESC"
		input = params.Get("name")
	} else if address {
		pgQuery = "SELECT id, name, address, district FROM restaurants WHERE " +
			"(regexp_replace(unaccent(address), '[[:digit:]/]', '', 'g') % unaccent($1)) " +
			"ORDER BY SIMILARITY(unaccent(address), unaccent($1)) DESC"
		input = params.Get("address")
	}
	var restaurants []*restaurantAutocomplete
	conn, err := dbInitialise()
	if err != nil {
		return restaurants, err
	}
	inputLen := len(input)
	if inputLen < 3 {
		_, err = conn.Exec("SELECT set_limit(0.1)")
	} else if inputLen < 5 {
		_, err = conn.Exec("SELECT set_limit(0.2)")
	} // else keep default of 0.3
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
	var restaurants []*RestaurantDB
	conn, err := dbInitialise()
	if err != nil {
		return restaurants, err
	}
	var andParams = [...]string{"vegetarian", "vegan", "gluten-free", "takeaway"}
	var nullParams = [...]string{"delivery-options"}
	var queries []string
	var orderBy = ""
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
	parameterArr, ok := params["search-name"]
	pgParam := fmt.Sprintf("(unaccent(name) %% unaccent($%d))", paramCtr)
	searchField := "name"
	if !ok {
		parameterArr, ok = params["search-address"]
		searchField = "address"
		pgParam = fmt.Sprintf("(regexp_replace(unaccent(address), '[[:digit:]/]', '', 'g') %% unaccent($%d)) ", paramCtr)
	}
	if ok {
		searchString := parameterArr[0]
		queries = append(queries, pgParam)
		values = append(values, searchString)
		orderBy = fmt.Sprintf(" ORDER BY SIMILARITY(unaccent(%s), unaccent($%d)) DESC", searchField, paramCtr)
		searchStringLen := len(searchString)
		if searchStringLen < 3 {
			_, err = conn.Exec("SELECT set_limit(0.1)")
		} else if searchStringLen < 5 {
			_, err = conn.Exec("SELECT set_limit(0.2)")
		} // else keep default of 0.3
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
	for _, r := range restaurants {
		if r.isInRadius(lat, lon, radiusParam) {
			if r.hasCuisines(cuisineParam) && r.isInPriceRange(priceRangeParam) && r.isInDistrict(districtParam) {
				filteredRestaurants = append(filteredRestaurants, r)
			}
		}
	}
	return filteredRestaurants
}

func writeResponse(w http.ResponseWriter, status int, response responseJSON) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	response.Status = status
	res, err := json.Marshal(response)
	if err != nil {
		log.Println("Error while marshalling JSON response")
		status = http.StatusInternalServerError
	}
	if status == http.StatusInternalServerError || err != nil {
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			panic(err)
		}
	} else {
		_, err := w.Write(res)
		if err != nil {
			panic(err)
		}
	}
}

func autocompleteHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "autocompleteHandler")
	params := r.URL.Query()
	autocompletedRestaurants, err := getAutocompleteCandidates(params)
	if err != nil {
		log.Println("Couldn't load restaurants from db")
		log.Println(err)
		writeResponse(w, http.StatusInternalServerError, responseJSON{})
		return
	}
	res := responseJSON{
		Msg:  "Success",
		Data: getAutocompleteInterfaces(autocompletedRestaurants),
	}
	writeResponse(w, http.StatusOK, res)
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
		writeResponse(w, http.StatusMethodNotAllowed, res)
		return
	}
	res.Msg = fmt.Sprintf("Invalid endpoint: %v", path)
	writeResponse(w, http.StatusBadRequest, res)
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
		writeResponse(w, http.StatusInternalServerError, responseJSON{})
		return
	}
	filteredRestaurants := filterRestaurants(loadedRestaurants, params, pcLat, pcLon)
	res := responseJSON{
		Msg:  "Success",
		Data: getRestaurantDBInterfaces(filteredRestaurants),
	}
	writeResponse(w, http.StatusOK, res)
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
		writeResponse(w, http.StatusBadRequest, res)
		return
	}
	loadedRestaurants, err := getDBRestaurants(params)
	if err != nil {
		log.Println("Couldn't load restaurants from db")
		writeResponse(w, http.StatusInternalServerError, responseJSON{})
		return
	}
	filteredRestaurants := filterRestaurants(loadedRestaurants, params, lat, lon)
	res.Msg = "Success"
	res.Data = getRestaurantDBInterfaces(filteredRestaurants)
	if err != nil {
		log.Println("Database not initialized")
		writeResponse(w, http.StatusInternalServerError, responseJSON{})
		return
	}
	writeResponse(w, http.StatusOK, res)
}

func main() {
	args := os.Args[1:]
	if scraper.SliceContains(args, "--initialize") {
		conn, err := dbInitialise()
		if err != nil {
			log.Println(err)
			log.Fatal("Make sure the DB_DSN environment variable is set")
		} else {
			log.Println("Connection to postgres established, downloading data...")
		}
		err = storeRestaurants(conn)
		if err != nil {
			log.Fatal(err)
		}
	}
	r := mux.NewRouter()
	port := ":8080"
	r.HandleFunc("/prague-college/restaurants", pcRestaurantsHandler).Methods(http.MethodGet)
	r.HandleFunc("/restaurants", restaurantsHandler).Methods(http.MethodGet)
	r.HandleFunc("/autocomplete", autocompleteHandler).Methods(http.MethodGet)
	r.PathPrefix("/").HandlerFunc(catchAllHandler)
	log.Println("Starting server on", port)
	log.Fatal(http.ListenAndServe(port, r))
}
