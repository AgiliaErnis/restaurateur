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
              src='images/Home/breakfast.webp'
              text='Top breakfast places in Prague'
              label='Breakfast'
            />
            <CardItem
              src='images/Home/italian.jpg'
              text='Lorem ipsum dolor sit amet, consectetur adipiscing elit.'
              label='Italian Food'
            />
          </ul>
          <ul className='cards__items'>
            <CardItem
              src='images/Home/coffee.jpg'
              text='Lorem ipsum dolor sit amet, consectetur adipiscing elit.'
              label='Coffee'
            />
            <CardItem
              src='images/Home/dinner.webp'
              text='Lorem ipsum dolor sit amet, consectetur adipiscing elit.'
              label='Dinner Places'
            />
            <CardItem
              src='images/Home/salad.webp'
              text='Lorem ipsum dolor sit amet, consectetur adipiscing elit.'
              label='Healthy menus'
            />
          </ul>
        </div>
      </div>
    </div>
  );
}

export default Cards;