package api

import (
	"fmt"
	"github.com/AgiliaErnis/restaurateur/backend/internal/db"
	"github.com/AgiliaErnis/restaurateur/backend/pkg/coordinates"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func pcRestaurantsHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "pcRestaurantsHandler")
	params := r.URL.Query()
	pcLat := 50.0785714
	pcLon := 14.4400922
	// Null is sometimes "null" sometimes null
	loadedRestaurants, err := db.GetDBRestaurants(params)
	if err != nil {
		log.Println("Couldn't load restaurants from db")
		log.Println(err)
		res := &responseSimpleJSON{}
		writeResponse(w, http.StatusInternalServerError, res)
		return
	}
	filteredRestaurants := filterRestaurants(loadedRestaurants, params, pcLat, pcLon)
	res := &responseFullJSON{
		Msg:  "Success",
		Data: filteredRestaurants,
	}
	auth, id := isAuthenticated(w, r)
	if auth {
		user, _ := db.GetUserByID(id)
		res.User = &userResponse{Name: user.Name, Email: user.Email}
	}
	writeResponse(w, http.StatusOK, res)
}

func restaurantsHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "restaurantsHandler")
	params := r.URL.Query()
	lat, lon, err := coordinates.GetCoordinates(params)
	res := &responseFullJSON{}
	if err != nil {
		resErr := &responseSimpleJSON{}
		resErr.Msg = fmt.Sprintf("%s", err)
		writeResponse(w, http.StatusBadRequest, resErr)
		return
	}
	loadedRestaurants, err := db.GetDBRestaurants(params)
	if err != nil {
		log.Println("Couldn't load restaurants from db")
		res := &responseSimpleJSON{}
		writeResponse(w, http.StatusInternalServerError, res)
		return
	}
	filteredRestaurants := filterRestaurants(loadedRestaurants, params, lat, lon)
	res.Msg = "Success"
	res.Data = filteredRestaurants
	auth, id := isAuthenticated(w, r)
	if auth {
		user, _ := db.GetUserByID(id)
		res.User = &userResponse{Name: user.Name, Email: user.Email}
	}
	if err != nil {
		log.Println("Database not initialized")
		res := &responseSimpleJSON{}
		writeResponse(w, http.StatusInternalServerError, res)
		return
	}
	writeResponse(w, http.StatusOK, res)
}

func restaurantHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "restaurantHandler")
	res := &responseFullJSON{}
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
		writeResponse(w, http.StatusInternalServerError, res)
		return
	}
	restaurant, err := db.GetRestaurantArrByID(id)
	if err != nil {
		log.Println(err)
		writeResponse(w, http.StatusInternalServerError, res)
		return
	}
	res.Data = restaurant
	if len(restaurant) != 1 {
		resErr := &responseSimpleJSON{}
		resErr.Msg = fmt.Sprintf("ID number %d not found in database", id)
		writeResponse(w, http.StatusBadRequest, resErr)
		return
	}
	auth, id := isAuthenticated(w, r)
	if auth {
		user, _ := db.GetUserByID(id)
		res.User = &userResponse{Name: user.Name, Email: user.Email}
	}
	res.Msg = "Success"
	writeResponse(w, http.StatusOK, res)
}

func filterRestaurants(restaurants []*db.RestaurantDB, params url.Values, lat, lon float64) []*db.RestaurantDB {
	var filteredRestaurants []*db.RestaurantDB
	radiusParam := params.Get("radius")
	cuisineParam := params.Get("cuisine")
	priceRangeParam := params.Get("price-range")
	districtParam := params.Get("district")
	for _, r := range restaurants {
		if r.IsInRadius(lat, lon, radiusParam) {
			if r.HasCuisines(cuisineParam) && r.IsInPriceRange(priceRangeParam) && r.IsInDistrict(districtParam) {
				filteredRestaurants = append(filteredRestaurants, r)
			}
		}
	}
	return filteredRestaurants
}
