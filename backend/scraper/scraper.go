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
	Name        string
	Address     string
	Images      []string
	Tags        []string
	Rating      string
	URL         string
	PhoneNumber string
	Lat         float64
	Lon         float64
	Vegan       bool
	Vegetarian  bool
	WeeklyMenu  map[string]string
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

func (restaurant *Restaurant) setCoordinates() error {
	address := strings.Split(restaurant.Address, ",")[0]
	url := "https://nominatim.openstreetmap.org/search?street=" +
		address + "&format=json"
	res, err := http.Get(url)
	log.Println("Getting coordinates for", restaurant.Name)
	log.Println("Getting coordinates for", restaurant.Address)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var nominatim nominatimJSON
	json.Unmarshal(body, &nominatim)
	if len(nominatim) == 0 {
		return errors.New("Couldn't get coordinates")
	}
	if lat, err := strconv.ParseFloat(nominatim[0].Lat, 64); err == nil {
		restaurant.Lat = lat
	}
	if lon, err := strconv.ParseFloat(nominatim[0].Lon, 64); err == nil {
		restaurant.Lon = lon
	}
	return nil
}

func visitLink(link, name, address string, ch chan<- restaurantPair) {
	newRestaurant := Restaurant{Name: name, Address: address}
	url := restuBaseURL + link
	res, err := http.Get(url)
	if err != nil {
		ch <- restaurantPair{newRestaurant, err}
		return
	}
	var images []string
	var tags []string
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
	scriptContent := doc.Find("script").Text()
	r, _ := regexp.Compile("\\+420[0-9]{9}")
	newRestaurant.PhoneNumber = r.FindString(scriptContent)
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
	doc.Find(".tag").Each(func(i int, s *goquery.Selection) {
		tag := s.Text()
		tags = append(tags, tag)
	})
	ratingChart := doc.Find(".rating-chart")
	newRestaurant.Rating = ratingChart.Find("figcaption").Text()
	for i, image := range images {
		if strings.Contains(image, "placeholder.svg") {
			newRestaurant.Images = images[:i]
			break
		}
	}
	for i, tag := range tags {
		if strings.Contains(tag, "Další") {
			newRestaurant.Tags = tags[:i]
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

func contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func (restaurant *Restaurant) setVegan(veganRestaurants []string) {
	found := contains(veganRestaurants, restaurant.Name)
	restaurant.Vegan = found
}

func (restaurant *Restaurant) setVegetarian(vegetarianRestaurants []string) {
	found := contains(vegetarianRestaurants, restaurant.Name)
	restaurant.Vegetarian = found
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
	restaurantMenus, err := getRestaurantMenus()
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
				continue
			}
			restaurant.setVegan(veganRestaurants)
			restaurant.setVegetarian(vegetarianRestaurants)
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
		log.Printf("Processing search page %d\n", pageNum)
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
			return restaurants, nil
		}
		pageNum++
	}
}

// getRestaurantMenus scrapes restu and returns
// all restaurants with a weekly menu
func getRestaurantMenus() ([]*RestaurantMenu, error) {
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
