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
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/swaggo/http-swagger"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"html"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	// ORIGIN_ALLOWED is `scheme://dns[:port]`, or `*` (insecure)
	originsOk = handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	headersOk = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methodsOk = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	store     *sessions.CookieStore
)

type response interface {
	WriteResponse()
}

type user struct {
	Name     string `json:"username" validate:"required,min=2,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=64"`
}

type userResponse struct {
	Name  string
	Email string
}

type responseJSON struct {
	Status int             `json:"Status" example:"200"`
	Msg    string          `json:"Msg" example:"Success"`
	Data   []*RestaurantDB `json:"Data"`
	User   *userResponse   `json:"User"`
}

type responseJSON struct {
	Status int             `json:"Status" example:"200"`
	Msg    string          `json:"Msg" example:"Success"`
	Data   []*RestaurantDB `json:"Data"`
	User   *userResponse   `json:"User"`
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
		_, err = w.Write([]byte("Internal server error"))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		_, err = w.Write(res)
		if err != nil {
			log.Fatal(err)
		}
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
		_, err = w.Write([]byte("Internal server error"))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		_, err = w.Write(res)
		if err != nil {
			log.Fatal(err)
		}
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
		_, err = w.Write([]byte("Internal server error"))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		_, err = w.Write(res)
		if err != nil {
			log.Fatal(err)
		}
	}
}

type restaurantAutocomplete struct {
	ID       int    `json:"ID" example:"1"`
	Name     string `json:"Name" example:"Steakhouse"`
	Address  string `json:"Address" example:"Polská 13"`
	District string `json:"District" example:"Praha 1"`
	Image    string `json:"Image" example:"url.com"`
}

var allowedEndpoints = [...]string{"/restaurants", "/prague-college/restaurants", "/autocomplete", "/user", "/login", "/register"}

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

