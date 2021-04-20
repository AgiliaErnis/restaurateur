package scraper

var cuisineTranslatedMap = map[string]string{
	"Americká":     "American",
	"Italská":      "Italian",
	"Asijská":      "Asian",
	"Indická":      "Indian",
	"Japonská":     "Japanese",
	"Vietnamská":   "Vietnamese",
	"Španělská":    "Spanish",
	"Středomořská": "Mediterranean",
	"Francouzská":  "French",
	"Thajská":      "Thai",
	"Mexická":      "Mexican",
	"Mezinárodní":  "International",
	"Česká":        "Czech",
	"Anglická":     "English",
	"Balkánská":    "Balkan",
	"Brazilská":    "Brazil",
	"Ruská":        "Russian",
	"Čínská":       "Chinese",
	"Řecká":        "Greek",
	"Arabská":      "Arabic",
	"Korejská":     "Korean",
}

var trackedCuisines = func() []string {
	keys := make([]string, len(cuisineTranslatedMap))
	for k := range cuisineTranslatedMap {
		keys = append(keys, k)
	}
	return keys
}()

func getCuisines(tags []string) []string {
	var foundCuisines []string
	for _, item := range tags {
		found := SliceContains(trackedCuisines, item)
		if found && item != "" {
			foundCuisines = append(foundCuisines, cuisineTranslatedMap[item])
		}
	}
	return foundCuisines
}

func getPriceRange(tags []string) string {
	for _, item := range tags {
		switch item {
		case "Do 300 Kč":
			return "0 - 300 Kč"
		case "300 - 600 Kč":
			return item
		case "Nad 600 Kč":
			return "600+ Kč"
		}
	}
	return "Not available"
}
