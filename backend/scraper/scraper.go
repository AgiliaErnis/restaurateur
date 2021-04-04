package scraper

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

const restuBaseURL = "https://www.restu.cz"

type nominatimJSON []struct {
	PlaceID     int      `json:"place_id"`
	Licence     string   `json:"licence"`
	OsmType     string   `json:"osm_type"`
	OsmID       int      `json:"osm_id"`
	Boundingbox []string `json:"boundingbox"`
	Lat         string   `json:"lat"`
	Lon         string   `json:"lon"`
	DisplayName string   `json:"display_name"`
	Class       string   `json:"class"`
	Type        string   `json:"type"`
	Importance  float64  `json:"importance"`
	Icon        string   `json:"icon,omitempty"`
}

// Restaurant contains information needed about the restaurant
type Restaurant struct {
	Name            string
	Address         string
	District        string
	Images          []string
	Cuisines        []string
	PriceRange      string
	Rating          string
	URL             string
	PhoneNumber     string
	Lat             float64
	Lon             float64
	Vegan           bool
	Vegetarian      bool
	GlutenFree      bool
	WeeklyMenu      map[string]string
	OpeningHours    map[string]string
	Takeaway        bool
	DeliveryOptions []string
}

// RestaurantMenu stores name of the restaurant along with the weekly menu
type RestaurantMenu struct {
	RestaurantName string
	WeeklyMenu     map[string]string
}

// RequestError is returned when code other than 200 is received
// (other codes are not expected)
type RequestError struct {
	StatusCode int
	Err        error
}

type restaurantPair struct {
	restaurant Restaurant
	err        error
}
type menuPair struct {
	menu RestaurantMenu
	err  error
}

func (req *RequestError) Error() string {
	return fmt.Sprintf("Status %d: Error: %v", req.StatusCode, req.Err)
}

func getRestaurantMenu(link, restaurantName string, ch chan<- menuPair) {
	menu := RestaurantMenu{RestaurantName: restaurantName,
		WeeklyMenu: make(map[string]string)}
	url := restuBaseURL + link + "menu"
	res, err := http.Get(url)
	if err != nil {
		ch <- menuPair{menu, err}
		return
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		ch <- menuPair{menu, err}
		return
	}
	doc.Find(".menu-section").Each(func(i int, s *goquery.Selection) {
		foundDate := s.Find("h4").Text()
		s.Find(".c-menu-item").Each(func(i int, s *goquery.Selection) {
			food := s.Find(".menu-section__item-desc").Text()
			price := s.Find(".menu-section__item-price").Text()
			item := food + " " + price
			menu.WeeklyMenu[foundDate] = item
		})
	})
	ch <- menuPair{menu, nil}
	return
}

func getNominatimJSON(restaurant *Restaurant) (nominatimJSON, error) {
	var nominatim nominatimJSON
	url := "https://nominatim.openstreetmap.org/search?street=" +
		restaurant.Address + "&city=" + restaurant.District + "&format=json"
	res, err := http.Get(url)
	log.Println("Getting coordinates for", restaurant.Address)
	if err != nil {
		return nominatim, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nominatim, err
	}
	if err := json.Unmarshal(body, &nominatim); err != nil {
		return nominatim, err
	}
	return nominatim, nil
}

func (restaurant *Restaurant) setCoordinates() error {
	nominatim, err := getNominatimJSON(restaurant)
	if len(nominatim) == 0 || err != nil {
		restaurant.Lat = 0
		restaurant.Lon = 0
		return fmt.Errorf("Couldn't get coordinates for %q", restaurant.Name)
	}
	if lat, err := strconv.ParseFloat(nominatim[0].Lat, 64); err == nil {
		restaurant.Lat = lat
	}
	if lon, err := strconv.ParseFloat(nominatim[0].Lon, 64); err == nil {
		restaurant.Lon = lon
	}
	return nil
}

