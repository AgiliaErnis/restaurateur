package main

import (
	"fmt"
	"github.com/AgiliaErnis/restaurateur/backend/scraper"
	"log"
)

func main() {
	restaurants, err := scraper.GetRestaurants("shop")
	if err != nil {
		log.Fatal(err)
	}
	for _, restaurant := range restaurants {
		fmt.Println(restaurant.Name)
		fmt.Println(restaurant.Address)
		fmt.Println(restaurant.Lat)
		fmt.Println(restaurant.Lon)
		fmt.Println(restaurant.PhoneNumber)
	}
	restaurantMenus, err := scraper.GetRestaurantMenus()
	if err != nil {
		log.Fatal(err)
	}
	for _, menu := range restaurantMenus {
		fmt.Println(menu.RestaurantName)
	}
}
