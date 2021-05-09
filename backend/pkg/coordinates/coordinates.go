package coordinates

import (
	"fmt"
	"github.com/AgiliaErnis/restaurateur/backend/pkg/scraper"
	"math"
	"net/url"
	"strconv"
	"strings"
)

func toRadians(num float64) float64 {
	return num * (math.Pi / 180)
}

// calculates distance (in meters) between two point based on their coordinates
func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
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

func GetCoordinates(params url.Values) (float64, float64, error) {
	addressParam := params.Get("address")
	radiusParam := params.Get("radius")
	if radiusParam == "ignore" {
		return 0, 0, nil
	} else if addressParam != "" {
		var lat float64
		var lon float64
		i := strings.Index(addressParam, "Praha")
		var address string
		var city string
		if i > -1 {
			address = addressParam[:i]
			city = addressParam[i:]
		} else {
			address = addressParam
			city = "Praha"
		}
		nominatim, err := scraper.GetNominatimJSON(address, city)
		if len(nominatim) == 0 || err != nil {
			return 0, 0, fmt.Errorf("Couldn't get coordinates for %q", addressParam)
		}
		lat, errLat := strconv.ParseFloat(nominatim[0].Lat, 64)
		lon, errLon := strconv.ParseFloat(nominatim[0].Lon, 64)
		if errLat != nil || errLon != nil {
			return 0, 0, fmt.Errorf("Coordinates for %s not found", addressParam)
		}
		return lat, lon, nil
	}
	latParam := params.Get("lat")
	lonParam := params.Get("lon")
	lat, errLat := strconv.ParseFloat(latParam, 64)
	lon, errLon := strconv.ParseFloat(lonParam, 64)
	if errLat != nil || errLon != nil {
		return 0, 0, fmt.Errorf("Invalid coordinates(Lat: %s, Lon: %s)", latParam, lonParam)
	}
	return lat, lon, nil
}
