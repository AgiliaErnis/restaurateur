import React, { useEffect, useState, useContext, useRef} from 'react';
import { Link } from 'react-router-dom';
import { Button } from '../button/Button';
import './Searchbox.css';
import Select from 'react-select'
import MobileNavbar from '../navbar/MobileNavbar'
import AdjustSearchbox from './AdjustSearchbox';
import SearchboxStyle from './SelectStyle';
import SelectLogic from './SelectLogic';
import { UserContext } from '../../UserContext';
import { Redirect } from 'react-router-dom';
import SuggestedRestaurantItem from './SuggestedRestaurantItem'

function Searchbox() {
   const { inputClassName, selectClassName,
          searchSize, searchIconClassname } = AdjustSearchbox();
  const { customStyles, customThemes } = SearchboxStyle();
  const { searchResult, searchOptions,
          setSearchResultHandler } = SelectLogic();
  const { button, showSearch } = MobileNavbar();
  const { setChosenRestaurant, setGeneralSearchPath, setPragueCollegePath } = useContext(UserContext);

  const [suggestedRestaurants, setSuggestedRestaurants] = useState([])
  const [input, setInput] = useState("")
  const searchByPath = (searchResult !== "location" ? "name=" : "address=")
  const [enterPressed, setEnterPressed] = useState(false)
  const [isVisible, setVisibility] = useState(false)
  const [cursor, setCursor] = useState(-1)

  const showSuggestions = () => {
    setVisibility(true);
    if (input === "") {
      setInput("a")
    }
  };

  const hideSuggestions = () => setVisibility(false);

  const searchContainer = useRef(null);
  const searchResultRef = useRef(null);

  useEffect(() => {
    showSearch();
  }, [showSearch]);

  window.addEventListener('resize', showSearch);

  const handleChange = e => {
    var userInput = e.target.value;
    setInput(userInput);
    setVisibility(true);
  }

  const generalSearch = (searchResult !== "location" ?
    `search-name=${input}` : `search-address=${input}`)

  const handleSearch = () => {
    setGeneralSearchPath(generalSearch);
    setChosenRestaurant(false)
    setVisibility(false)
    setPragueCollegePath(false)
  }

  const handleClickOutside = (event) => {
    if (searchContainer.current &&
      !searchContainer.current.contains(event.target)) {
      hideSuggestions();
    }
  }

  useEffect(() => {
    window.addEventListener("mousedown", handleClickOutside)
    return () => {
      window.removeEventListener("mousedown", handleClickOutside)
    }
  })

  useEffect(() => {
    if (cursor < 0 || cursor > suggestedRestaurants.length
      ||
      !searchResultRef) {
      return () => {}
    }

    const scrollSuggestedRestaurants = position => {
    if (cursor > -1) {
      searchResultRef.current.children[cursor].scrollIntoView({
        top: position,
        smooth: true
      })
      window.scrollTo(0,0)
      }
    }

    let listItems = Array.from(searchResultRef.current.children);
    listItems[cursor] && scrollSuggestedRestaurants(listItems[cursor].offsetTop);
  },[cursor,suggestedRestaurants])

  useEffect(() => {
    if (isVisible) {
      fetch(`${process.env.REACT_APP_PROXY}autocomplete?${searchByPath}${input}`).then(response =>
        response.json()).then(
          json => setSuggestedRestaurants(json.Data))
    }
  }, [input, searchByPath,isVisible])

  const keyboardNavigation = (e) => {
    if (e.key === "ArrowDown") {
      isVisible ?
        setCursor(c => (c < suggestedRestaurants.length - 1 ? c + 1 : c))
        : showSuggestions();
    }

    if (e.key === "ArrowUp") {
      setCursor(c => (c > -1 ? c - 1 : c));
    }

    if (e.key === "Enter") {
      if (cursor > -1) {
        setChosenRestaurant(suggestedRestaurants[cursor].ID)
      } else {
        setGeneralSearchPath(generalSearch);
        setChosenRestaurant(false)
      }
    hideSuggestions();
    setEnterPressed(true)
    setPragueCollegePath(false)
    }

    if (e.key === "Escape") {
      hideSuggestions();
    }
  }

  return (
    <div className="search-box" ref={searchContainer}>
      <div className="main-search-container">
        {button && <i className={searchIconClassname} />}
        <input
          type="text"
          className={inputClassName}
          placeholder={searchResult === "name" ?
            "Type restaurant, cafe... name" :
            (searchResult === "location" ?
              "Type address or nearby place of restaurant" :
              "Type restaurant, cafe... name")}
          onChange={handleChange}
          onKeyDown={e => keyboardNavigation(e)}
          onClick={showSuggestions}
        />{enterPressed && <Redirect push to="/restaurants" />}
        {input.length !== 0 &&
          <div className={
            isVisible === false ? "suggested-restaurants-hidden" :
            (suggestedRestaurants !== null &&
              suggestedRestaurants.length < 3 ?
              "suggested-restaurants-container"
              :
              "suggested-restaurants-container scroll")}
            >
          <ul ref={searchResultRef}>
            {suggestedRestaurants !== null ?
              suggestedRestaurants.map(restaurant => {
                return <Link to='/restaurants' style={{ textDecoration: "none" }}>
                  <SuggestedRestaurantItem
                    key={restaurant.ID}
                    image={restaurant.Image}
                    name={restaurant.Name}
                    address={restaurant.Address}
                    district={restaurant.District}
                    onSelectItem={() => {
                      hideSuggestions();
                      setChosenRestaurant(restaurant.ID);
                      setPragueCollegePath(false)
                    }}
                    isHighlighted={cursor === suggestedRestaurants.indexOf(restaurant) ?
                      true
                      :
                      false}
                    {...restaurant}
                    />
                </Link>
              })
              :
              <p className="no-results">
                No restaurants found
              </p>
            }
          </ul>
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