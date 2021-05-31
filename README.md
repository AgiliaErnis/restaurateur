# Restaurateur
Restaurateur is a restaurant recommender that takes
into account all the restaurants located
throughout Prague and recommends the most suitable
ones based on applied filtering preferences.

The recommender also works for a fixed location - Prague College.

- [Functionality](#functionality)
- [Usage](#usage)
- [Running the application](#running-the-application)
- [Deployment Information](#deployment-information)

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
    * Weekly Menu

* Filtering Options:

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


### Stage 3

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
    * URL
    * Vegan, Vegeterian and Gluten Free Options
    * Opening Hours

## Usage

### Stage 1

https://user-images.githubusercontent.com/56120787/115972966-c1b25500-a562-11eb-9fd9-ca9e72279b14.mp4

### Stage 2 && Stage 3

### Registration

https://user-images.githubusercontent.com/56120787/119279329-0dc4e800-bc3c-11eb-82a2-b8d955f4d481.mp4

### Authorization

https://user-images.githubusercontent.com/56120787/119279367-55e40a80-bc3c-11eb-8354-1ff9bcb2e307.mp4

### General Usage

https://user-images.githubusercontent.com/56120787/119279240-72cc0e00-bc3b-11eb-94de-9e7aff14596f.mp4

## Running the application

The detailed insturctions for running the application
can be seen in the following README files:

#### Frontend
https://github.com/AgiliaErnis/restaurateur/tree/main/frontend#readme
#### Backend

https://github.com/AgiliaErnis/restaurateur/tree/main/backend#readme

## Deployment Information
The application has two environments - https://test.restaurateur.tech for testing, and https://restaurateur.tech for production builds. Backend API is accessible at https://testapi.restaurateur.tech and https://api.restaurateur.tech.

### Architecture

The application is contained inside 2 Docker containers:

* db - Postgres container with a persistent storage on the server
* app - a container based on NodeJS for both Node and Go parts

There is an nginx reverse proxy set up on the server on ports 80 and 443. It is used for easy SSL and domain configuration.

### Deployment

`docker compose` is used to build the containers on the server. `docker-compose-dev.yml` is used for test.restaurateur.tech, and `docker-compose.yml` for restaurateur.tech.

The deployment process is automated through GitHub actions. GitHub runner connects to the server through SSH and executes the necessary commands. Test environment is released to test.restaurateur.tech on prerelease GitHub event. Production environment is released to restaurateur.tech on release GitHub event. `deploy-prerelese.yml` and `deploy-release.yml` files are the workflows responsible for that.
