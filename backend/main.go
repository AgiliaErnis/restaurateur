package main

import (
	"fmt"
	"github.com/AgiliaErnis/restaurateur/backend/scraper"
	"log"
)

func main() {
	restaurants, err := scraper.GetRestaurants("praha")
	if err != nil {
		log.Fatal(err)
	}
}
