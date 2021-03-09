package main

import (
    "fmt"
    "github.com/AgiliaErnis/restaurateur/backend/scraper"
)


func main() {
	var restaurants []scraper.Restaurant
	scraper.WalkResults("shop", 1, &restaurants)
	for i := range restaurants {
		fmt.Println(restaurants[i].Name)
	}
	var restaurantMenus []scraper.RestaurantMenu
	scraper.WalkMenuResults(1, &restaurantMenus)
	for i := range restaurantMenus {
		fmt.Println(restaurantMenus[i].RestaurantName)
		fmt.Println(restaurantMenus[i].WeeklyMenu)
	}
}
