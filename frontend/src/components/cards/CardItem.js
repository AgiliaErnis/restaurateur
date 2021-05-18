import React, { useContext } from 'react';
import { Link } from 'react-router-dom';
import { UserContext } from '../../UserContext';
import './Cards.css';

function CardItem(props) {
  const { setClickedSuggestion } = useContext(UserContext)

  return (
    <>
      <li
        className='cards__item'
        onClick={() => setClickedSuggestion(props.filterValue)}
      >
          <Link className='cards__item__link' to="/restaurants">
          <figure
            className='cards__item__pic-wrap'
            data-category={props.label}>
              <img
                className='cards__item__img'
                alt='Food Category'
                src={props.src}
              />
            </figure>
            <div className='cards__item__info'>
              <h5 className='cards__item__text'>{props.text}</h5>
            </div>
          </Link>
        </li>
    </>
  );
}

export default CardItem;