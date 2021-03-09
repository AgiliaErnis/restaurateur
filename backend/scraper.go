package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Restaurant struct {
	Name    string
	Address string
	Images  []string
	Tags    []string
	Rating  string
}

type RestaurantMenu struct {
	RestaurantName string
	WeeklyMenu     map[string]string
}

func (restaurantMenu *RestaurantMenu) updateMenu(link string) {
	url := "https://www.restu.cz" + link + "menu"
	res, _ := http.Get(url)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".menu-section").Each(func(i int, s *goquery.Selection) {
		foundDate := s.Find("h4").Text()
		s.Find(".c-menu-item").Each(func(i int, s *goquery.Selection) {
			food := s.Find(".menu-section__item-desc").Text()
			price := s.Find(".menu-section__item-price").Text()
			item := food + " " + price
			restaurantMenu.WeeklyMenu[foundDate] = item
		})
	})
}

func visitLink(link string) ([]string, []string, string) {
	url := "https://www.restu.cz" + link
	res, _ := http.Get(url)
	var images []string
	var tags []string
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
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
	rating := ratingChart.Find("figcaption").Text()
	for i := range images {
		if strings.Contains(images[i], "placeholder.svg") {
			images = images[:i]
			break
		}
	}
	for i := range tags {
		if strings.Contains(tags[i], "Další") {
			tags = tags[:i]
			break
		}
	}
	return images, tags, rating
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

func walkResults(searchTerm string, pageNum int, restaurants *[]Restaurant) {
	url := "https://www.restu.cz/vyhledavani/?term=" + searchTerm +
		"&page=" + strconv.Itoa(pageNum)
	res, _ := http.Get(url)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	links := getLinks(doc)
	names := getNames(doc)
	addresses := getAddresses(doc)
	for i := range links {
		name := names[i]
		address := addresses[i]
		images, tags, rating := visitLink(links[i])
		restaurant := Restaurant{
			Name:    name,
			Address: address,
			Images:  images,
			Tags:    tags,
			Rating:  rating}
		*restaurants = append(*restaurants, restaurant)
	}
	// continue to the next page if the current one had results
	if len(links) != 0 {
		walkResults(searchTerm, pageNum+1, restaurants)
	}
}

func walkMenuResults(pageNum int, restaurantMenus *[]RestaurantMenu) {
	url := "https://www.restu.cz/praha/maji-denni-menu" +
		"/?page=" + strconv.Itoa(pageNum)
	res, _ := http.Get(url)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	links := getLinks(doc)
	names := getNames(doc)
	for i := range links {
		restaurantMenu := RestaurantMenu{RestaurantName: names[i],
			WeeklyMenu: make(map[string]string)}
		restaurantMenu.updateMenu(links[i])
		*restaurantMenus = append(*restaurantMenus, restaurantMenu)
	}
	// continue to the next page if the current one had results
	if len(links) != 0 {
		walkMenuResults(pageNum+1, restaurantMenus)
	}
}

func main() {
	var restaurants []Restaurant
	walkResults("shop", 1, &restaurants)
	for i := range restaurants {
		fmt.Println(restaurants[i].Name)
	}
	var restaurantMenus []RestaurantMenu
	walkMenuResults(1, &restaurantMenus)
	for i := range restaurantMenus {
		fmt.Println(restaurantMenus[i].RestaurantName)
		fmt.Println(restaurantMenus[i].WeeklyMenu)
	}
}
