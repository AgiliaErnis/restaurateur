import React from 'react';
import DistrictItem from './District_item';
import './Districts.css';
import { PopularDistrictsData }  from './PopularDistrictsData'

function Districts() {
  return (
    <div className='districts'>
      <h1>Popular districts in <b>Prague</b></h1>
      <div className="districts__container">
        <div className="districts__wrapper">
          <ul className="districts__items">
            {PopularDistrictsData.map(district =>
            { return PopularDistrictsData.indexOf(district) < 3 &&
                   <DistrictItem
                    district= {district.district}
                    num_of_restaurants={district.num_of_restaurants}
                  />
            })}
          </ul>
          <ul className="districts__items">
            {PopularDistrictsData.map(district =>
            { return PopularDistrictsData.indexOf(district) >= 3 &&
                   <DistrictItem
                    district= {district.district}
                    num_of_restaurants={district.num_of_restaurants}
                  />
            })}
          </ul>
          </div>
        </div>
    </div>
  );
}

export default Districts;
