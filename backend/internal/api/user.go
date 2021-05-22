package api

import (
	"encoding/json"
	"github.com/AgiliaErnis/restaurateur/backend/internal/db"
	"golang.org/x/crypto/bcrypt"
	"html"
	"log"
	"net/http"
)

type userPassword struct {
	Password string `json:"password"`
}

type userUpdate struct {
	OldPassword string `json:"oldPassword" validate:"required,min=6,max=64"`
	NewPassword string `json:"newPassword" validate:"required,min=6,max=64"`
	NewUsername string `json:"newUsername" validate:"required,min=2,max=32"`
}

// userGetHandler godoc
// @Summary Get info about a user
// @Description Returns a JSON with user info if the request headers contain an authenticated cookie.
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} responseUserJSON
// @Success 400 {object} responseSimpleJSON
// @Failure 500 {string} []byte
// @Router /user [get]
func userGetHandler(w http.ResponseWriter, r *http.Request) {
	id := getUserIDFromCookie(r)
	res := &responseUserJSON{}
	user, _ := db.GetUserByID(id)
	res.User = &userResponseFull{Name: user.Name, Email: user.Email}
	savedRestaurants, err := db.GetSavedRestaurantsArr(id)
	if err != nil {
		log.Println("Couldn't get saved restaurants for user id:", id)
		log.Println(err)
	} else {
		res.User.SavedRestaurants = savedRestaurants
	}
	res.Msg = "Success"
	writeResponse(w, http.StatusOK, res)
}

// userDeleteHandler godoc
// @Summary Deletes a user
// @Description Deletes a user if the request headers contain an authenticated cookie and the body contains a JSON with a valid password.
// @Tags user
// @Accept  json
// @Produce  json
// @Param password body userPassword true "Password of the current user"
// @Success 200 {object} responseSimpleJSON
// @Success 400 {object} responseSimpleJSON
// @Failure 500 {string} []byte
// @Router /user [delete]
// userDeleteHandler godoc
func userDeleteHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "userDeleteHandler")
	id := getUserIDFromCookie(r)
	res := &responseSimpleJSON{}
	user, err := db.GetUserByID(id)
	if err != nil {
		log.Println(err)
		res.Msg = "User doesn't exist in the db"
		writeResponse(w, http.StatusBadRequest, res)
		return
	}
	password := &userPassword{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(password)
	if err != nil {
		log.Println(err)
		res.Msg = "Missing a password"
		writeResponse(w, http.StatusBadRequest, res)
		return
	}
	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password.Password))
	if errf != nil {
		log.Println(errf)
		res.Msg = "Invalid password"
		writeResponse(w, http.StatusForbidden, res)
		return
	}
	err = db.DeleteUser(id)
	if err != nil {
		res.Msg = "Couldn't delete the record"
		writeResponse(w, http.StatusBadRequest, res)
		return
	}
	res.Msg = "Successfuly deleted the user!"
	writeResponse(w, http.StatusOK, res)
}

