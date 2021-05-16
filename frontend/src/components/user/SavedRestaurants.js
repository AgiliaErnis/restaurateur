import React from 'react'
import RestaurantItem from '../restaurants/RestaurantItem';
import { ImagePlaceHolder } from
    '../../components/restaurants/PhotoSlider/ImagePlaceHolder';

function SavedRestaurants() {
    return (
        <div className="saved-restaurants-container">
            <RestaurantItem
                photos={ImagePlaceHolder}
                name="Terasa U Zlate studne"
                rating="3.7"
                tags="Czech, International"
                address="U Zlaté studně 166/4, Praha 1"
                price="Price Range: Over 600 CZK"
                takeaway="true"
                delivery="false"
            />
            <RestaurantItem
                photos={ImagePlaceHolder}
                name="Gallery 44 Restaurant"
                rating="2.7"
                tags="International"
                address="Italska 166/4, Praha 1"
                price="Price Range: 300 - 600 CZK"
                takeaway="false"
                delivery="true"
            />
            <RestaurantItem
                photos={ImagePlaceHolder}
                name="Terasa U Zlate studne"
                rating="3.7"
                tags="Czech, International"
                address="U Zlaté studně 166/4, Praha 1"
                price="Price Range: Over 600 CZK"
                takeaway="true"
                delivery="false"
            />
            <RestaurantItem
                photos={ImagePlaceHolder}
                name="Terasa U Zlate studne"
                rating="2.7"
                tags="International"
                address="Italska 166/4, Praha 1"
                price="Price Range: 300 - 600 CZK"
                takeaway="false"
                delivery="true"
            />
        </div>
    )
}

export default SavedRestaurants
