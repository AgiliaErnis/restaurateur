package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/AgiliaErnis/restaurateur/backend/pkg/coordinates"
	"github.com/AgiliaErnis/restaurateur/backend/pkg/scraper"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	restaurantsSchema = `CREATE TABLE restaurants (
		 id SERIAL NOT NULL PRIMARY KEY,
		 name TEXT NOT NULL,
		 address TEXT NOT NULL,
		 district TEXT NOT NULL,
	 	 images TEXT,
	 	 cuisines TEXT,
	 	 price_range TEXT,
	 	 rating TEXT,
	 	 url TEXT,
	 	 phone_number TEXT,
	 	 lat NUMERIC NOT NULL,
	 	 lon NUMERIC NOT NULL,
	 	 vegan BOOLEAN,
	 	 vegetarian BOOLEAN,
		 gluten_free BOOLEAN,
		 weekly_menu JSON,
		 opening_hours JSON,
		 takeaway BOOLEAN,
		 delivery_options TEXT
	 );`
	restaurateurUsersSchema = `CREATE TABLE restaurateur_users (
        id SERIAL PRIMARY KEY,
        name TEXT,
        email TEXT UNIQUE,
        password TEXT
    );`
)

// User holds information about a user provided from a JSON
type User struct {
	Name     string `json:"username" validate:"required,min=2,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=64"`
}

// UserDB holds information about a user provided from the postgres DB
type UserDB struct {
	ID       int
	Name     string
	Email    string
	Password string
}

// RestaurantDB holds all data about a restaurant that is stored in the db
type RestaurantDB struct {
	ID              int            `db:"id" json:"ID" example:"1"`
	Name            string         `db:"name" json:"Name" example:"Steakhouse"`
	Address         string         `db:"address" json:"Address" example:"Polská 12"`
	District        string         `db:"district" json:"District" example:"Praha 1"`
	Images          pq.StringArray `db:"images" json:"Images" example:"image1.com, image2.com"`
	Cuisines        pq.StringArray `db:"cuisines" json:"Cuisines" example:"Italian,Czech"`
	PriceRange      string         `db:"price_range" json:"PriceRange" example:"300-600 Kč"`
	Rating          string         `db:"rating" json:"Rating" example:"4.6"`
	URL             string         `db:"url" json:"URL" example:"http://restaurant.com"`
	PhoneNumber     string         `db:"phone_number" json:"PhoneNumber" example:"+420123456789"`
	Lat             float64        `db:"lat" json:"Lat" example:"50.03493"`
	Lon             float64        `db:"lon" json:"Lon" example:"14.30320"`
	Vegan           bool           `db:"vegan" json:"Vegan"`
	Vegetarian      bool           `db:"vegetarian" json:"Vegetarian"`
	GlutenFree      bool           `db:"gluten_free" json:"GlutenFree"`
	WeeklyMenu      string         `db:"weekly_menu" json:"WeeklyMenu"`
	OpeningHours    string         `db:"opening_hours" json:"OpeningHours"`
	Takeaway        bool           `db:"takeaway" json:"Takeaway"`
	DeliveryOptions pq.StringArray `db:"delivery_options" json:"DeliveryOptions"`
}

// SortBy is a type for sorting the RestaurantDB struct
type SortBy func(r1, r2 *RestaurantDB) bool

// Sort sorts the RestaurantDB struct based on a provided sorting method
func (by SortBy) Sort(restaurants []*RestaurantDB) {
	rs := &restaurantDBSorter{
		restaurants: restaurants,
		by:          by,
	}
	sort.Sort(rs)
}

type restaurantDBSorter struct {
	restaurants []*RestaurantDB
	by          func(r1, r2 *RestaurantDB) bool
}

func (s *restaurantDBSorter) Len() int {
	return len(s.restaurants)
}

func (s *restaurantDBSorter) Swap(i, j int) {
	s.restaurants[i], s.restaurants[j] = s.restaurants[j], s.restaurants[i]
}

func (s *restaurantDBSorter) Less(i, j int) bool {
	return s.by(s.restaurants[i], s.restaurants[j])
}

// IsInRadius checks if a given restaurant is inside a radius compared to a point signified by coordinates
func (restaurant *RestaurantDB) IsInRadius(lat, lon float64, radiusParam string) bool {
	if radiusParam == "ignore" {
		return true
	}
	radius, errRad := strconv.ParseFloat(radiusParam, 64)
	if errRad != nil {
		// default value
		radius = 1000
	}
	distance := coordinates.Haversine(lat, lon, restaurant.Lat, restaurant.Lon)
	return distance <= radius
}

