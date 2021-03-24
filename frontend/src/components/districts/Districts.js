import React from 'react';
import '../cards/Cards.css';
import DistrictItem from './District_item';
import './Districts.css'


function Districts() {
  return (
    <div className='districts'>
      <h1>Popular districts in <b>Prague</b></h1>
      <div className="districts__container">
        <div className="districts__wrapper">
          <ul className="districts__items">
            <DistrictItem
              path='/restaurants'
              district="Prague 1"
              num_of_restaurants="50"
            />
            <DistrictItem
              path='/restaurants'
              district="Prague 2"
              num_of_restaurants="50"
            />
            <DistrictItem
              path='/restaurants'
              district="Prague 3"
              num_of_restaurants="50"
            />
        </ul>
        <ul className="districts__items">
            <DistrictItem
              path='/restaurants'
              district="Prague 4"
              num_of_restaurants="50"
            />
            <DistrictItem
              path='/restaurants'
              district="Prague 5"
              num_of_restaurants="50"
            />
            <DistrictItem
              path='/restaurants'
              district="Prague 6"
              num_of_restaurants="50"
            />
          </ul>
          </div>
        </div>
    </div>
  );
}

export default Districts;
