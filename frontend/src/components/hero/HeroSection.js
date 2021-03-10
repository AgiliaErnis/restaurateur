import React from 'react';
import './HeroSection.css';
import Searchbox from './search/Searchbox'

function HeroSection() {
  return (
    <div className='hero-container'>
      <video src='/videos/test.mp4' autoPlay loop muted />
      <h1>Discover the best dining destinations in Prague</h1>
        <Searchbox />
    </div>
  );
}

export default HeroSection;