// IsInPriceRange checks if a given restaurant is inside a provided price range
func (restaurant *RestaurantDB) IsInPriceRange(priceRangeString string) bool {
	if priceRangeString == "" {
		return true
	}
	priceRanges := strings.Split(priceRangeString, ",")
	for _, priceRange := range priceRanges {
		replacer := strings.NewReplacer(" ", "", "Kč", "", "+", "-")
		cleanedPriceRange := replacer.Replace(restaurant.PriceRange)
		if cleanedPriceRange == priceRange {
			return true
		}
	}
	return false
}

// IsInDistrict checks if a restaurant is located in a given district
func (restaurant *RestaurantDB) IsInDistrict(districtString string) bool {
	if districtString == "" {
		return true
	}
	replacer := strings.NewReplacer(" ", "")
	districtString = replacer.Replace(districtString)
	districts := strings.Split(districtString, ",")
	for _, district := range districts {
		if replacer.Replace(restaurant.District) == strings.Title(district) {
			return true
		}
	}
	return false
}

// HasCuisines checks if a restaurant has food from a given cuisine
func (restaurant *RestaurantDB) HasCuisines(cuisinesString string) bool {
	if cuisinesString == "" {
		return true
	}
	cuisinesString = strings.Replace(cuisinesString, " ", "", -1)
	cuisines := strings.Split(cuisinesString, ",")
	for _, cuisine := range cuisines {
		if !scraper.SliceContains(restaurant.Cuisines, strings.Title(cuisine)) {
			return false
		}
	}
	return true
}

