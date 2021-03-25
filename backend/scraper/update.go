package scraper

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"

	"github.com/lib/pq"
)

func StoreRestaurants(dsn string) error {
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer conn.Close()


	restaurants, err := GetRestaurants("praha")
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

func insert(r *Restaurant, db *sql.DB) error {
	stmt, err := db.Prepare(`INSERT INTO restaurants VALUES 
								   (name, address, images, cuisines, price_range, rating, url, phone_number, lat, lon, vegan, vegetarian, weekly_menu)
								   ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`)
	if err != nil {
		return err
	}


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
		jsonpbSerializableMap(r.WeeklyMenu))

	return err
}

type jsonpbSerializableMap map[string]string

func (sm *jsonpbSerializableMap) Value() (driver.Value, error) {
	return json.Marshal(sm)
}
