import React, { useState, useEffect, useContext } from 'react';
import './RestaurantItem.css'
import Rating from '@material-ui/lab/Rating';
import { PhotoSlider } from './PhotoSlider/PhotoSlider';
import PhoneModal from './PhoneModal';
import { UserContext } from '../../UserContext';
import MenuModal from './MenuModal';

export const RestaurantItem = React.memo((props) => {
  const [click, setClick] = useState(restaurantIsClicked());
  const [savedRestaurant, setSavedRestaurant] = useState(false)
  const [clickOnPhone, setClickOnPhone] = useState(false)
  const [clickOnMenu, setClickOnMenu] = useState(false)
  const [clickOnView, setClickOnView] = useState(false)
  const { setNewSavedRestaurant } = useContext(UserContext)
  const [deleteSavedOne, setDeleteSavedOne] = useState(false)
  const { successfullLogin } = useContext(UserContext)
  
  const openingHours = (props.OpeningHours !== null && JSON.parse(props.OpeningHours))
  const days = [...Object.keys(openingHours)]
  const hours = [...Object.values(openingHours)]

  const handleClick = () => setClick(!click);

  function restaurantIsClicked () {
    switch(window.location.pathname){
      case '/user':
        return true
      default:
        return false;
    }
  }

  function hideSaveBtn () {
    switch(window.location.pathname){
      case '/restaurants':
        return true
      case '/user':
        return false;
      default:
        return null;
    }
  }

  const handleClickOnPhone = () => {
    setClickOnPhone(!clickOnPhone)
  }

  const handleClickOnMenu = () => {
    setClickOnMenu(!clickOnMenu)
  }

  const handleClickOnView = () => {
    setClickOnView(!clickOnView)
  }

  useEffect(() => {
    if (savedRestaurant !== false) {
      const restaurantID = {
        "restaurantID": savedRestaurant
      }
      if (click) {
        const saveRestaurantRequest = {
          method: 'POST',
          body: JSON.stringify(restaurantID),
          credentials: 'include',
          headers: {
            'Content-Type': 'application/json'
          }
        }

        fetch(`${process.env.REACT_APP_PROXY}/auth/user/saved-restaurants`, saveRestaurantRequest)
          .then(response => response.json())
          .then(res => {
            if (res.Status === 200) {
              setNewSavedRestaurant(restaurantID)
            }
          })
      } else {
         const deleteRestaurantRequest = {
          method: 'DELETE',
          body: JSON.stringify(restaurantID),
          credentials: 'include',
          headers: {
            'Content-Type': 'application/json'
          }
        }

        fetch(`${process.env.REACT_APP_PROXY}/auth/user/saved-restaurants`, deleteRestaurantRequest)
          .then(response => response.json())
          .then(res => {
            if (res.Status === 200) {
              setNewSavedRestaurant(restaurantID)
              setDeleteSavedOne(false)
            }
          })
      }

    }
  }, [click, savedRestaurant, setNewSavedRestaurant])

  return (
    <>
      <li className='restaurant_card_item'>
        <div className="content">
          <PhotoSlider slides={props.photos} />
          <div className="restaurant_description">
            <span className="restaurant-name">
              <div className="name-container">{props.name}</div>
              {restaurantIsClicked() &&
                <div className="save-container" onClick={() => {
                  handleClick();
                  setSavedRestaurant(props.ID)
              }}>
                 <p style={{color: "red", fontSize: "13px"}}>remove</p>
              </div>}
              {hideSaveBtn() && successfullLogin &&
                <div className="save-container">
                <i onClick={
                  () => {
                    if (!restaurantIsClicked() &&
                      props.RestaurantIsSaved.indexOf(true) !== -1) {
                      setClick(false)
                      setSavedRestaurant(props.ID)
                      setDeleteSavedOne(true)
                    }
                    else {
                      setSavedRestaurant(props.ID);
                      handleClick();
                      setDeleteSavedOne(false)
                    }
                  }
                }
                  className={!restaurantIsClicked() && props.RestaurantIsSaved.indexOf(true) !== -1 && !deleteSavedOne ? "save-btn-active" : click ? "save-btn-active" : "save-btn"}
                >
                </i>
                <div className={!click ? 'save-on-hover'
                  : 'save-on-hover-hidden'} >
                  Click to Save
                </div>
                <div className={!click ? "saved-hidden" : "saved"}>Saved</div>
              </div>}
            </span>
            <div className="rating-container">
              <Rating name="read-only"
                value={props.rating}
                precision={0.1}
                readOnly
              />
              <span className="rating-num">({props.rating})</span>
            </div>
            <span className="tags">
              {props.tags}
            </span>
                <span style={{ color: "green", fontSize: "14px", marginLeft:"2px" }}>
                  {props.vegan !== false && "Vegan "}
                  {props.vegetarian !== false && "Vegetarian "}
                  {props.glutenFree !== false && "Gluten Free"}
                </span>
            <span className="address">{props.address}, {props.district}</span>
            <span className="price-range">Price Range: {props.price}</span>
            <span className="takeaway">
              <i className={props.takeaway === true ?
                "fas fa-check" : "fas fa-times"}>
              </i>
              Takeaway
            </span>
            <span className="delivery">
              <i className={props.delivery != null ?
                "fas fa-check" : "fas fa-times"}></i>Delivery</span>
            <div className={clickOnView ? "view-less" : "view-more"}>
              <span style={{ marginTop: "0.2rem" }}>
                URL:<a href={props.url} target="_blank" rel="noreferrer">{props.url}</a>
              </span>
              <span
                style={{ display: "flex", flexDirection: "row", marginTop: "0.2rem" }}>
                Opening hours:
                <span>
                  {props.openingHours === null ? "Opening Hours are Not Available" :
                   days.length !== 0 && days.map(day =>
                            <div style={{ display: "flex", flexDirection: "row" }}>
                                <p className="day">{day}:</p>
                                <p style={{ textAlign: "start" }}>
                                    {hours[days.indexOf(day)]}
                                </p>
                            </div>)
                }
                </span>
              </span>
            </div>
            <div className="more-options">
              <div className="option" onClick={handleClickOnPhone}>
                Phone
                    <i className="fas fa-phone"></i>
                {clickOnPhone && <PhoneModal name={props.name} phone={props.phone}/>}
              </div>
              <div className="option" onClick={handleClickOnMenu}>
                Weekly Menu
                    <i className="fab fa-elementor"></i>
                {clickOnMenu && <MenuModal name={props.name} menu={props.menu} date={props.menuDates}/>}
              </div>
              <div className="option" onClick={handleClickOnView}>
                {clickOnView ? "View Less" : "View More"}
                <i className={clickOnView ? "fas fa-angle-up" : "fas fa-angle-down"}></i>
              </div>
            </div>
          </div>
        </div>
      </li>
    </>
  );
})