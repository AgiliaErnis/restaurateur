import React, { useContext } from 'react'
import { RestaurantItem } from '../restaurants/RestaurantItem';
import { ImagePlaceHolder } from
    '../../components/restaurants/PhotoSlider/ImagePlaceHolder';
import { UserContext } from '../../UserContext';

function SavedRestaurants() {
  const savedRestaurants = useContext(UserContext)

    return (
        <div className="saved-restaurants-container">
            {savedRestaurants.savedRestaurants !== null &&
                savedRestaurants.savedRestaurants.map(restaurant => {
                  return <RestaurantItem
                phone={restaurant.phoneNumber}
                ID={restaurant.id}
                photos={restaurant.images.length !== 0 ?
                  restaurant.images : ImagePlaceHolder}
                name={restaurant.name}
                rating={restaurant.rating === "" ?
                  "Rating is not available" : restaurant.rating}
                tags={restaurant.cuisines !== null ?
                  restaurant.cuisines.map((cuisine) => {
                    if (restaurant.cuisines.indexOf(cuisine) ===
                      restaurant.cuisines.length - 1) {
                      return cuisine
                    } else {
                        return cuisine + ","
                      }
                    })
                    :
                    "Cuisines are not available"}
                address={restaurant.address}
                district={restaurant.district}
                price={restaurant.priceRange}
                takeaway={restaurant.takeaway}
                delivery={restaurant.deliveryOptions}
                click={true}
                menu={restaurant.weeklyMenu !== "null" ? restaurant.weeklyMenu : null}
                vegan={restaurant.vegan}
                vegetarian={restaurant.vegetarian}
                glutenFree={restaurant.glutenFree}
                url={restaurant.url !== "" ? restaurant.url : "URL is not available"}
                OpeningHours={restaurant.openingHours}
                  />
            })}
        </div>
    )
}

export default SavedRestaurants
