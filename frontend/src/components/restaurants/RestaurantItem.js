import React from 'react';
import './RestaurantItem.css'
import Rating from '@material-ui/lab/Rating';
import { PhotoSlider } from './PhotoSlider/PhotoSlider';
import MobileNavbar from '../navbar/MobileNavbar'

function RestaurantItem(props) {
  const {click, handleClick }
    = MobileNavbar();
  return (
    <>
      <li className='restaurant_card_item'>
          <div className="content">
              <PhotoSlider slides={props.photos} />
              <div className="restaurant_description">
                <span className="restaurant-name">{props.name}
                  <div className="save-container">
                    <i onClick={handleClick}
                      className={click ? "save-btn-active"
                                  : "save-btn"}>
                    </i>
                <div className={click ? 'save-on-hover-hidden'
                                 : 'save-on-hover'}>
                      Click to Save
                    </div>
                  <div className={click ? "saved" : "saved-hidden"}>Saved</div></div>
                </span>
                  <div className="rating-container">
                  <Rating name="read-only" value={props.rating}
                          precision={0.1} readOnly
                  />
                  <span className="rating-num">({props.rating})</span>
                  </div>
                  <span className="tags">{props.tags}</span>
                  <span className="address">{props.address}</span>
                  <span className="price-range">{props.price}</span>
                  <span className="takeaway">
                  <i className={props.takeaway === "true" ?
                    "fas fa-check" : "fas fa-times"}></i>Takeaway</span>
                  <span className="delivery">
                    <i className={props.delivery === "true" ?
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
}

export default RestaurantItem;