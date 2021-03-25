/*
import React from 'react';
import { Link, BrowserRouter as Router } from 'react-router-dom';
import './Cards.css'

function CardItem(props) {
  return (
    <>
      <Router>
        <li className='cards__item'>
          <Link className='cards__item__link'>
            <figure className='cards__item__pic-wrap' data-category={props.label}>
              <img
                className='cards__item__img'
                alt='Food Category Image'
                src={props.src}
              />
            </figure>
            <div className='cards__item__info'>
              <h5 className='cards__item__text'>{props.text}</h5>
            </div>
          </Link>
        </li>
      </Router>
    </>
  );
}

export default CardItem;
*/


import React from 'react';
import { Link, BrowserRouter as Router } from 'react-router-dom';
import './Cards.css'

function CardItem(props) {
  return (
    <>
      <Router>
        <li className='cards__item'>
          <Link className='cards__item__link' path={props.path}>
            <figure className='cards__item__pic-wrap' label={props.label}>
              <img
                className='cards__item__img'
                alt='Food Category Image'
                src={props.src}
              />
            </figure>
            <div className='cards__item__info'>
              <h5 className='cards__item__text'>{props.text}</h5>
            </div>
          </Link>
        </li>
      </Router>
    </>
  );
}

export default CardItem;
