import React from 'react';
import { VerticalFilter } from '../components/filtration/VerticalFilter';
import RestaurantItem from '../components/restaurants/RestaurantItem';
import Select from 'react-select'
import Navbar from '../components/navbar/Navbar';
import SelectStyle from '../components/search/SelectStyle';
import SelectLogic from '../components/search/SelectLogic';
import './Restaurants.css'
import { RestaurantPhotos } from
  '../components/restaurants/PhotoSlider/RestaurantPhotos';
import RestaurantPagination from '../components/restaurants/pagination/Pagination';


export default function Restaurants() {
  const { customThemes, customStyles } = SelectStyle();
  const { sortOptions, setSortResultHandler } = SelectLogic();

  return (
    <>
      <Navbar/>
      <div className="restaurants-hero-container">
        <VerticalFilter />
        <div className="restaurant-cards-container">
          <div className="restaurant-cards-header">
            <h1>Restaurants in Prague</h1>
            <Select
              defaultValue="Sort by"
              options={sortOptions}
              styles={customStyles}
              theme={customThemes}
              onChange={setSortResultHandler}
              className="sort"
              placeholder="Sort by"
              isSearchable
            />
          </div>
          <RestaurantItem
            photos={RestaurantPhotos}
            name="Terasa U Zlate studne"
            rating="3.7"
            tags="Czech, International"
            address="U Zlaté studně 166/4, 118 00 Malá Strana, Czechia"
            price="Price Range: Over 600 CZK"
            takeaway= "true"
            delivery = "false"
          />
          <RestaurantItem
            photos={RestaurantPhotos}
            name="Gallery 44 Restaurant"
            rating="2.7"
            tags="International, Vegan, Vegetarian"
            address="Italska 166/4, 118 00 Malá Strana, Czechia"
            price="Price Range: 300 - 600 CZK"
            takeaway= "false"
            delivery = "true"
          />
           <RestaurantItem
            photos={RestaurantPhotos}
            name="Terasa U Zlate studne"
            rating="3.7"
            tags="Czech, International"
            address="U Zlaté studně 166/4, 118 00 Malá Strana, Czechia"
            price="Price Range: Over 600 CZK"
            takeaway= "true"
            delivery = "false"
          />
          <RestaurantItem
            photos={RestaurantPhotos}
            name="Terasa U Zlate studne"
            rating="2.7"
            tags="International, Vegan, Vegetarian"
            address="Italska 166/4, 118 00 Malá Strana, Czechia"
            price="Price Range: 300 - 600 CZK"
            takeaway= "false"
            delivery = "true"
          />
          <RestaurantPagination />
        </div>
      </div>

    </>
  );
}