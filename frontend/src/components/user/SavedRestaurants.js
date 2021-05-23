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
                phone={restaurant.PhoneNumber}
                ID={restaurant.ID}
                photos={restaurant.Images.length !== 0 ?
                  restaurant.Images : ImagePlaceHolder}
                name={restaurant.Name}
                rating={restaurant.Rating === "" ?
                  "Rating is not available" : restaurant.Rating}
                tags={restaurant.Cuisines !== null ?
                  restaurant.Cuisines.map((cuisine) => {
                    if (restaurant.Cuisines.indexOf(cuisine) ===
                      restaurant.Cuisines.length - 1) {
                      return cuisine
                    } else {
                        return cuisine + ","
                      }
                    })
                    :
                    "Cuisines are not available"}
                address={restaurant.Address}
                district={restaurant.District}
                price={restaurant.PriceRange}
                takeaway={restaurant.Takeaway}
                    delivery={restaurant.DeliveryOptions}
                    click={true}
              />
            })}
        </div>
    )
}

export default SavedRestaurants
