package main

import (
	"database/sql"
	"encoding/json"
	"github.com/AgiliaErnis/restaurateur/backend/scraper"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
	"os"
)

const (
	SCHEMA = `CREATE TABLE restaurants (
		         id SERIAL NOT NULL PRIMARY KEY,
             name TEXT NOT NULL,
					 	 address TEXT NOT NULL,
					 	 images TEXT,
					 	 cuisines TEXT,
					 	 price_range TEXT,
					 	 rating TEXT,
					 	 url TEXT,
					 	 phone_number TEXT,
					 	 lat NUMERIC,
					 	 lon NUMERIC,
					 	 vegan BOOLEAN,
					 	 vegetarian BOOLEAN,
						 weekly_menu JSON
           );`
)

func dbCheck(conn *sqlx.DB) error {
	var table string
	err := conn.Get(&table, "SELECT table_name FROM information_schema.tables WHERE table_name=$1", "restaurants")
	if err == sql.ErrNoRows {
		log.Println("No table found, creating")
		_, err = conn.Exec(SCHEMA)
	}

	return err
}

func dbInitialise() (*sqlx.DB, error) {
	var DB_DSN = os.Getenv("DB_DSN")
	conn, err := sqlx.Connect("postgres", DB_DSN)
	if err != nil {
		return nil, err
	}

	err = dbCheck(conn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func StoreRestaurants(conn *sqlx.DB) error {
	restaurants, err := scraper.GetRestaurants("chodov")
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
	stmt, err := db.Prepare(`INSERT INTO restaurants
								   (name, address, images, cuisines, price_range, rating, url, phone_number, lat, lon, vegan, vegetarian, weekly_menu)
									 VALUES
								   ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`)
	if err != nil {
		return err
	}

	WeeklyMenu, _ := json.Marshal(r.WeeklyMenu)

	_, err = stmt.Exec(r.Name,
		r.Address,
		pq.Array(r.Images),
		pq.Array(r.Cuisines),
		r.PriceRange,
		r.Rating,
		r.URL,
		r.PhoneNumber,
		r.Lat,
		r.Lon,
		r.Vegan,
		r.Vegetarian,
		WeeklyMenu)

	return err
}
