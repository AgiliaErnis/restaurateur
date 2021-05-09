import React, { useEffect, useState, useContext} from 'react';
import { Link } from 'react-router-dom';
import { Button } from '../button/Button';
import './Searchbox.css';
import Select from 'react-select'
import MobileNavbar from '../navbar/MobileNavbar'
import AdjustSearchbox from './AdjustSearchbox';
import SearchboxStyle from './SelectStyle';
import SelectLogic from './SelectLogic';
import { ImagePlaceHolder } from '../restaurants/PhotoSlider/ImagePlaceHolder'
import { UserContext } from '../../UserContext';

function Searchbox() {
  const { inputClassName, selectClassName,
          searchSize, searchIconClassname } = AdjustSearchbox();
  const { customStyles, customThemes } = SearchboxStyle();
  const { searchResult, searchOptions,
          setSearchResultHandler } = SelectLogic();
  const { button, showSearch } = MobileNavbar();
  const { setChosenRestaurant, setGeneralSearchPath,
    clickOnScreen, setClickOnScreen } = useContext(UserContext);

  const [suggestedRestaurants, setSuggestedRestaurants] = useState([])
  const [input, setInput] = useState("")
  const searchByPath = (searchResult !== "location" ? "name=" : "address=")

  useEffect(() => {
    showSearch();
  }, [showSearch]);

  window.addEventListener('resize', showSearch);

  const handleChange = e => {
    var userInput = e.target.value;
    setInput(userInput);
    setClickOnScreen(false)
  }

  const generalSearch = (searchResult !== "location" ?
    `search-name=${input}` : `search-address=${input}`)

  const handleSearch = () => {
    setGeneralSearchPath(generalSearch);
    setClickOnScreen(true)
    setChosenRestaurant(false)
  }

  useEffect(() => {
    fetch(`/autocomplete?${searchByPath}${input}`).then(response =>
      response.json()).then(
        json => setSuggestedRestaurants(json.Data))
  }, [input,searchByPath])

  return (
    <div className="search-box">
      <div className="main-search-container">
        {button && <i className={searchIconClassname}/> }
        <input
          type="text"
          className={inputClassName}
          placeholder={searchResult === "name" ?
            "Type restaurant, cafe... name" :
            (searchResult === "location" ?
              "Type address or nearby place of restaurant" :
              "Type restaurant, cafe... name")}
          onChange={handleChange}
        />
        {input.length !== 0 &&
          <div className={clickOnScreen !== false ?
          "suggested-restaurants-hidden"
          :
          suggestedRestaurants.length < 3 ?
            "suggested-restaurants-container"
            :
            "suggested-restaurants-container scroll"}>
            {suggestedRestaurants.length !== 0 ?
              suggestedRestaurants.map(restaurant => {
                return <Link to='/restaurants' style={{textDecoration: "none"}}>
                <div className="suggested-restaurant"
                  onClick={() => setChosenRestaurant(restaurant.ID)}>
                  <div className="restaurant-image">
                      <img src={restaurant.Image !== "" ?
                        restaurant.Image
                        :
                        ImagePlaceHolder}
                        alt="suggested-restaurant" />
                  </div>
                  <div className="text">
                    <p className="restaurant-name">
                      {restaurant.Name}
                    </p>
                    <p className="restaurant-address">
                      {restaurant.Address}, {restaurant.District}
                    </p>
                  </div>
                </div></Link>
              }) : <p className="no-results">No restaurants found
              </p>}
          </div>
          }
      </div>
      <Select
        defaultValue={searchOptions[0]}
        options={searchOptions}
        styles={customStyles}
        theme={customThemes}
        onChange={setSearchResultHandler}
        className={selectClassName}
        placeholder="Search by"
        isSearchable
      />
      <Link to='/restaurants'>
        <Button
          buttonSize={searchSize}
            buttonStyle='btn--search'
            onClick={handleSearch}
        >
            Search
        </Button>
      </Link>
    </div>
  )
}

export default Searchbox;