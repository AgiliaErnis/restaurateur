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
      <div className='footer-links'>
        <div className='footer-link-wrapper'>
          <div className='footer-link-items'>
            <h2>Home</h2>
            <Link to='/'>Top Suggestions</Link>
            <Link to='/'>Popular Localities</Link>
          </div>
          <div className='footer-link-items'>
            <h2>Restaurants</h2>
            <Link to='/'>Filtration</Link>
            <Link to='/'>Destinations</Link>
          </div>
          <div className='footer-link-items'>
            <h2>For You</h2>
            <Link to='/'>Privacy</Link>
            <Link to='/'>Terms</Link>
            <Link to='/'>Security</Link>
          </div>
        </div>
      </div>
      <section className='icons'>
        <div className='wrapper'>
          <div className='footer-logo'>
            <Link to='/' className='logo'>
            Restaurateur<i className="fas fa-utensils" />
            </Link>
          </div>
          <small className='website-rights'> © 2021</small>
        </div>
      </section>
    </div>
  );
}

export default Footer;