// CheckDB checks if the db is set up and initializes everything that is not set up yet
func CheckDB() {
	var table string
	conn, err := GetConn()
	if err != nil {
		log.Println("Make sure the DB_DSN environment variable is set")
		log.Fatal(err)
	}
	defer conn.Close()
	err = conn.Get(&table, "SELECT table_name FROM information_schema.tables WHERE table_name=$1", "restaurants")
	if err == sql.ErrNoRows {
		log.Println("No restaurants table found, creating")
		_, err = conn.Exec("CREATE EXTENSION pg_trgm;")
		_, err = conn.Exec("CREATE EXTENSION unaccent;")
		_, err = conn.Exec(restaurantsSchema)
		if err != nil {
			log.Fatal(err)
		}
		err = storeRestaurants(conn)
		if err != nil {
			log.Println("Couldn't store restaurants")
			log.Fatal(err)
		}
	}
	err = conn.Get(&table, "SELECT table_name FROM information_schema.tables WHERE table_name=$1", "restaurateur_users")
	if err == sql.ErrNoRows {
		log.Println("No restaurateur_users table found, creating")
		_, err = conn.Exec(restaurateurUsersSchema)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Database ready")
}

// GetConn fetches a connection to the db
func GetConn() (*sqlx.DB, error) {
	dbDSN := os.Getenv("DB_DSN")
	conn, err := sqlx.Connect("postgres", dbDSN)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func storeRestaurants(conn *sqlx.DB) error {
	restaurants, err := scraper.GetRestaurants("praha")
	if err != nil {
		return err
	}

	for _, r := range restaurants {
		err := insert(r, conn)
		if err != nil {
			return err
		}
	}

	return nil
}

func insert(r *scraper.Restaurant, db *sqlx.DB) error {
	stmt, err := db.Prepare(`INSERT INTO restaurants (name, address, district, images,
								cuisines, price_range, rating, url, phone_number, lat, lon,
								vegan, vegetarian, gluten_free, weekly_menu, opening_hours, takeaway, delivery_options)
								VALUES
								($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)`)
	if err != nil {
		return err
	}

	WeeklyMenu, _ := json.Marshal(r.WeeklyMenu)
	OpeningHours, _ := json.Marshal(r.OpeningHours)

	_, err = stmt.Exec(r.Name,
		r.Address,
		r.District,
		pq.Array(r.Images),
		pq.Array(r.Cuisines), r.PriceRange,
		r.Rating,
		r.URL,
		r.PhoneNumber,
		r.Lat,
		r.Lon,
		r.Vegan,
		r.Vegetarian,
		r.GlutenFree,
		WeeklyMenu,
		OpeningHours,
		r.Takeaway,
		pq.Array(r.DeliveryOptions))

	return err
}

func loadRestaurants(conn *sqlx.DB) ([]*RestaurantDB, error) {
	var restaurants []*RestaurantDB
	err := conn.Select(&restaurants, `SELECT * FROM restaurants`)

	return restaurants, err
}

// GetUserByID returns a user struct based on an ID
func GetUserByID(id int) (*User, error) {
	user := &User{}
	conn, err := GetConn()
	if err != nil {
		return user, err
	}
	err = conn.QueryRowx(`SELECT name, email, password FROM restaurateur_users where id=$1`, id).StructScan(user)
	return user, err
}

// SaveUser saves a user to the db
func SaveUser(user *User) error {
	conn, err := GetConn()
	if err != nil {
		return err
	}
	stmt, err := conn.Prepare("INSERT INTO restaurateur_users (name, email, password) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Name, user.Email, user.Password)
	return err
}

// DeleteUser deletes a user from the db based on an ID
func DeleteUser(id int) error {
	conn, err := GetConn()
	if err != nil {
		return err
	}
	stmt, err := conn.Prepare("DELETE from restaurateur_users WHERE id=$1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}

// GetUserByEmail returns a user struct based on an email
func GetUserByEmail(email string) (*UserDB, error) {
	user := &UserDB{}
	conn, err := GetConn()
	if err != nil {
		return user, err
	}
	err = conn.QueryRowx("SELECT * FROM restaurateur_users where email=$1", email).StructScan(user)
	return user, err
}

// UpdateOne updates a field in the user table
func UpdateOne(column, updateValue string, id int) error {
	conn, err := GetConn()
	if err != nil {
		return err
	}
	stmt, err := conn.Prepare(fmt.Sprintf("UPDATE restaurateur_users set %s=$1 WHERE id=$2", column))
	if err != nil {
		return err
	}
	_, err = stmt.Exec(updateValue, id)
	return err
}

// GetRestaurantArrByID returns a single restaurant but in an array
func GetRestaurantArrByID(id int) ([]*RestaurantDB, error) {
	var restaurant []*RestaurantDB
	queryString := "SELECT * FROM restaurants where id=$1"
	conn, err := GetConn()
	if err != nil {
		return restaurant, err
	}
	err = conn.Select(&restaurant, queryString, id)
	if err != nil {
		return restaurant, err
	}
	return restaurant, nil
}

// GetDBRestaurants returns an array of restaurants that satisfy given criteria
func GetDBRestaurants(params url.Values) ([]*RestaurantDB, error) {
	var restaurants []*RestaurantDB
	conn, err := GetConn()
	if err != nil {
		return restaurants, err
	}
	defer conn.Close()
	var andParams = [...]string{"vegetarian", "vegan", "gluten-free", "takeaway"}
	var nullParams = [...]string{"delivery-options"}
	var queries []string
	var orderBy = ""
	pgQuery := "SELECT * from restaurants"
	paramCtr := 1
	var values []interface{}
	for _, param := range andParams {
		_, ok := params[param]
		if ok && params.Get(param) != "false" {
			param = strings.Replace(param, "-", "_", -1)
			pgParam := fmt.Sprintf("%s=$%d", param, paramCtr)
			queries = append(queries, pgParam)
			value := params.Get(param)
			if value == "" {
				values = append(values, true)
			} else {
				values = append(values, value)
			}
			paramCtr++
		}
	}
	parameterArr, ok := params["search-name"]
	pgParam := fmt.Sprintf("(unaccent(name) %% unaccent($%d))", paramCtr)
	searchField := "name"
	if !ok {
		parameterArr, ok = params["search-address"]
		searchField = "address"
		pgParam = fmt.Sprintf("(regexp_replace(unaccent(address), '[[:digit:]/]', '', 'g') %% unaccent($%d)) ", paramCtr)
	}
	if ok {
		searchString := parameterArr[0]
		queries = append(queries, pgParam)
		values = append(values, searchString)
		orderBy = fmt.Sprintf(" ORDER BY SIMILARITY(unaccent(%s), unaccent($%d)) DESC", searchField, paramCtr)
		searchStringLen := len(searchString)
		if searchStringLen < 3 {
			_, err = conn.Exec("SELECT set_limit(0.1)")
		} else if searchStringLen < 5 {
			_, err = conn.Exec("SELECT set_limit(0.2)")
		} // else keep default of 0.3
		paramCtr++
	}
	for _, param := range nullParams {
		_, ok := params[param]
		if ok {
			param = strings.Replace(param, "-", "_", -1)
			pgParam := fmt.Sprintf("%s IS NOT NULL", param)
			queries = append(queries, pgParam)
		}
	}
	if len(queries) > 0 {
		pgQuery += " WHERE "
	}
	pgQuery += strings.Join(queries, " AND ") + orderBy
	err = conn.Select(&restaurants, pgQuery, values...)
	if err != nil {
		return restaurants, err
	}
	return restaurants, nil
}
