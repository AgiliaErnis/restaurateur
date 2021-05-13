package api

import (
	"github.com/AgiliaErnis/restaurateur/backend/internal/db"
	"log"
	"net/http"
	"net/url"
)

type restaurantAutocomplete struct {
	ID       int    `json:"ID" example:"1"`
	Name     string `json:"Name" example:"Steakhouse"`
	Address  string `json:"Address" example:"Polská 13"`
	District string `json:"District" example:"Praha 1"`
	Image    string `json:"Image" example:"url.com"`
}

func getAutocompleteCandidates(params url.Values) ([]*restaurantAutocomplete, error) {
	pgQuery := ""
	input := ""
	_, name := params["name"]
	_, address := params["address"]
	if name {
		pgQuery = "SELECT id, name, address, district, coalesce(substring(images, '(?<=\")\\S+?(?=\")'), '') as image FROM restaurants WHERE " +
			"(unaccent(name) % unaccent($1))" +
			" ORDER BY SIMILARITY(unaccent(name), unaccent($1)) DESC"
		input = params.Get("name")
	} else if address {
		pgQuery = "SELECT id, name, address, district, coalesce(substring(images, '(?<=\")\\S+?(?=\")'), '') as image FROM restaurants WHERE " +
			"(regexp_replace(unaccent(address), '[[:digit:]/]', '', 'g') % unaccent($1)) " +
			"ORDER BY SIMILARITY(unaccent(address), unaccent($1)) DESC"
		input = params.Get("address")
	}
	var restaurants []*restaurantAutocomplete
	conn, err := db.GetConn()
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
		res := &responseAutocompleteJSON{}
		writeResponse(w, http.StatusInternalServerError, res)
		return
	}
	res := &responseAutocompleteJSON{
		Msg:  "Success",
		Data: autocompletedRestaurants,
	}
	writeResponse(w, http.StatusOK, res)
}
