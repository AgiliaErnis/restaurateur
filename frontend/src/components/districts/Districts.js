import React from 'react';
import DistrictItem from './District_item';
import './Districts.css';

function Districts() {
  return (
    <div className='districts'>
      <h1>Popular districts in <b>Prague</b></h1>
      <div className="districts__container">
        <div className="districts__wrapper">
          <ul className="districts__items">
            <DistrictItem
              district="Prague 1"
              num_of_restaurants="50"
            />
            <DistrictItem
              district="Prague 2"
              num_of_restaurants="50"
            />
            <DistrictItem
              district="Prague 3"
              num_of_restaurants="50"
            />
        </ul>
        <ul className="districts__items">
            <DistrictItem
              district="Prague 4"
              num_of_restaurants="50"
            />
            <DistrictItem
              district="Prague 5"
              num_of_restaurants="50"
            />
            <DistrictItem
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