func getAutocompleteCandidates(params url.Values) ([]*restaurantAutocomplete, error) {
	pgQuery := ""
	input := ""
	_, name := params["name"]
	_, address := params["address"]
	if name {
		pgQuery = "SELECT id, name, address, district, coalesce(substring(images, '(?<=\")\\S+?(?=\")'), '') as image FROM restaurants_bak WHERE " +
			"(unaccent(name) % unaccent($1))" +
			" ORDER BY SIMILARITY(unaccent(name), unaccent($1)) DESC"
		input = params.Get("name")
	} else if address {
		pgQuery = "SELECT id, name, address, district, coalesce(substring(images, '(?<=\")\\S+?(?=\")'), '') as image FROM restaurants_bak WHERE " +
			"(regexp_replace(unaccent(address), '[[:digit:]/]', '', 'g') % unaccent($1)) " +
			"ORDER BY SIMILARITY(unaccent(address), unaccent($1)) DESC"
		input = params.Get("address")
	}
	var restaurants []*restaurantAutocomplete
	conn, err := dbGetConn()
	if err != nil {
		return restaurants, err
	}
	defer conn.Close()
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

func getRestaurantArrByID(id int) ([]*RestaurantDB, error) {
	var restaurant []*RestaurantDB
	queryString := "SELECT * FROM restaurants_bak where id=$1"
	conn, err := dbGetConn()
	if err != nil {
		return restaurant, err
	}
	err = conn.Select(&restaurant, queryString, id)
	if err != nil {
		return restaurant, err
	}
	return restaurant, nil
}

func getDBRestaurants(params url.Values) ([]*RestaurantDB, error) {
	var restaurants []*RestaurantDB
	conn, err := dbGetConn()
	if err != nil {
		return restaurants, err
	}
	defer conn.Close()
	var andParams = [...]string{"vegetarian", "vegan", "gluten-free", "takeaway"}
	var nullParams = [...]string{"delivery-options"}
	var queries []string
	var orderBy = ""
	pgQuery := "SELECT * from restaurants_bak"
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
	autocompletedRestaurants, err := getAutocompleteCandidates(params)
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
// @Param radius query string false "Radius (in meters) of the area around a provided or pre-selected starting point. Restaurants in this area will be returned. Radius can be ignored when specified with radius=ignore and lat and lon parameters will no longer be required. When no radius is provided, a default value of 1000 meters is used."
// @Param cuisine query string false "Filters restaurants based on a list of cuisines, separated by commas -> cuisine=Czech,English. A restaurant will be returned only if it satisfies all provided cuisines.Available cuisines: American, Italian, Asian, Indian, Japanese, Vietnamese, Spanish, Mediterranean, French, Thai, Mexican, International, Czech, English, Balkan, Brazil, Russian, Chinese, Greek, Arabic, Korean."
// @Param price-range query string false "Filters restaurants based on a list of price ranges, separated by commas -> price-range=0-300,600-. A restaurant will be returned if it satisfies at least one provided price range. Available price ranges: 0-300,300-600,600-"
// @Param vegetarian query bool false "Filters out all non vegetarian restaurants."
// @Param vegan query bool false "Filters out all non vegan restaurants."
// @Param gluten-free query bool false "Filters out all non gluten free restaurants."
// @Param takeaway query bool false "Filters out all restaurants that don't have a takeaway option."
// @Param delivery-options query bool false "Filters out all restaurants that don't have a delivery option."
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
		res := responseErrorJSON{}
		res.WriteResponse(w, http.StatusInternalServerError)
		return
	}
	filteredRestaurants := filterRestaurants(loadedRestaurants, params, pcLat, pcLon)
	res := responseJSON{
		Msg:  "Success",
		Data: filteredRestaurants,
	}
	auth, id := isAuthenticated(w, r)
	if auth {
		user, _ := getUserByID(id)
		res.User = &userResponse{Name: user.Name, Email: user.Email}
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

// restaurantHandler godoc
// @Summary Provides info about a specific restaurant
// @Tags restaurant
// @Param restaurant-id path int true "Restaurant ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} responseJSON
// @Success 400 {object} responseErrorJSON
// @Failure 500 {string} []byte
// @Router /restaurant/{restaurant-id} [get]
func restaurantHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "restaurantHandler")
	res := responseJSON{}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
		res.WriteResponse(w, http.StatusInternalServerError)
		return
	}
	restaurant, err := getRestaurantArrByID(id)
	if err != nil {
		log.Println(err)
		res.WriteResponse(w, http.StatusInternalServerError)
		return
	}
	res.Data = restaurant
	if len(restaurant) != 1 {
		resErr := responseErrorJSON{}
		resErr.Msg = fmt.Sprintf("ID number %d not found in database", id)
		resErr.WriteResponse(w, http.StatusBadRequest)
		return
	}
	auth, id := isAuthenticated(w, r)
	if auth {
		user, _ := getUserByID(id)
		res.User = &userResponse{Name: user.Name, Email: user.Email}
	}
	res.Msg = "Success"
	res.WriteResponse(w, http.StatusOK)
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
		resErr := responseErrorJSON{}
		resErr.Msg = fmt.Sprintf("%s", err)
		resErr.WriteResponse(w, http.StatusBadRequest)
		return
	}
	loadedRestaurants, err := getDBRestaurants(params)
	if err != nil {
		log.Println("Couldn't load restaurants from db")
		res := responseErrorJSON{}
		res.WriteResponse(w, http.StatusInternalServerError)
		return
	}
	filteredRestaurants := filterRestaurants(loadedRestaurants, params, lat, lon)
	res.Msg = "Success"
	res.Data = filteredRestaurants
	auth, id := isAuthenticated(w, r)
	if auth {
		user, _ := getUserByID(id)
		res.User = &userResponse{Name: user.Name, Email: user.Email}
	}
	if err != nil {
		log.Println("Database not initialized")
		res := responseErrorJSON{}
		res.WriteResponse(w, http.StatusInternalServerError)
		return
	}
	res.WriteResponse(w, http.StatusOK)
}

func isAuthenticated(w http.ResponseWriter, r *http.Request) (bool, int) {
	session, err := store.Get(r, "session-id")
	if err != nil {
		log.Println(err)
		return false, 0
	}
	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth ||
		!(session.Values["expires"].(int64) > time.Now().Unix()) {
		return false, 0
	}
	// add 15 min to cookie
	session.Values["expires"] = time.Now().Add(time.Minute * 15).Unix()
	session.Options.MaxAge = 60 * 15
	err = session.Save(r, w)
	if err != nil {
		log.Println(err)
		return false, 0
	}
	return true, session.Values["user-id"].(int)
}

func init() {
	authKey := securecookie.GenerateRandomKey(64)
	encryptionKey := securecookie.GenerateRandomKey(32)
	store = sessions.NewCookieStore(
		authKey,
		encryptionKey,
	)

	store.Options = &sessions.Options{
		MaxAge:   60 * 15, // 15 min
		HttpOnly: true,
	}
}

