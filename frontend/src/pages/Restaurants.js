import React, { useState, useEffect } from 'react';
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

export default function Restaurants() {
  const { customThemes, customStyles } = SelectStyle();
  const { sortOptions, setSortResultHandler, } = SelectLogic();

  const [priceRange, setPriceRange] = useState([]);
  const [featured, setFeatured] = useState([]);

  const [restaurants, setRestaurants] = useState([]);

  const [currentPage, setCurrentPage] = useState(1);
  const [restaurantsPerPage] = useState(5);

  const indexOfLastRestaurant = currentPage * restaurantsPerPage;
  const indexOfFirstRestaurant = indexOfLastRestaurant - restaurantsPerPage;

  const paginate = (pageNumber) => {
    setCurrentPage(pageNumber)
  }

  const handlePriceRangeFilters = (filters) => {
    const newFilters = [...filters]
    setPriceRange(newFilters)
  }

  const handleFeaturedFilters = (filters) => {
    const newFilters = [...filters]
    setFeatured(newFilters)
  }

 const showFilteredResults = () => {
    var pragueCollegeRestaurants = "/prague-college/restaurants?"
    if (priceRange.length === 0 & featured.length === 0) {
      return pragueCollegeRestaurants
    }
    if (priceRange.length !== 0 & featured.length !== 0) {
      return pragueCollegeRestaurants +
        "price-range=" + priceRange + "&" + featured.join("&")
    }
    if (priceRange.length !== 0) {
      return pragueCollegeRestaurants +
        "price-range=" + priceRange
    }
    if (featured.length !== 0) {
        return pragueCollegeRestaurants + featured.join("&")
   }
  }

  var path = showFilteredResults();

  useEffect(() => {
    fetch(`${path}`).then(response => response.json()).then(
      json => setRestaurants(json.Data))
      paginate(1);
  }, [path])

  if (restaurants !== null) {
    var currentRestaurants = restaurants.slice(
      indexOfFirstRestaurant, indexOfLastRestaurant);
  }
  else {
    currentRestaurants = null
  }

  return (
    <>
      <Navbar/>
      <div className="restaurants-hero-container">
        <VerticalFilter
          handlePriceRangeFilters={filters =>
            handlePriceRangeFilters(filters, "arrayOfPriceRanges")}
          handleFeaturedFilters={filters =>
            handleFeaturedFilters(filters, "arrayOfFeatured")}
        />
        <div className="restaurant-cards-container">
          <div className="restaurant-cards-header">
            <h1>Restaurants around Prague College</h1>
            <Select
              defaultValue="Sort by"
              options={sortOptions}
              styles={customStyles}
              theme={customThemes}
              onChange={setSortResultHandler}
              className="sort"
              placeholder="Sort by"
              isSearchable
            />
          </div>
          {restaurants ?
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
                      filteredRestaurant.Cuisines.length - 1)
                  { return cuisine }
                  else { return cuisine + ","}
                  }) : "Cuisines are not available"}
                address={filteredRestaurant.Address}
                district={filteredRestaurant.District}
                price={filteredRestaurant.PriceRange}
                takeaway={filteredRestaurant.Takeaway}
                delivery={filteredRestaurant.DeliveryOptions}
              />
            })
            : <h1 className="error">
              There are no results for these filters
              </h1>
          }
        </div>
      </div>
      {restaurants &&
        <RestaurantPagination
          restaurantsPerPage={restaurantsPerPage}
          totalRestaurants={restaurants.length}
          paginate={paginate} />}
    </>
  );
}