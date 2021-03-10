package scraper

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

const RestuBaseUrl = "https://www.restu.cz"

type restaurant struct {
	Name    string
	Address string
	Images  []string
	Tags    []string
	Rating  string
}

type restaurantMenu struct {
	RestaurantName       string
	WeeklyMenu map[string]string
}

type requestError struct {
	StatusCode int
	Err        error
}

func (req *requestError) Error() string {
	return fmt.Sprintf("Status %d: Error: %v", req.StatusCode, req.Err)
}

func getRestaurantMenu(link, restaurantName string) (restaurantMenu, error) {
	menu := restaurantMenu{RestaurantName: restaurantName,
		WeeklyMenu: make(map[string]string)}
	url := RestuBaseUrl + link + "menu"
	res, err := http.Get(url)
	if err != nil {
		return menu, err
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return menu, err
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
	return menu, nil
}

func visitLink(link, name, address string) (restaurant, error) {
	newRestaurant := restaurant{Name: name, Address: address}
	url := RestuBaseUrl + link
	res, err := http.Get(url)
	if err != nil {
		return newRestaurant, err
	}
	var images []string
	var tags []string
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return newRestaurant, &requestError{
			StatusCode: res.StatusCode,
			Err:        errors.New("Couldn't visit the link " + link),
		}
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return newRestaurant, err
	}
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
	return newRestaurant, nil
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

func GetRestaurants(searchTerm string) ([]restaurant, error) {
	var restaurants []restaurant
	pageNum := 1
	for {
		url := RestuBaseUrl + "/vyhledavani/?term=" + searchTerm +
			"&page=" + strconv.Itoa(pageNum)
		res, err := http.Get(url)
		if err != nil {
			return restaurants, err
		}
		fmt.Printf("%c[2K", 27)
		fmt.Printf(" Processing search page %d\r", pageNum)
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return restaurants, &requestError{
				StatusCode: res.StatusCode,
				Err:        errors.New("Couldn't acces URL " + url),
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
			name := names[i]
			address := addresses[i]
			newRestaurant, err := visitLink(links[i], name, address)
			if err != nil {
				return restaurants, err
			}
			restaurants = append(restaurants, newRestaurant)
		}
		if len(links) == 0 {
			return restaurants, nil
		} else {
			pageNum += 1
		}
	}
}

func GetRestaurantMenus() ([]restaurantMenu, error) {
	var restaurantMenus []restaurantMenu
	pageNum := 1
	for {
		fmt.Printf("%c[2K", 27)
		fmt.Printf(" Processing menus page %d\r", pageNum)
		url := RestuBaseUrl + "/praha/maji-denni-menu" +
			"/?page=" + strconv.Itoa(pageNum)
		res, err := http.Get(url)
		if err != nil {
			return restaurantMenus, err
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return restaurantMenus, &requestError{
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
			restaurantMenu, err := getRestaurantMenu(links[i], names[i])
			if err != nil {
				return restaurantMenus, nil
			}
			restaurantMenus = append(restaurantMenus, restaurantMenu)
		}
		if len(links) == 0 {
			return restaurantMenus, nil
		} else {
			pageNum += 1
		}
	}
}
