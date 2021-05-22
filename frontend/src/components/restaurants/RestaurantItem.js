import React, { useState, useEffect } from 'react';
import './RestaurantItem.css'
import Rating from '@material-ui/lab/Rating';
import { PhotoSlider } from './PhotoSlider/PhotoSlider';
import MobileNavbar from '../navbar/MobileNavbar'

export const RestaurantItem = React.memo((props) => {
  const {click, handleClick } = MobileNavbar();
  const [savedRestaurant, setSavedRestaurant] = useState(false)

  function restaurantIsClicked () {
    switch(window.location.pathname){
      case '/user':
         return true
      default:
        return false;
    }
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

        fetch('http://localhost:8080/auth/user/saved-restaurants', saveRestaurantRequest)
          .then(response => response.json())
          .then(res => {
            if (res.Status === 200) {
              console.log(res)
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

        fetch('http://localhost:8080/auth/user/saved-restaurants', deleteRestaurantRequest)
          .then(response => response.json())
          .then(res => {
            console.log(res)
            if (res.Status === 200) {
              console.log(res)
            }

          })
      }

    }
  },[click,savedRestaurant])


  return (
    <>
      <li className='restaurant_card_item'>
        <div className="content">
          <PhotoSlider slides={props.photos} />
          <div className="restaurant_description">
            <span className="restaurant-name">
              <div className="name-container">{props.name}</div>
              <div className="save-container">
                <i onClick={restaurantIsClicked ?
                  () => {
                    setSavedRestaurant(props.ID)
                    handleClick();
                  } :
                  () => {
                    setSavedRestaurant(props.ID);
                    handleClick()
                  }
                }
                  className={restaurantIsClicked ? !click ? "save-btn" : "save-btn-active" : !click ? "save-btn" : "save-btn-active"  }

                >
                </i>
                <div className={!click ?'save-on-hover'
                  : 'save-on-hover-hidden'} >
                  Click to Save
                </div>
                <div className={!click? "saved-hidden" : "saved"}>Saved</div>
              </div>
            </span>
            <div className="rating-container">
              <Rating name="read-only"
                value={props.rating}
                precision={0.1}
                readOnly
              />
              <span className="rating-num">({props.rating})</span>
            </div>
            <span className="tags">{props.tags}</span>
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
            <div className="more-options">
              <div className="option">
                Call
                    <i className="fas fa-phone"></i>
              </div>
              <div className="option">
                Menu
                    <i className="fab fa-elementor"></i>
              </div>
              <div className="option view">
                View More
                    <i className="fas fa-angle-right" />
              </div>
            </div>
          </div>
        </div>
      </li>
    </>
  );
})