// userPatchHandler godoc
// @Summary Updates a user's password or username
// @Description Updates user's password or username based on the provided JSON. Only 1 field can be updated at a time. For password you need to provide "oldPassword" and "newPassword" fields, omitting the "newUsername" field and vice versa if you'd like to update the username
// @Tags user
// @Param updateJSON body userUpdate true "Create a new user"
// @Accept  json
// @Produce  json
// @Success 200 {object} responseSimpleJSON
// @Success 400 {object} responseSimpleJSON
// @Failure 500 {string} []byte
// @Router /user [patch]
func userPatchHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "userPatchHandler")
	id := getUserIDFromCookie(r)
	userUpdate := &userUpdate{}
	res := &responseUserJSON{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(userUpdate)
	if err != nil {
		resErr := &responseSimpleJSON{}
		resErr.Msg = "Wrong or missing fields in JSON"
		writeResponse(w, http.StatusBadRequest, resErr)
		return
	}
	if userUpdate.NewUsername != "" {
		username := html.EscapeString(userUpdate.NewUsername)
		err = db.UpdateOne("name", username, id)
		if err != nil {
			resErr := &responseSimpleJSON{}
			resErr.Msg = "Couldn't update username"
			writeResponse(w, http.StatusBadRequest, resErr)
		}
		res.Msg = "Successfuly updated the username!"
		writeResponse(w, http.StatusOK, res)
		return
	} else if userUpdate.OldPassword != "" && userUpdate.NewPassword != "" {
		user, _ := db.GetUserByID(id)
		errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userUpdate.OldPassword))
		if errf != nil {
			log.Println(errf)
			resErr := &responseSimpleJSON{}
			resErr.Msg = "Invalid password"
			writeResponse(w, http.StatusForbidden, resErr)
			return
		}
		pass, err := bcrypt.GenerateFromPassword([]byte(userUpdate.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
			writeResponse(w, http.StatusInternalServerError, res)
			return
		}
		password := string(pass)
		err = db.UpdateOne("password", password, id)
		if err != nil {
			res.Msg = "Couldn't update password"
			writeResponse(w, http.StatusBadRequest, res)
		}
		res.Msg = "Successfuly updated the user's password!"
		writeResponse(w, http.StatusOK, res)
		return
	}
	resErr := &responseSimpleJSON{}
	resErr.Msg = "Wrong or missing fields in JSON"
	writeResponse(w, http.StatusBadRequest, resErr)
}

// savedPostHandler godoc
// @Summary Saves a restaurant mapped to a user to db
// @Tags Saved restaurants
// @Accept  json
// @Produce  json
// @Param restaurantID body restaurantIDJSON true "ID of restaurant to save"
// @Success 200 {object} responseSimpleJSON
// @Success 400 {object} responseSimpleJSON
// @Failure 500 {string} []byte
// @Router /user/saved-restaurants [post]
// savedPostHandler godoc
func savedPostHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "savedPostHandler")
	res := &responseSimpleJSON{}
	id := getUserIDFromCookie(r)
	restaurantIDJSON := &restaurantIDJSON{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(restaurantIDJSON)
	if err != nil {
		log.Println(err)
		res.Msg = "Missing a field"
		writeResponse(w, http.StatusBadRequest, res)
		return
	}
	err = db.AddSavedRestaurant(restaurantIDJSON.RestaurantID, id)
	if err != nil {
		log.Println(err)
		res.Msg = "Couldn't add the restaurant"
		writeResponse(w, http.StatusBadRequest, res)
		return
	}
	res.Msg = "Successfuly added the restaurant!"
	writeResponse(w, http.StatusOK, res)
}

// savedDeleteHandler godoc
// @Summary Deletes a saved restaurant
// @Description Deletes a saved restaurant if the request headers contain an authenticated cookie
// @Tags Saved restaurants
// @Accept  json
// @Produce  json
// @Param restaurantID body restaurantIDJSON true "ID of restaurant to delete"
// @Success 200 {object} responseSimpleJSON
// @Success 400 {object} responseSimpleJSON
// @Failure 500 {string} []byte
// @Router /user/saved-restaurants [delete]
// savedDeleteHandler godoc
func savedDeleteHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "savedDeleteHandler")
	res := &responseSimpleJSON{}
	id := getUserIDFromCookie(r)
	restaurantIDJSON := &restaurantIDJSON{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(restaurantIDJSON)
	if err != nil {
		log.Println(err)
		res.Msg = "Missing a field"
		writeResponse(w, http.StatusBadRequest, res)
		return
	}
	err = db.DeleteSavedRestaurant(restaurantIDJSON.RestaurantID, id)
	if err != nil {
		log.Println(err)
		res.Msg = "Couldn't delete the restaurant"
		writeResponse(w, http.StatusBadRequest, res)
		return
	}
	res.Msg = "Successfuly deleted the restaurant!"
	writeResponse(w, http.StatusOK, res)
}
