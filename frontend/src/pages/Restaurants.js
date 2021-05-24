import React, { useState, useEffect, useContext } from 'react';
import { VerticalFilter } from '../components/filtration/VerticalFilter';
import { RestaurantItem } from '../components/restaurants/RestaurantItem';
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
  const { sortOptions, setSortResultHandler,sortResult } = SelectLogic();
  const { chosenRestaurant, generalSearchPath,savedRestaurants } = useContext(UserContext);

  const [checkedFilters, setCheckedFilters] = useState([]);
  const { restaurants, setRestaurants,
          clickedDistrict, clickedSuggestion,
          checkedDistance, pragueCollegePath } = useContext(UserContext)

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
    if (chosenRestaurant !== false) {
      var chosenRestaurantPath = `/restaurant/${chosenRestaurant}`
      return chosenRestaurantPath
    } else if (pragueCollegePath === true) {
        var pragueCollegeRestaurants =
          `/prague-college/restaurants?radius=${checkedDistance}&`

        if (clickedDistrict !== false) {
          pragueCollegeRestaurants += `district=${clickedDistrict}`
        }

        if (clickedSuggestion !== false) {
          if (clickedSuggestion === "vegetarian" ||
            clickedSuggestion === "gluten-free") {
            pragueCollegeRestaurants += `${clickedSuggestion}`
          } else {
            pragueCollegeRestaurants +=
              `cuisine=${clickedSuggestion}`
          }
        }
      return pragueCollegeRestaurants + arrayOfPathValues.join("&") +
        (arrayOfPathValues.length !== 0 ? "&" : "") + `sort=${sortResult}`
    } else {
        var path = "/restaurants?radius=ignore&"
        if (clickedDistrict !== false) {
          path += `district=${clickedDistrict}&`
        } else if (clickedSuggestion !== false) {
            if (clickedSuggestion === "vegetarian"
              ||
              clickedSuggestion === "gluten-free") {
              path += clickedSuggestion + "&"
            } else {
              path += `cuisine=${clickedSuggestion}&`
            }
        } else if (generalSearchPath !== false) {
            path += generalSearchPath + "&"
        }
      return path + arrayOfPathValues.join("&") +
        (arrayOfPathValues.length !== 0 ? "&" : "") + `sort=${sortResult}`
      }
  }

  const path = showFilteredResults();

  useEffect(() => {
    fetch(`http://localhost:8080/${path}`).then(response => response.json()).then(
      json => setRestaurants(json.Data))
    paginate(1);
  }, [path,setRestaurants])

  if (restaurants !== null) {
    var currentRestaurants = restaurants.slice(
      indexOfFirstRestaurant, indexOfLastRestaurant);
  } else {
      currentRestaurants = null
  }


  return (
    <>
      <Navbar/>
      <div className="restaurants-hero-container" >
        <VerticalFilter
          handlecheckedFilters={filters =>
            handlecheckedFilters(filters, "arrayOfcheckedFilterss")}
        />
        <div className="restaurant-cards-container">
          <div className="restaurant-cards-header">
            <h1>
              {pragueCollegePath === true
                ?
                "Restaurants around Prague College"
                :
                "Restaurants in Prague"
              }
            </h1>
            {chosenRestaurant === false &&
              <Select
                defaultValue="Sort by"
                options={sortOptions}
                styles={customStyles}
                theme={customThemes}
                onChange={setSortResultHandler}
                className="sort"
                placeholder="Sort by"
              />}
          </div>
          {restaurants !== null ?
            currentRestaurants.map(filteredRestaurant => {
              return <RestaurantItem
                RestaurantIsSaved={savedRestaurants !== null ?(savedRestaurants.map(index => {
                  if (index.ID === filteredRestaurant.ID)
                    {
                    return true;
                  } else {
                    return false;
                  }})): []}
                ID={filteredRestaurant.ID}
                key={filteredRestaurant.ID}
                photos={filteredRestaurant.Images !== null &&
                  filteredRestaurant.Images.length !== 0 ?
                  filteredRestaurant.Images : ImagePlaceHolder}
                name={filteredRestaurant.Name}
                rating={filteredRestaurant.Rating === "" ?
                  0 : parseFloat(filteredRestaurant.Rating)}
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
                phone={filteredRestaurant.PhoneNumber}
                menu={filteredRestaurant.WeeklyMenu !== "null" ? filteredRestaurant.WeeklyMenu : null}
              />
            })
            :
            <h1 className="error">
              No Restaurants Found
            </h1>
          }
        </div>
      </div>
      {restaurants &&
          <RestaurantPagination
            restaurantsPerPage={restaurantsPerPage}
            totalRestaurants={restaurants.length}
            paginate={paginate}
            page={currentPage}
            />}
    </>
  );
}