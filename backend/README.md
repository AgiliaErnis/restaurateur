# WIP

## Starting the server

`$ cd backend `

`$ go build `

`$ ./backend`

## Database set-up

postgres >= 13

Environment variable DB_DSN.

Example:
`export DB_DSN='user=postgres dbname=postgres password=postgres sslmode=disable'`

**Importing the restaurants table**

There is a dump of the restaurants table in the backend folder `restaurants.sql`.

In the command line the table can be imported in a following way:

`$ psql < restaurants.sql`

**Dumping the restaurants table**

Command line example:

`$ pg_dump postgres -U postgres -h localhost -t restaurants > restaurants.sql`


---

For now the server listens on port 8080, later the port should be configurable.

## Currently implemented

HTTP server with two endpoints

| Method | Path                        | Handler       |
|--------|-----------------------------|---------------|
| GET    | /prague-college/restaurants | pcRestaurantsHandler |
| GET | /restaurants | restaurantsHandler

## Query parameters

### radius

Radius (in meters) of the area around a provided or pre-selected starting point. Restaurants in this area will be returned.
Radius can be ignored when specified. When no radius is provided, a default value of 1000 meters is used.

**Examples**

`radius=ignore`

`radius=800`

### address

Starting point for a search in a given radius.

**Example**

 `address=Mánesova 13 Praha 2`

### lat

Latitude in degrees.

**Example**

`lat=50.0785714`

### lon

Longitude in degrees.

**Example**

`lon=14.4400922`

### cuisine

Filters restaurants based on a list of cuisines, separated by commas.
A restaurant will be returned only if it satisfies all provided cuisines.

Available cuisines:
American, Italian, Asian, Indian, Japanese, Vietnamese, Spanish, Mediterranean, French, Thai, Mexican, International, Czech, English, Balkan, Brazil, Russian, Chinese, Greek, Arabic, Korean

**Examples**

`cuisine=Czech,English`

`cuisine=International`

### price-range

Filters restaurants based on a list of price ranges, separated by commas.
A restaurant will be returned if it satisfies at least one provided price range.

Available price ranges:

* 0-300
* 300-600
* 600-

**Examples**

`price-range=0-300,600-`

`price-range=300-600`


### district

Filters restaurants based on a list districts, separated by commas.
A restaurant will be returned if it is in one of the provided districts.

**Example**

`district=Praha 2, Praha 10`

`district=Praha 3`

### vegetarian

Filters out all non vegetarian restaurants.

**Example**

`vegetarian`

### vegan

Filters out all non vegan restaurants.

**Example**

`vegan`

### gluten-free

Filters out all non gluten free restaurants.

**Example**

`gluten-free`

### takeaway

Filters out all restaurants that don't have a takeaway option.

**Example**

 `takeaway`

### delivery-options

Filters out restaurants without a delivery option.

**Example**

`delivery-options`

---

## Example JSON responses
```json
{
  "Status": 200,
  "Msg": "Success",
  "Data": [
    {
      "ID": 24,
      "Name": "SOVA",
      "Address": "Balbínova 4",
      "District": "Praha 2",
      "Images": [
        "https://www.restu.cz/ir/restaurant/7ff/7ffb3fdcd62f3183bcc1dc11fb5d3dac--cxc400.jpg",
        "https://www.restu.cz/ir/restaurant/793/79342c1365cee1d7a476f0f9358f02b2--cxc400.jpg"
      ],
      "Cuisines": [
        "Czech",
        "International"
      ],
      "PriceRange": "0 - 300 Kč",
      "Rating": "4.7",
      "URL": "http://www.rest-sova.com",
      "PhoneNumber": "+420 222 541 451",
      "Lat": 50.0768013,
      "Lon": 14.43389,
      "Vegan": false,
      "Vegetarian": false,
      "GlutenFree": false,
      "WeeklyMenu": "null",
      "OpeningHours": "{\"Friday\":\"17:00 - 23:00\",\"Monday\":\"Zavřeno\",\"Saturday\":\"11:00 - 23:00\",\"Sunday\":\"11:00 - 23:00\",\"Thursday\":\"17:00 - 23:00\",\"Tuesday\":\"Zavřeno\",\"Wednesday\":\"Zavřeno\"}",
      "Takeaway": false,
      "DeliveryOptions": null
    },
    {
      "ID": 43,
      "Name": "Amunì",
      "Address": "Vinohradská 1183/83",
      "District": "Praha 2",
      "Images": [
        "https://www.restu.cz/ir/restaurant/91e/91e404d2b08096767029306576a39c1f--cxc400.jpg",
        "https://www.restu.cz/ir/restaurant/9d7/9d762f29816282786ca51ee5febef520--cxc400.jpg"
      ],
      "Cuisines": [
        "Italian"
      ],
      "PriceRange": "300 - 600 Kč",
      "Rating": "4.7",
      "URL": "http://amunipraha.cz/",
      "PhoneNumber": "+420 606 068 765",
      "Lat": 50.0771978,
      "Lon": 14.4464371,
      "Vegan": false,
      "Vegetarian": false,
      "GlutenFree": false,
      "WeeklyMenu": "null",
      "OpeningHours": "{\"Friday\":\"12:00 - 22:00\",\"Monday\":\"12:00 - 22:00\",\"Saturday\":\"12:00 - 22:00\",\"Sunday\":\"12:00 - 22:00\",\"Thursday\":\"12:00 - 22:00\",\"Tuesday\":\"12:00 - 22:00\",\"Wednesday\":\"12:00 - 22:00\"}",
      "Takeaway": true,
      "DeliveryOptions": [
        "https://wolt.com/cs/cze/prague/restaurant/amuni?fbclid=IwAR3y_q5PwgLCDQ9rUwiSeo72AbGHg_wmoCqqTDZUM3zLrzxrodWYbJ5cz2A",
        "https://onemenu.cz/menu/AMUNI"
      ]
    }
  ]
}
```
```json
{
  "Status": 400,
  "Msg": "Invalid coordinates(Lat: wrong1, Lon: wrong2)",
  "Data": null
}
```
```json
{
  "Status": 404,
  "Msg": "Invalid endpoint: /test",
  "Data": null
}
```
