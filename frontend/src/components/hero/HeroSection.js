import React from 'react';
import './HeroSection.css';
import Searchbox from './search/Searchbox'

function HeroSection() {
  return (
    <div className='hero-container'>
      <img src='/images/Home/hero.jpg'/>
      <h1>Discover the best dining destinations in Prague</h1>
        <Searchbox />
    </div>
  );
}

export default HeroSection;
