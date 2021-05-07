package main

import (
	"database/sql"
	"encoding/json"
	"github.com/AgiliaErnis/restaurateur/backend/scraper"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
	"os"
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

type userDB struct {
	ID       int
	Name     string
	Email    string
	Password string
}

// RestaurantDB struct compatible with postgres
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

func (restaurant *RestaurantDB) isInRadius(lat, lon float64, radiusParam string) bool {
	if radiusParam == "ignore" {
		return true
	}
	radius, errRad := strconv.ParseFloat(radiusParam, 64)
	if errRad != nil {
		// default value
		radius = 1000
	}
	distance := haversine(lat, lon, restaurant.Lat, restaurant.Lon)
	return distance <= radius
}

func (restaurant *RestaurantDB) isInPriceRange(priceRangeString string) bool {
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

func (restaurant *RestaurantDB) isInDistrict(districtString string) bool {
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

func (restaurant *RestaurantDB) hasCuisines(cuisinesString string) bool {
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

func dbCheck() {
	var table string
	conn, err := dbGetConn()
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

func dbGetConn() (*sqlx.DB, error) {
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

func getUserByID(id int) (*user, error) {
	user := &user{}
	conn, err := dbGetConn()
	if err != nil {
		return user, err
	}
	err = conn.QueryRowx(`SELECT name, email, password FROM restaurateur_users where id=$1`, id).StructScan(user)
	return user, err
}

func saveUser(user *user) error {
	conn, err := dbGetConn()
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

func deleteUser(id int) error {
	conn, err := dbGetConn()
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

func getUserByEmail(email string) (*userDB, error) {
	user := &userDB{}
	conn, err := dbGetConn()
	if err != nil {
		return user, err
	}
	err = conn.QueryRowx("SELECT * FROM restaurateur_users where email = $1", email).StructScan(user)
	return user, err
}
