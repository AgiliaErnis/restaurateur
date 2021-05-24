package api

import (
	"encoding/json"
	"github.com/AgiliaErnis/restaurateur/backend/internal/db"
	"log"
	"net/http"
)

type response interface {
	SetStatus(status int)
}

type restaurantIDJSON struct {
	RestaurantID int `json:"restaurantID"`
}

type userResponseFull struct {
	Name             string
	Email            string
	SavedRestaurants []*db.RestaurantDB
}

type userResponseSimple struct {
	Name                string
	Email               string
	SavedRestaurantsIDs []int
}

type responseFullJSON struct {
	Status int                 `json:"Status" example:"200"`
	Msg    string              `json:"Msg" example:"Success"`
	Data   []*db.RestaurantDB  `json:"Data"`
	User   *userResponseSimple `json:"User"`
}

type responseSimpleJSON struct {
	Status int    `json:"Status"`
	Msg    string `json:"Msg"`
}

type responseUserJSON struct {
	Status int               `json:"Status" example:"200"`
	Msg    string            `json:"Msg" example:"Success"`
	User   *userResponseFull `json:"User"`
}

type responseAutocompleteJSON struct {
	Status int                       `json:"Status" example:"200"`
	Msg    string                    `json:"Msg" example:"Success"`
	Data   []*restaurantAutocomplete `json:"Data"`
}

func (r *responseAutocompleteJSON) SetStatus(status int) {
	r.Status = status
}

func (r *responseUserJSON) SetStatus(status int) {
	r.Status = status
}

func (r *responseFullJSON) SetStatus(status int) {
	r.Status = status
}

func (r *responseSimpleJSON) SetStatus(status int) {
	r.Status = status
}

func writeResponse(w http.ResponseWriter, status int, res response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	res.SetStatus(status)
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
