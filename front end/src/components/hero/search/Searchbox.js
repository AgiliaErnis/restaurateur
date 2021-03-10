import React, { useState, useEffect } from 'react';
import { Button } from '../../button/Button';
/*import { Link } from 'react-router-dom';*/
import './Searchbox.css';

function Searchbox () {
    return (
        <div class="form-box">
            <input class="search-field cafes" name="cafe" type="text" placeholder="Restaurants, Cafes..."></input>
            <input class="search-field location" name="location" type="text" placeholder="Location"></input>
            <Button
            className='btns search'
            buttonSize='btn--large'
            >
            Search
            </Button>
            </div>
        )
}



export default Searchbox;
