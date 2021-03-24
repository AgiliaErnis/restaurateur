import React from 'react';
import { Link, BrowserRouter as Router } from 'react-router-dom';
import './Districts.css'

function DistrictItem(props) {
  return (
    <>
      <Router>
        <li className='districts__item'>
          <Link className='districts__item__link' to={props.path}>
            <h4>
              {props.district} ({props.num_of_restaurants} places)
            </h4>
            <i class="fas fa-angle-right" />
          </Link>
        </li>
      </Router>
    </>
  );
}

export default DistrictItem;