package main

import (
	"github.com/AgiliaErnis/restaurateur/backend/scraper"
	"math"
)

func toRadians(num float64) float64 {
	return num * (math.Pi / 180)
}

// calculates distance (in meters) between two point based on their coordinates
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	lat1Rad := toRadians(lat1)
	lon1Rad := toRadians(lon1)
	lat2Rad := toRadians(lat2)
	lon2Rad := toRadians(lon2)

	deltaLon := lon2Rad - lon1Rad
	deltaLat := lat2Rad - lat1Rad
	a := math.Pow(math.Sin(deltaLat/2), 2) + math.Cos(lat1Rad)*
		math.Cos(lat2Rad)*math.Pow(math.Sin(deltaLon/2), 2)
	c := 2 * math.Asin(math.Sqrt(a))
	var radius float64 = 6371000 // Radius of earth in meters
	return c * radius
}

func getRestaurantsInRadius(restaurants []*scraper.Restaurant, lat, lon, radius float64) []*scraper.Restaurant {
	var filteredRestaurants []*scraper.Restaurant
	for _, r := range restaurants {
		distance := haversine(lat, lon, r.Lat, r.Lon)
		if distance <= radius {
			filteredRestaurants = append(filteredRestaurants, r)
		}
	}
	return filteredRestaurants
}
