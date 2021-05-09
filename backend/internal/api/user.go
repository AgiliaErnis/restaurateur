package api

import (
	"encoding/json"
	"github.com/AgiliaErnis/restaurateur/backend/internal/db"
	"golang.org/x/crypto/bcrypt"
	"html"
	"log"
	"net/http"
)

type userUpdate struct {
	OldPassword string `json:"oldPassword" validate:"required,min=6,max=64"`
	NewPassword string `json:"newPassword" validate:"required,min=6,max=64"`
	NewUsername string `json:"newUsername" validate:"required,min=2,max=32"`
}

func userGetHandler(w http.ResponseWriter, r *http.Request) {
	auth, id := isAuthenticated(w, r)
	if auth {
		res := &responseUserJSON{}
		user, _ := db.GetUserByID(id)
		res.User = &userResponse{Name: user.Name, Email: user.Email}
		res.Msg = "Success"
		writeResponse(w, http.StatusOK, res)
		return
	}
	resErr := &responseSimpleJSON{}
	resErr.Msg = "Not authenticated"
	writeResponse(w, http.StatusForbidden, resErr)
}

func userDeleteHandler(w http.ResponseWriter, r *http.Request) {
	auth, id := isAuthenticated(w, r)
	if auth {
		res := &responseSimpleJSON{}
		err := db.DeleteUser(id)
		if err != nil {
			res.Msg = "Couldn't delete the record"
			writeResponse(w, http.StatusBadRequest, res)
			return
		}
		res.Msg = "Successfuly deleted the user!"
		writeResponse(w, http.StatusOK, res)
		return
	}
	resErr := &responseSimpleJSON{}
	resErr.Msg = "Not authenticated"
	writeResponse(w, http.StatusForbidden, resErr)
}

func userPatchHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "userPatchHandler")
	auth, id := isAuthenticated(w, r)
	if auth {
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
		return
	}
	resErr := &responseSimpleJSON{}
	resErr.Msg = "Not authenticated"
	writeResponse(w, http.StatusForbidden, resErr)
}