// registerHandler godoc
// @Summary Registers a user
// @Tags register
// @Accept  json
// @Produce  json
// @Param user body user true "Create a new user"
// @Success 200 {object} responseJSON
// @Success 400 {object} responseErrorJSON
// @Failure 500 {string} []byte
// @Router /register [post]
func registerHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "registerHandler")
	user := &user{}
	res := responseJSON{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		log.Println(err)
		res.WriteResponse(w, http.StatusInternalServerError)
		return
	}
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		log.Println(err)
		res.Msg = "Credentials do not comply with requirements"
		res.WriteResponse(w, http.StatusBadRequest)
		return
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		res.WriteResponse(w, http.StatusInternalServerError)
		return
	}
	user.Password = string(pass)
	user.Name = html.EscapeString(user.Name)
	err = saveUser(user)
	if err != nil {
		log.Println(err)
		if err.Error() == "pq: duplicate key value violates unique constraint \"restaurateur_users_email_key\"" {
			res.Msg = "Email already used"
			res.WriteResponse(w, http.StatusBadRequest)
			return
		}
		res.WriteResponse(w, http.StatusInternalServerError)
		return
	}
	res.Msg = "Registration successful!"
	res.WriteResponse(w, http.StatusOK)
}

// userHandler godoc
// @Summary Get info about a user
// @Description Returns a JSON with user info if the request headers contain an authenticated cookie.
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} responseJSON
// @Success 400 {object} responseErrorJSON
// @Failure 500 {string} []byte
// @Router /user [get]
func userHandler(w http.ResponseWriter, r *http.Request) {
	auth, id := isAuthenticated(w, r)
	if auth {
		res := responseJSON{}
		user, _ := getUserByID(id)
		res.User = &userResponse{Name: user.Name, Email: user.Email}
		res.Msg = "Success"
		res.WriteResponse(w, http.StatusOK)
		return
	}
	resErr := responseErrorJSON{}
	resErr.Msg = "Not authenticated"
	resErr.WriteResponse(w, http.StatusForbidden)
}

// userDeleteHandler godoc
// @Summary Deletes a user
// @Description Deletes a user if the request headers contain an authenticated cookie.
// @Tags userDelete
// @Accept  json
// @Produce  json
// @Success 200 {object} responseJSON
// @Success 400 {object} responseErrorJSON
// @Failure 500 {string} []byte
// @Router /user [delete]
func userDeleteHandler(w http.ResponseWriter, r *http.Request) {
	auth, id := isAuthenticated(w, r)
	if auth {
		res := responseJSON{}
		err := deleteUser(id)
		if err != nil {
			res.Msg = "Couldn't delete the record"
			res.WriteResponse(w, http.StatusInternalServerError)
			return
		}
		res.Msg = "Successfuly deleted the user!"
		res.WriteResponse(w, http.StatusOK)
		return
	}
	resErr := responseErrorJSON{}
	resErr.Msg = "Not authenticated"
	resErr.WriteResponse(w, http.StatusForbidden)
}

// loginHandler godoc
// @Summary Logs in a user
// @Description Logs in a user if the request headers contain an authenticated cookie.
// @Tags login
// @Accept  json
// @Produce  json
// @Param user body user true "Logs in a new user"
// @Success 200 {object} responseJSON
// @Success 400 {object} responseErrorJSON
// @Failure 500 {string} []byte
// @Router /login [post]
func loginHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "loginHandler")
	session, _ := store.Get(r, "session-id")
	user := &user{}
	res := responseJSON{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		log.Println(err)
		res.Msg = "Invalid JSON data"
		res.WriteResponse(w, http.StatusBadRequest)
		return
	}
	dbUser, err := getUserByEmail(user.Email)
	log.Println(err)
	log.Println(dbUser.Name)
	log.Println(dbUser.Email)
	log.Println(dbUser.Password)
	errf := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if errf != nil {
		resErr := responseErrorJSON{}
		resErr.Msg = "Invalid password"
		resErr.WriteResponse(w, http.StatusForbidden)
		return
	}
	session.Values["user-id"] = dbUser.ID
	session.Values["authenticated"] = true
	t := time.Now().Add(time.Minute * 15).Unix()
	session.Values["expires"] = t
	err = session.Save(r, w)
	if err != nil {
		log.Println(err)
		res.WriteResponse(w, http.StatusInternalServerError)
		return
	}
	res.Msg = "Log in successful!"
	res.User = &userResponse{Name: dbUser.Name, Email: dbUser.Email}
	res.WriteResponse(w, http.StatusOK)
}

func main() {
	var portNum int
	flag.IntVar(&portNum, "p", 8080, "Port number")
	flag.Parse()
	dbCheck()
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
	r.HandleFunc("/user", userHandler).Methods(http.MethodGet)
	r.HandleFunc("/user", userDeleteHandler).Methods(http.MethodDelete)
	r.HandleFunc("/login", loginHandler).Methods(http.MethodPost)
	r.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
	r.PathPrefix("/").HandlerFunc(catchAllHandler)
	log.Println("Starting server on", port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
