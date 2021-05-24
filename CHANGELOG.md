# Changelog

## [0.1.0] - 2021-04-24

### Added

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

## [0.1.1] - 2021-05-06

### Fixed

* The problem of pagination not highlighting the
correct current page was solved.

* Fixed console errors.

### Added

* Added the scroll to top function which is triggered
whenever the vertical filtration or pagination is clicked.

## [0.2.0] - 2021-05-23

### Added

* The recommender provides the following
information about each restaurant:
    * Name
    * Address, District
    * Rating
    * Price Range
    * Cuisines
    * Takeaway and Delivery Options
    * Phone Number
    * Weekly Menu

* Updated filtering options:

    * For Prague College local restaurants:
        * Cuisine
        * Price Range
        * District
        * Delivery and Takeaway Options
        * Vegetarian and Vegan Options
        * Gluten Free Option
        * Weekly Menu
        * Distance

    * For all restaurants:
        * Cuisine
        * Price Range
        * District
        * Delivery and Takeaway Options
        * Vegetarian and Vegan Options
        * Gluten Free Option
        * Weekly Menu

 * Added search engine which suggests restaurants
  based on user's input and works as an autocomlete.
  The restaurants can be searched by the given location or name.

* Added sorting of retrived restaurants with the following options:
    * Price range - descending
    * Price range - ascending
    * Rating - descending

* Implemented user registration/logging in with the
 possibility to save favourite restaurants or modify
 the account settings(password, username, delete account) as needed.
