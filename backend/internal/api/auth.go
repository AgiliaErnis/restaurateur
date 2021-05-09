package api

import (
	"encoding/json"
	"github.com/AgiliaErnis/restaurateur/backend/internal/db"
	"github.com/gorilla/handlers"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"html"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	// ORIGIN_ALLOWED is `scheme://dns[:port]`, or `*` (insecure)
	originsOk = handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	headersOk = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methodsOk = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	store     *sessions.CookieStore
)

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

func isAuthenticated(w http.ResponseWriter, r *http.Request, addTime bool) (bool, int) {
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
	if addTime {
		session.Values["expires"] = time.Now().Add(time.Minute * 15).Unix()
		session.Options.MaxAge = 60 * 15
		err = session.Save(r, w)
		if err != nil {
			log.Println(err)
			return false, 0
		}
	}
	return true, session.Values["user-id"].(int)
}

// registerHandler godoc
// @Summary Registers a user
// @Tags register
// @Accept  json
// @Produce  json
// @Param user body db.User true "Create a new user"
// @Success 200 {object} responseUserJSON
// @Success 400 {object} responseSimpleJSON
// @Failure 500 {string} []byte
// @Router /register [post]
func registerHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "registerHandler")
	user := &db.User{}
	res := &responseFullJSON{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(user)
	if err != nil {
		log.Println(err)
		writeResponse(w, http.StatusInternalServerError, res)
		return
	}
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		log.Println(err)
		resErr := &responseSimpleJSON{}
		resErr.Msg = "Credentials do not comply with requirements"
		writeResponse(w, http.StatusBadRequest, resErr)
		return
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		writeResponse(w, http.StatusInternalServerError, res)
		return
	}
	user.Password = string(pass)
	user.Name = html.EscapeString(user.Name)
	err = db.SaveUser(user)
	if err != nil {
		log.Println(err)
		if err.Error() == "pq: duplicate key value violates unique constraint \"restaurateur_users_email_key\"" {
			resErr := &responseSimpleJSON{}
			resErr.Msg = "Email already used"
			writeResponse(w, http.StatusBadRequest, resErr)
			return
		}
		writeResponse(w, http.StatusInternalServerError, res)
		return
	}
	res.Msg = "Registration successful!"
	writeResponse(w, http.StatusOK, res)
}

// loginHandler godoc
// @Summary Logs in a user
// @Description Logs in a user if the request headers contain an authenticated cookie.
// @Tags login
// @Accept  json
// @Produce  json
// @Param user body db.User true "Logs in a new user"
// @Success 200 {object} responseUserJSON
// @Success 400 {object} responseSimpleJSON
// @Failure 500 {string} []byte
// @Router /login [post]
func loginHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "loginHandler")
	session, _ := store.Get(r, "session-id")
	user := &db.User{}
	res := &responseUserJSON{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(user)
	if err != nil {
		log.Println(err)
		resErr := &responseSimpleJSON{}
		resErr.Msg = "Invalid JSON data"
		writeResponse(w, http.StatusBadRequest, resErr)
		return
	}
	dbUser, err := db.GetUserByEmail(user.Email)
	log.Println(err)
	errf := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if errf != nil {
		resErr := &responseSimpleJSON{}
		resErr.Msg = "Invalid password"
		writeResponse(w, http.StatusForbidden, resErr)
		return
	}
	session.Values["user-id"] = dbUser.ID
	session.Values["authenticated"] = true
	t := time.Now().Add(time.Minute * 15).Unix()
	session.Values["expires"] = t
	err = session.Save(r, w)
	if err != nil {
		log.Println(err)
		writeResponse(w, http.StatusInternalServerError, res)
		return
	}
	res.Msg = "Login successful!"
	res.User = &userResponse{Name: dbUser.Name, Email: dbUser.Email}
	writeResponse(w, http.StatusOK, res)
}

// logoutHandler godoc
// @Summary Logs out a user
// @Tags logout
// @Accept  json
// @Produce  json
// @Success 200 {object} responseSimpleJSON
// @Success 400 {object} responseSimpleJSON
// @Failure 500 {string} []byte
// @Router /logout [get]
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	auth, _ := isAuthenticated(w, r, false)
	res := &responseSimpleJSON{}
	if auth {
		session, _ := store.Get(r, "cookie-name")
		session.Values["authenticated"] = false
		session.Options.MaxAge = -1
		err := session.Save(r, w)
		if err != nil {
			log.Println(err)
			writeResponse(w, http.StatusInternalServerError, res)
			return
		}
		res.Msg = "Successfuly logged out!"
		writeResponse(w, http.StatusOK, res)
		return
	}
	res.Msg = "Not authenticated"
	writeResponse(w, http.StatusForbidden, res)
}
