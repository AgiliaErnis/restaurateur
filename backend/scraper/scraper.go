package scraper

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

const restuBaseURL = "https://www.restu.cz"

// Restaurant contains information needed about the restaurant
type Restaurant struct {
	Name        string
	Address     string
	Images      []string
	Tags        []string
	Rating      string
	URL         string
	PhoneNumber string
}

// RestaurantMenu stores name of the restaurant along with the weekly menu
type RestaurantMenu struct {
	RestaurantName string
	WeeklyMenu     map[string]string
}

// RequestError is returned when code other than 200 is returned
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

// GetRestaurants queries restu with the provided searchTerm
// and returns information about found restaurants
func GetRestaurants(searchTerm string) ([]Restaurant, error) {
	var restaurants []Restaurant
	restaurantChannel := make(chan restaurantPair)
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
			restaurants = append(restaurants, restaurant)
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
		fmt.Printf("%c[2K", 27)
		fmt.Printf(" Processing search page %d\r", pageNum)
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
			go func(link, name, address string) {
				workerWaitGroup.Add(1)
				defer workerWaitGroup.Done()
				visitLink(link, name, address, restaurantChannel)
			}(link, name, address)
		}
		if len(links) == 0 {
			break
		}
		pageNum++
	}
	workerWaitGroup.Wait()
	close(restaurantChannel)
	collectorWaitGroup.Wait()
	fmt.Println()
	return restaurants, nil
}

// GetRestaurantMenus scrapes restu and returns
// all restaurants with a weekly menu
func GetRestaurantMenus() ([]RestaurantMenu, error) {
	var restaurantMenus []RestaurantMenu
	menuChannel := make(chan menuPair)
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
			restaurantMenus = append(restaurantMenus, menu)
		}
	}()
	pageNum := 1
	for {
		fmt.Printf("%c[2K", 27)
		fmt.Printf(" Processing menus page %d\r", pageNum)
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
			fmt.Println()
			return restaurantMenus, nil
		}
		pageNum++
	}
}
