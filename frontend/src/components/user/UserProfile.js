import React, { useContext } from 'react'
import "./UserProfile.css"
import { UserContext } from '../../UserContext';
import RestaurantItem from '../restaurants/RestaurantItem';
import { ImagePlaceHolder } from
    '../../components/restaurants/PhotoSlider/ImagePlaceHolder';
import { UserMenuItemsData }   from './UserMenuItemsData'

function UserProfile() {
    const { clickedUserMenuItem, setClickedUserMenuItem } = useContext(UserContext)

    return (
        <div className="account-container">
        <div className="user-profile">
            <div className="profile-cover-container">
                <img src='/images/Home/hero.jpg' alt='hero-background'className="user-background"/>
                <div className="user-info">
                    <div className="user">
                        <i class="far fa-user-circle"></i>
                        <p>Username</p>
                    </div>
                    <div className="saved-restaurants">
                        <span className="number">4</span>
                         <span>Saved Restaurants</span>
                    </div>
                </div>
            </div>
        </div>
            <div className="user-info-box-container">
                <div className="user-info-box">
                    <div className="info-box-menu">
                        {UserMenuItemsData.map(item =>
                         <div className="info-box-menu-item">
                                <h4 className="menu-item-header">{item.menuItem}</h4>
                                {item.options.map(option =>
                                   <div className="menu-item-options ">
                                        <span className={clickedUserMenuItem === option.keyword ?
                                            "menu-item-option clicked"
                                            :
                                            "menu-item-option"
                                            }
                                            onClick={() => setClickedUserMenuItem(option.keyword)}
                                        >
                                            {option.option}
                                        </span>
                                    </div>
                                )}
                         </div>
                        )}
                    </div>
                </div>
                <div className="user-info-content">
                    <h4 className="menu-item-header content">Saved Restaurants</h4>
                    <div className="saved-restaurants-container">
                        <RestaurantItem
                            photos={ImagePlaceHolder}
                            name="Terasa U Zlate studne"
                            rating="3.7"
                            tags="Czech, International"
                            address="U Zlaté studně 166/4, Praha 1"
                            price="Price Range: Over 600 CZK"
                            takeaway= "true"
                            delivery = "false"
                        />
                        <RestaurantItem
                            photos={ImagePlaceHolder}
                            name="Gallery 44 Restaurant"
                            rating="2.7"
                            tags="International"
                            address="Italska 166/4, Praha 1"
                            price="Price Range: 300 - 600 CZK"
                            takeaway= "false"
                            delivery = "true"
                        />
                        <RestaurantItem
                            photos={ImagePlaceHolder}
                            name="Terasa U Zlate studne"
                            rating="3.7"
                            tags="Czech, International"
                            address="U Zlaté studně 166/4, Praha 1"
                            price="Price Range: Over 600 CZK"
                            takeaway= "true"
                            delivery = "false"
                        />
                        <RestaurantItem
                            photos={ImagePlaceHolder}
                            name="Terasa U Zlate studne"
                            rating="2.7"
                            tags="International"
                            address="Italska 166/4, Praha 1"
                            price="Price Range: 300 - 600 CZK"
                            takeaway= "false"
                            delivery = "true"
                        />
                    </div>
                </div>
            </div>
        </div>
    )
}

export default UserProfile
