package api

import (
	"encoding/json"
	"github.com/AgiliaErnis/restaurateur/backend/internal/db"
	"log"
	"net/http"
)

type response interface {
	setStatus(status int)
}

type restaurantIDJSON struct {
	RestaurantID int `json:"restaurantID"`
}

type userResponseFull struct {
	Name             string             `json:"name" example:"name"`
	Email            string             `json:"email" example:"test@mail.com`
	SavedRestaurants []*db.RestaurantDB `json:"savedRestaurants"`
}

type userResponseSimple struct {
	Name                string `json:"name" example:"name"`
	Email               string `json:"email" example:"test@mail.com`
	SavedRestaurantsIDs []int  `json:"savedRestaurantsIDs" example:"1,2"`
}

type responseFullJSON struct {
	Status int                 `json:"status" example:"200"`
	Msg    string              `json:"msg" example:"Success"`
	Data   []*db.RestaurantDB  `json:"data"`
	User   *userResponseSimple `json:"user"`
}

type responseSimpleJSON struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

type responseUserJSON struct {
	Status int               `json:"status" example:"200"`
	Msg    string            `json:"msg" example:"Success"`
	User   *userResponseFull `json:"user"`
}

type responseAutocompleteJSON struct {
	Status int                       `json:"status" example:"200"`
	Msg    string                    `json:"msg" example:"Success"`
	Data   []*restaurantAutocomplete `json:"data"`
}

func (r *responseAutocompleteJSON) setStatus(status int) {
	r.Status = status
}

func (r *responseUserJSON) setStatus(status int) {
	r.Status = status
}

func (r *responseFullJSON) setStatus(status int) {
	r.Status = status
}

func (r *responseSimpleJSON) setStatus(status int) {
	r.Status = status
}

func writeResponse(w http.ResponseWriter, status int, res response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	res.setStatus(status)
	r, err := json.Marshal(res)
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
		_, err := w.Write(r)
		if err != nil {
			panic(err)
		}
	}
}
