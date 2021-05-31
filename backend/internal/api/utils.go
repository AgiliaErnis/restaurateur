package api

import (
	"fmt"
	"github.com/AgiliaErnis/restaurateur/backend/internal/db"
	"github.com/AgiliaErnis/restaurateur/backend/pkg/scraper"
	"log"
	"net/http"
)

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

func catchAllHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r, "catchAllHandler")
	path := r.URL.Path
	res := &responseSimpleJSON{}
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

func updateWeeklyMenus() {
	log.Println("Updating weekly menus...")
	menus, err := scraper.GetRestaurantMenus()
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Updating weekly menus in the db")
		db.UpdateWeeklyMenus(menus)
		log.Println("Weekly menus successfuly updated!")
	}
}
