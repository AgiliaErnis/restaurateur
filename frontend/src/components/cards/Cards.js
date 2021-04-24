import React from 'react';
import './Cards.css';
import CardItem from './CardItem';

function Cards() {
  return (
    <div className='cards'>
      <h1>Check Out the Top Suggestions! </h1>
      <div className='cards__container'>
        <div className='cards__wrapper'>
          <ul className='cards__items'>
            <CardItem
              src='images/Home/international.webp'
              text="View the city's more than 1000 International restaurants"
              label='International'
            />
            <CardItem
              src='images/Home/italian.jpg'
              text="Explore all the Italian restaurants of Prague"
              label='Italian'
            />
          </ul>
          <ul className='cards__items'>
            <CardItem
              src='images/Home/czech.jpg'
              text='Find your most suitable Czech restaurant'
              label='Czech'
            />
            <CardItem
              src='images/Home/glutenfree.webp'
              text='Check out  the restaurants with gluten free options'
              label='Gluten Free'
            />
            <CardItem
              src='images/Home/vegan.jpeg'
              text='Check out the restaurants with vegan and vegetarian meals'
              label='Vegan/Vegetarian'
            />
          </ul>
        </div>
      </div>
    </div>
  );
}

export default Cards;