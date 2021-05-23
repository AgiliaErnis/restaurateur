# Restaurateur
Restaurateur is a restaurant recommender that takes
into account all the restaurants located
throughout Prague and recommends the most suitable
ones based on applied filtering preferences.

The recommender also works for a fixed location - Prague College.

## Functionality

### Stage 1

* The recommender works for a fixed location
given in advance ('Prague College local restaurants').

* The recommender provides the following
information about each restaurant:
    * Name
    * Address, District
    * Rating
    * Price Range
    * Cuisines
    * Takeaway and Delivery Options

* Recommendations are being done on fixed filtering references.
Preferences to take into account are:
    * Price Range
    * Vegetarian and Vegan Options
    * Gluten Free Option

### Stage 2

* The recmmender works for all the restaurants
located throughout Prague.

* The recommender provides the following
information about each restaurant:
    * Name
    * Address, District
    * Rating
    * Price Range
    * Cuisines
    * Takeaway and Delivery Options
    * Phone Number

* Filtering Options:

    For Prague College local restaurants:
        * Price Range
        * District
        * Delivery and Takeaway Options
        * Vegetarian and Vegan Options
        * Gluten Free Option
        * Distance

    For all restaurants:
        * Cuisine
        * Price Range
        * District
        * Delivery and Takeaway Options
        * Vegetarian and Vegan Options
        * Gluten Free Option

* Search engine suggests restaurants based on user's input and
 works as an autocomplete.

The restaurants can be searched by the given location or name.

* The recommender can sort retrieved restaurants
 by:
    * Price range - descending
    * Price range - ascending
    * Rating - descending

* The application allows the user to create an account,
 save favourite restaurants and modify account settings as needed.

## Usage

https://user-images.githubusercontent.com/56120787/115972966-c1b25500-a562-11eb-9fd9-ca9e72279b14.mp4

## User Interface

The UI of the Restaurateur web application is fully
based on JavaScript Library – React.

### Running the application

#### Frontend

`$ cd frontend `

`$ npm install `

`$ npm start `

#### Backend

`$ cd backend `

`$ go build `

`$ ./backend `

For more information check out the backend README file - https://github.com/AgiliaErnis/restaurateur/tree/main/backend#readme