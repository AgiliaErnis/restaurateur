import React, { useState, useEffect, useContext } from 'react';
import { VerticalFilter } from '../components/filtration/VerticalFilter';
import RestaurantItem from '../components/restaurants/RestaurantItem';
import Select from 'react-select'
import Navbar from '../components/navbar/Navbar';
import SelectStyle from '../components/search/SelectStyle';
import SelectLogic from '../components/search/SelectLogic';
import './Restaurants.css'
import RestaurantPagination from
  '../components/restaurants/pagination/Pagination';
import { ImagePlaceHolder } from
  '../components/restaurants/PhotoSlider/ImagePlaceHolder';
import { UserContext } from '../UserContext';

export default function Restaurants() {
  const { customThemes, customStyles } = SelectStyle();
  const { sortOptions, setSortResultHandler } = SelectLogic();
  const pragueCollegePath = useContext(UserContext)
  const clickedDistrict = useContext(UserContext)
  const clickedSuggestion = useContext(UserContext)
  const checkedDistance = useContext(UserContext)

  const [checkedFilters, setCheckedFilters] = useState([]);
  const [restaurants, setRestaurants] = useState([]);
  const [unsortedRestaurants, setUnsortedRestaurants] = useState([]);

  const [currentPage, setCurrentPage] = useState(1);
  const [restaurantsPerPage] = useState(5);

  const indexOfLastRestaurant = currentPage * restaurantsPerPage;
  const indexOfFirstRestaurant = indexOfLastRestaurant - restaurantsPerPage;

  const paginate = (pageNumber) => {
    setCurrentPage(pageNumber)
  }

  const handlecheckedFilters = (filters) => {
    const newFilters = [...filters]
    setCheckedFilters(newFilters)
  }
  const arrayOfPathValues = checkedFilters.filter(filter =>
    filter.checkedOptions.length !== 0).map(noneEmptyFilter => {
      if (noneEmptyFilter.category === "other") {
        return noneEmptyFilter.checkedOptions.join("&")
      }
      else {
        return noneEmptyFilter.category + "=" + noneEmptyFilter.checkedOptions
      }
    }
  )

  const showFilteredResults = () => {
    if (pragueCollegePath.pragueCollegePath === true) {
      var pragueCollegeRestaurants =
        `/prague-college/restaurants?radius=${checkedDistance.checkedDistance}&`

      if (clickedDistrict.clickedDistrict !== false) {
        pragueCollegeRestaurants += `district=${clickedDistrict.clickedDistrict}`
      }

      if (clickedSuggestion.clickedSuggestion !== false) {
        if (clickedSuggestion.clickedSuggestion === "vegetarian" ||
          clickedSuggestion.clickedSuggestion === "gluten-free") {
          pragueCollegeRestaurants += `${clickedSuggestion.clickedSuggestion}`
        } else {
          pragueCollegeRestaurants +=
            `cuisine=${clickedSuggestion.clickedSuggestion}`
        }
      }
      return pragueCollegeRestaurants + arrayOfPathValues.join("&")
    } else {
        var path = "/restaurants?radius=ignore&"
        if (clickedDistrict.clickedDistrict !== false) {
          path += `district=${clickedDistrict.clickedDistrict}`
      }
      if (clickedSuggestion.clickedSuggestion !== false) {
        if (clickedSuggestion.clickedSuggestion === "vegetarian"
          ||
          clickedSuggestion.clickedSuggestion === "gluten-free") {
            path += clickedSuggestion.clickedSuggestion
          } else {
              path += `cuisine=${clickedSuggestion.clickedSuggestion}`
        }
      }
      return path + arrayOfPathValues.join("&")
    }
  }

  const path = showFilteredResults();

  function sortRestaurants(sortParameter) {
    let restaurantsCopy
    // Store a copy of restaurants, because if for example all restaurants without
    // ratings would be removed, sorted by rating and then sorted by price,
    // it could happen that a restaurant
    // with a valid price would be removed and now it would be missing.
    if (unsortedRestaurants.length > 0) {
      restaurantsCopy = unsortedRestaurants
    } else {
      restaurantsCopy = restaurants
      setUnsortedRestaurants(restaurants)
    }
    let tempRestaurants = []
    if (sortParameter === undefined || restaurantsCopy === null) {
      return
    }
    if (sortParameter === "rating") {
      // Remove all restaurants without a rating
      for (let restaurant of restaurantsCopy) {
        if (restaurant["Rating"] !== "") {
          tempRestaurants.push(restaurant)
        }
      }
      tempRestaurants.sort((a, b) => (parseFloat(a["Rating"]) < parseFloat(b["Rating"]) ? 1 : -1))
    } else {
      // Remove all restaurants without a price range
      for (let restaurant of restaurantsCopy) {
        if (restaurant["PriceRange"] !== "Not available") {
          tempRestaurants.push(restaurant)
        }
      }
      if (sortParameter === "price-ascending") {
        tempRestaurants.sort((a, b) => (a["PriceRange"] > b["PriceRange"] ? 1 : -1))
      } else if (sortParameter === "price-descending") {
        tempRestaurants.sort((a, b) => (a["PriceRange"] < b["PriceRange"] ? 1 : -1))
      }
    }
    return tempRestaurants
  }

  function sortHandler(e) {
    setSortResultHandler(e)
    let sortParameter = e.value
    let tempRestaurants = sortRestaurants(sortParameter)
    setRestaurants(tempRestaurants)
    paginate(1);
  }

  useEffect(() => {
    fetch(`${path}`).then(response => response.json()).then(
      json => {
        setUnsortedRestaurants([])
        setRestaurants(json.Data)
      })
      paginate(1);
  }, [path])

  if (restaurants !== null) {
    var currentRestaurants = restaurants.slice(
      indexOfFirstRestaurant, indexOfLastRestaurant);
  } else {
      currentRestaurants = null
  }

  return (
    <>
      <Navbar/>
      <div className="restaurants-hero-container">
        <VerticalFilter
          handlecheckedFilters={filters =>
            handlecheckedFilters(filters, "arrayOfcheckedFilterss")}
        />
        <div className="restaurant-cards-container">
          <div className="restaurant-cards-header">
            <h1>
              {pragueCollegePath.pragueCollegePath === true
                ?
                "Restaurants around Prague College"
                :
                "Restaurants in Prague"
              }
            </h1>
            <Select
              defaultValue="Sort by"
              options={sortOptions}
              styles={customStyles}
              theme={customThemes}
              onChange={sortHandler}
              className="sort"
              placeholder="Sort by"
              isSearchable
            />
          </div>
          {restaurants.length !== 0 ?
            currentRestaurants.map(filteredRestaurant => {
              return <RestaurantItem
                photos={filteredRestaurant.Images.length !== 0 ?
                  filteredRestaurant.Images : ImagePlaceHolder}
                name={filteredRestaurant.Name}
                rating={filteredRestaurant.Rating === "" ?
                  "Rating is not available" : filteredRestaurant.Rating}
                tags={filteredRestaurant.Cuisines !== null ?
                  filteredRestaurant.Cuisines.map((cuisine) => {
                    if (filteredRestaurant.Cuisines.indexOf(cuisine) ===
                      filteredRestaurant.Cuisines.length - 1) {
                      return cuisine
                    } else {
                        return cuisine + ","
                      }
                    })
                    :
                    "Cuisines are not available"}
                address={filteredRestaurant.Address}
                district={filteredRestaurant.District}
                price={filteredRestaurant.PriceRange}
                takeaway={filteredRestaurant.Takeaway}
                delivery={filteredRestaurant.DeliveryOptions}
              />
            })
            :
            <h1 className="error">
              There are no results for these filters
            </h1>
          }
        </div>
      </div>
      {restaurants &&
        <RestaurantPagination
          restaurantsPerPage={restaurantsPerPage}
          totalRestaurants={restaurants.length}
          paginate={paginate}
        />
      }
    </>
  );
}
