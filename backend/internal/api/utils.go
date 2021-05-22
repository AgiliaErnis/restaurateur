package api

import (
	"fmt"
	"github.com/AgiliaErnis/restaurateur/backend/internal/db"
	"github.com/AgiliaErnis/restaurateur/backend/pkg/scraper"
	"log"
	"net/http"
	"time"
)

func logRequest(r *http.Request, handlerName string) {
	method := r.Method
	endpoint := r.URL
	clientAddr := r.RemoteAddr
	log.Printf("Incoming request\n"+
		"Method: %q\n"+
		"Client's address: %q\n"+
		"Request URL: %q\n"+
		"Handler: %q\n"+
		"Headers: %q\n",
		method, clientAddr, endpoint, handlerName, r.Header)
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

func menuUpdater() {
	t := time.Now()
	for {
		nextRun := time.Date(t.Year(), t.Month(), t.Day(), 11, 0, 0, 0, t.Location())
		time.Sleep(time.Until(nextRun))
		log.Println("Updating weekly menus...")
		menus, err := scraper.GetRestaurantMenus()
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Updating weekly menus in the db")
			db.UpdateWeeklyMenus(menus)
			log.Println("Weekly menus successfuly updated!")
		}
		t = t.Add(24 * time.Hour)
	}
}
