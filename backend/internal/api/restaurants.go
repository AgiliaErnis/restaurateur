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
// @Param has-menu query bool false "Filters out all restaurants that don't have a weekly menu."
// @Param sort query string false "Sorts restaurants. Available sort options: price-asc, price-desc, rating"
// @Success 200 {object} responseFullJSON
// @Failure 405 {object} responseSimpleJSON
// @Router /prague-college/restaurants [get]
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
	auth := isAuthenticated(w, r)
	if auth {
		id := getUserIDFromCookie(r)
		user, _ := db.GetUserByID(id)
		res.User = &userResponseSimple{Name: user.Name, Email: user.Email}
		savedRestaurants, _ := db.GetSavedRestaurantsID(id)
		res.User.SavedRestaurantsIDs = savedRestaurants
	}
	writeResponse(w, http.StatusOK, res)
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
// @Param has-menu query bool false "Filters out all restaurants that don't have a weekly menu."
// @Param sort query string false "Sorts restaurants. Available sort options: price-asc, price-desc, rating"
// @Success 200 {object} responseFullJSON
// @Failure 400 {object} responseSimpleJSON
// @Failure 500 {string} []byte
// @Router /restaurants [get]
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
	auth := isAuthenticated(w, r)
	if auth {
		id := getUserIDFromCookie(r)
		user, _ := db.GetUserByID(id)
		res.User = &userResponseSimple{Name: user.Name, Email: user.Email}
		savedRestaurants, _ := db.GetSavedRestaurantsID(id)
		res.User.SavedRestaurantsIDs = savedRestaurants
	}
	if err != nil {
		log.Println("Database not initialized")
		res := &responseSimpleJSON{}
		writeResponse(w, http.StatusInternalServerError, res)
		return
	}
	writeResponse(w, http.StatusOK, res)
}

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
	auth := isAuthenticated(w, r)
	if auth {
		id := getUserIDFromCookie(r)
		user, _ := db.GetUserByID(id)
		res.User = &userResponseSimple{Name: user.Name, Email: user.Email}
		savedRestaurants, _ := db.GetSavedRestaurantsID(id)
		res.User.SavedRestaurantsIDs = savedRestaurants
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
	_, shouldSort := params["sort"]
	if shouldSort {
		filteredRestaurants = sortRestaurants(filteredRestaurants, params.Get("sort"))
	}
	return filteredRestaurants
}

func sortRestaurants(restaurants []*db.RestaurantDB, sortOption string) []*db.RestaurantDB {
	switch sortOption {
	case "rating":
		restaurants = removeEmptyRatings(restaurants)
		db.SortBy(func(r1, r2 *db.RestaurantDB) bool {
			return r1.Rating > r2.Rating
		}).Sort(restaurants)
	case "price-asc":
		restaurants = removeEmptyPriceRanges(restaurants)
		db.SortBy(func(r1, r2 *db.RestaurantDB) bool {
			return r1.PriceRange < r2.PriceRange
		}).Sort(restaurants)
	case "price-desc":
		restaurants = removeEmptyPriceRanges(restaurants)
		db.SortBy(func(r1, r2 *db.RestaurantDB) bool {
			return r1.PriceRange > r2.PriceRange
		}).Sort(restaurants)
	}
	return restaurants
}

func removeEmptyRatings(restaurants []*db.RestaurantDB) []*db.RestaurantDB {
	var cleanedRestaurants []*db.RestaurantDB
	for _, restaurant := range restaurants {
		if restaurant.Rating != "" {
			cleanedRestaurants = append(cleanedRestaurants, restaurant)
		}
	}
	return cleanedRestaurants
}

func removeEmptyPriceRanges(restaurants []*db.RestaurantDB) []*db.RestaurantDB {
	var cleanedRestaurants []*db.RestaurantDB
	for _, restaurant := range restaurants {
		if restaurant.PriceRange != "Not available" {
			cleanedRestaurants = append(cleanedRestaurants, restaurant)
		}
	}
	return cleanedRestaurants
}