func visitLink(link, name, fullAddress string, ch chan<- restaurantPair) {
	address := strings.Split(fullAddress, ", ")[0]
	newRestaurant := Restaurant{Name: name, Address: address}
	url := restuBaseURL + link
	res, err := http.Get(url)
	if err != nil {
		ch <- restaurantPair{newRestaurant, err}
		return
	}
	var images []string
	defer res.Body.Close()
	if res.StatusCode != 200 {
		ch <- restaurantPair{newRestaurant, &RequestError{
			StatusCode: res.StatusCode,
			Err:        errors.New("Couldn't visit the link " + link),
		}}
		return
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		ch <- restaurantPair{newRestaurant, err}
		return
	}
	daysArray := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	ctr := 0
	openingHours := make(map[string]string)
	doc.Find(".opening-list__time").Each(func(i int, s *goquery.Selection) {
		if ctr >= 7 {
			return
		}
		openingHours[daysArray[ctr]] = s.Text()
		ctr++
	})
	doc.Find(".restaurant-detail-header-delivery-takeaway-container").Each(func(i int, s *goquery.Selection) {
		text := s.Find("h4").Text()
		if text == "" {
			newRestaurant.Takeaway = false
		} else {
			newRestaurant.Takeaway = true
			var deliveryOptions []string
			s.Find("a").Each(func(i int, s *goquery.Selection) {
				link, _ := s.Attr("href")
				deliveryOptions = append(deliveryOptions, link)
			})
			newRestaurant.DeliveryOptions = deliveryOptions
		}

	})
	newRestaurant.OpeningHours = openingHours
	scriptContent := doc.Find("script").Text()
	rTags, _ := regexp.Compile("restaurantTopics.*")
	restaurantTopics := rTags.FindString(scriptContent)
	rDistrict, _ := regexp.Compile("Praha [0-9]+")
	newRestaurant.District = rDistrict.FindString(fullAddress)
	tags := strings.Split(strings.Replace(restaurantTopics, "restaurantTopics': ", "", 1), ",")
	doc.Find(".tag").Each(func(i int, s *goquery.Selection) {
		tag := s.Text()
		if !(sliceContains(tags, tag)) {
			tags = append(tags, tag)
		}
	})
	newRestaurant.Cuisines = getCuisines(tags)
	newRestaurant.PriceRange = getPriceRange(tags)
	s := doc.Find(".restaurant-phone-popup__phone")
	phoneNum, _ := s.Attr("href")
	newRestaurant.PhoneNumber = strings.Replace(phoneNum, "tel:", "", 1)
	doc.Find(".track-restaurant-web").Each(func(i int, s *goquery.Selection) {
		restaurantURL, _ := s.Attr("href")
		newRestaurant.URL = restaurantURL
	})
	doc.Find("picture").Each(func(i int, s *goquery.Selection) {
		s.Find("img").Each(func(i int, s *goquery.Selection) {
			image, _ := s.Attr("src")
			images = append(images, image)
		})
	})
	ratingChart := doc.Find(".rating-chart")
	newRestaurant.Rating = ratingChart.Find("figcaption").Text()
	for i, image := range images {
		if strings.Contains(image, "placeholder.svg") {
			newRestaurant.Images = images[:i]
			break
		}
	}
	ch <- restaurantPair{newRestaurant, nil}
	return
}

func getLinks(doc *goquery.Document) []string {
	var links []string
	doc.Find(".card-item-link").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		links = append(links, link)
	})
	return links
}

func getNames(doc *goquery.Document) []string {
	var names []string
	doc.Find(".card-item__title").Each(func(i int, s *goquery.Selection) {
		name := s.Find("span").Text()
		names = append(names, name)
	})
	return names
}

func getAddresses(doc *goquery.Document) []string {
	var addresses []string
	doc.Find(".card-item__restaurant-address").Each(func(i int, s *goquery.Selection) {
		address := s.Find("span").Text()
		addresses = append(addresses, address)
	})
	return addresses
}

func sliceContains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func (restaurant *Restaurant) setVegan(veganRestaurants []string) {
	found := sliceContains(veganRestaurants, restaurant.Name)
	restaurant.Vegan = found
}

func (restaurant *Restaurant) setVegetarian(vegetarianRestaurants []string) {
	found := sliceContains(vegetarianRestaurants, restaurant.Name)
	restaurant.Vegetarian = found
}

func (restaurant *Restaurant) setGlutenFree(glutenFreeRestaurants []string) {
	found := sliceContains(glutenFreeRestaurants, restaurant.Name)
	restaurant.GlutenFree = found
}

func (restaurant *Restaurant) setWeeklyMenu(menus []*RestaurantMenu) {
	for _, menu := range menus {
		if menu.RestaurantName == restaurant.Name {
			restaurant.WeeklyMenu = menu.WeeklyMenu
			return
		}
	}
}

