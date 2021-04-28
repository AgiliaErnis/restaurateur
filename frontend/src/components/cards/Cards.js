import React from 'react';
import './Cards.css';
import CardItem from './CardItem';
import { TopSuggestionsData } from './TopSuggestionsData'

function Cards() {

  return (
    <div className='cards'>
      <h1>Check Out the Top Suggestions! </h1>
      <div className='cards__container'>
        <div className='cards__wrapper'>
          <ul className='cards__items'>
            {TopSuggestionsData.map(suggestion => {
              return TopSuggestionsData.indexOf(suggestion) < 2 &&
              <CardItem
              src={suggestion.src}
              text={suggestion.text}
              label={suggestion.label}
              filterValue={suggestion.filterValue}
            />
            })}
          </ul>
          <ul className='cards__items'>
            {TopSuggestionsData.map(suggestion => {
              return TopSuggestionsData.indexOf(suggestion) >= 2 &&
              <CardItem
              src={suggestion.src}
              text={suggestion.text}
              label={suggestion.label}
              filterValue={suggestion.filterValue}
            />
            })}
          </ul>
        </div>
      </div>
    </div>
  );
}

export default Cards;