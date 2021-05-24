import React, { useContext } from 'react';
import { Link } from 'react-router-dom';
import { UserContext } from '../../UserContext';
import './Districts.css';

function DistrictItem(props) {
  const { setClickedDistrict }  = useContext(UserContext)

  return (
    <>
      <li
        className='districts__item'
        onClick={() => setClickedDistrict(props.district)}
      >
        <Link className='districts__item__link' to='/restaurants'>
          <h4>
            {props.district} ({props.num_of_restaurants} places)
          </h4>
          <i class="fas fa-angle-right" />
        </Link>
      </li>
    </>
  );
}

export default DistrictItem;