// GetRestaurants queries restu with the provided searchTerm
// and returns information about found restaurants
func GetRestaurants(searchTerm string) ([]*Restaurant, error) {
	var restaurants []*Restaurant
	restaurantMenus, err := GetRestaurantMenus()
	if err != nil {
		return restaurants, err
	}
	veganRestaurants, err := getFilteredRestaurants("veganske-restaurace")
	if err != nil {
		return restaurants, err
	}
	vegetarianRestaurants, err := getFilteredRestaurants("vegetarianske-restaurace")
	if err != nil {
		return restaurants, err
	}
	glutenFreeRestaurants, err := getFilteredRestaurants("bezlepkove-restaurace")
	if err != nil {
		return restaurants, err
	}
	restaurantChannel := make(chan restaurantPair)
	// set max go routines to not hammer the website too much (waitng for OSM afterwards anyway)
	maxGoroutines := 5
	guard := make(chan struct{}, maxGoroutines)
	var workerWaitGroup sync.WaitGroup
	var collectorWaitGroup sync.WaitGroup
	collectorWaitGroup.Add(1)
	go func() {
		defer collectorWaitGroup.Done()
		for {
			pair, ok := <-restaurantChannel
			if !ok {
				return
			}
			restaurant := pair.restaurant
			err := pair.err
			if err != nil {
				panic(err)
			}
			// calling setCoordinates only in 1 thread, because
			// OSM has limits on calls per second
			err = restaurant.setCoordinates()
			if err != nil {
				log.Println(err)
			}
			restaurant.setVegan(veganRestaurants)
			restaurant.setVegetarian(vegetarianRestaurants)
			restaurant.setGlutenFree(glutenFreeRestaurants)
			restaurant.setWeeklyMenu(restaurantMenus)
			restaurants = append(restaurants, &restaurant)
		}
	}()
	pageNum := 1
	for {
		url := restuBaseURL + "/vyhledavani/?term=" + searchTerm +
			"&page=" + strconv.Itoa(pageNum)
		res, err := http.Get(url)
		if err != nil {
			return restaurants, err
		}
		log.Printf("Processing %s page %d\n", searchTerm, pageNum)
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return restaurants, &RequestError{
				StatusCode: res.StatusCode,
				Err:        errors.New("Couldn't access URL " + url),
			}
		}
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return restaurants, nil
		}
		links := getLinks(doc)
		names := getNames(doc)
		addresses := getAddresses(doc)
		for i := range links {
			link := links[i]
			name := names[i]
			address := addresses[i]
			guard <- struct{}{}
			go func(link, name, address string) {
				workerWaitGroup.Add(1)
				defer workerWaitGroup.Done()
				visitLink(link, name, address, restaurantChannel)
				<-guard
			}(link, name, address)
		}
		if len(links) == 0 {
			workerWaitGroup.Wait()
			close(restaurantChannel)
			collectorWaitGroup.Wait()
			restaurants = verifyCoordinates(restaurants)
			return restaurants, nil
		}
		pageNum++
	}
}

// GetRestaurantMenus scrapes restu and returns
// all restaurants with a weekly menu
func GetRestaurantMenus() ([]*RestaurantMenu, error) {
	var restaurantMenus []*RestaurantMenu
	menuChannel := make(chan menuPair, 1)
	var workerWaitGroup sync.WaitGroup
	var collectorWaitGroup sync.WaitGroup
	collectorWaitGroup.Add(1)
	go func() {
		defer collectorWaitGroup.Done()
		for {
			pair, ok := <-menuChannel
			if !ok {
				return
			}
			menu := pair.menu
			err := pair.err
			if err != nil {
				panic(err)
			}
			restaurantMenus = append(restaurantMenus, &menu)
		}
	}()
	pageNum := 1
	for {
		log.Printf("Processing menus page %d\n", pageNum)
		url := restuBaseURL + "/praha/maji-denni-menu" +
			"/?page=" + strconv.Itoa(pageNum)
		res, err := http.Get(url)
		if err != nil {
			return restaurantMenus, err
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return restaurantMenus, &RequestError{
				StatusCode: res.StatusCode,
				Err:        errors.New("Couldn't access menu URL " + url),
			}
		}
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return restaurantMenus, err
		}
		links := getLinks(doc)
		names := getNames(doc)
		for i := range links {
			link := links[i]
			name := names[i]
			go func(link, name string) {
				defer workerWaitGroup.Done()
				workerWaitGroup.Add(1)
				getRestaurantMenu(link, name, menuChannel)
			}(link, name)
		}
		if len(links) == 0 {
			workerWaitGroup.Wait()
			close(menuChannel)
			collectorWaitGroup.Wait()
			return restaurantMenus, nil
		}
		pageNum++
	}
}

// getFilteredRestaurants takes a url suffix with a filter and
// returns names of restaurants matching that filter
func getFilteredRestaurants(urlSuffix string) ([]string, error) {
	var restaurantNames []string
	pageNum := 1
	for {
		url := restuBaseURL + "/" + urlSuffix + "/praha/?page=" + strconv.Itoa(pageNum)
		log.Printf("Processing %s page %d\n", urlSuffix, pageNum)
		res, err := http.Get(url)
		if err != nil {
			return restaurantNames, err
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return restaurantNames, &RequestError{
				StatusCode: res.StatusCode,
				Err:        errors.New("Couldn't access menu URL " + url),
			}
		}
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return restaurantNames, err
		}
		names := getNames(doc)
		for _, name := range names {
			restaurantNames = append(restaurantNames, name)
		}
		if len(names) == 0 {
			return restaurantNames, nil
		}
		pageNum++
	}
}

func verifyCoordinates(restaurants []*Restaurant) []*Restaurant {
	var newRestaurants []*Restaurant
	for _, restaurant := range restaurants {
		if restaurant.Lat == 0 || restaurant.Lon == 0 {
			log.Printf("Retrying to set coordinates for %q\n", restaurant.Name)
			err := restaurant.setCoordinates()
			time.Sleep(time.Second) // nominatim has a limit of 1 call per second
			if restaurant.Lat != 0 && restaurant.Lon != 0 {
				newRestaurants = append(newRestaurants, restaurant)
			} else {
				log.Println(err)
			}
		} else {
			newRestaurants = append(newRestaurants, restaurant)
		}
	}
	return newRestaurants
}
