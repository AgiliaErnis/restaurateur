import React from 'react';
import './Footer.css';
import { Link } from 'react-router-dom';

function Footer() {
  return (
    <div className='footer-container'>
      <section className='footer-heading-container'>
        <p className='footer-heading'>
          Join us and discover the best restaurants in Prague
        </p>
      </section>
      <div class='footer-links'>
        <div className='footer-link-wrapper'>
          <div class='footer-link-items'>
            <h2>Home</h2>
            <Link to='/'>Top Suggestions</Link>
            <Link to='/'>Popular Localities</Link>
          </div>
          <div class='footer-link-items'>
            <h2>Restaurants</h2>
            <Link to='/'>Filtration</Link>
            <Link to='/'>Destinations</Link>
          </div>
          <div class='footer-link-items'>
            <h2>For You</h2>
            <Link to='/'>Privacy</Link>
            <Link to='/'>Terms</Link>
            <Link to='/'>Security</Link>
          </div>
        </div>
      </div>
      <section class='icons'>
        <div class='wrapper'>
          <div class='footer-logo'>
            <Link to='/' className='logo'>
            Restaurateur<i class="fas fa-utensils" />
            </Link>
          </div>
          <small class='website-rights'> © 2021</small>
        </div>
      </section>
    </div>
  );
}

export default Footer;