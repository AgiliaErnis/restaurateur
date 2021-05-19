import React from 'react'
import { ImagePlaceHolder } from '../restaurants/PhotoSlider/ImagePlaceHolder'
import "./Searchbox.css"

function SuggestedRestaurantItem(props) {
    return (
        <li className={
            `suggested-restaurant ${props.isHighlighted ? `active` : ""}`}
            onClick={props.onSelectItem}>
            <div className="restaurant-image">
                <img src={props.image !== "" ?
                    props.image
                    :
                    ImagePlaceHolder}
                    alt="suggested-restaurant" />
            </div>
            <div className="text">
                <p className="restaurant-name">
                    {props.name}
                </p>
                <p className="restaurant-address">
                  {props.address}, {props.district}
                </p>
            </div>
            </li>
    )
}

export default SuggestedRestaurantItem
