import React, { useEffect } from 'react';
import { Button } from '../button/Button';
import MobileNavbar from './MobileNavbar'
import { Link } from 'react-router-dom';
import './Navbar.css';
import { Modal } from '../forms/Modal'
import Searchbox from '../search/Searchbox';
import NavbarLogic from './NavbarLogic';

function Navbar() {
  const { click, button, showButton, handleClick, closeMobileMenu }
    = MobileNavbar();
  const { navMenuClassName, searchbox, showLogInModal,
          showSignUpModal, openLogInModal, openSignUpModal,
          setShowLogInModal, setShowSignUpModal }
    = NavbarLogic();

  useEffect(() => {
   showButton();
  }, [showButton]);

  window.addEventListener('resize', showButton);

  return (
    <>
      <nav className={click ? 'navbar active' : 'navbar'}>
        <div className='navbar-container'>
          <Link to='/' className='navbar-logo' onClick={closeMobileMenu}>
           Restaurateur<i class="fas fa-utensils" />
          </Link>
          <div className={click ? 'hidden' : searchbox}>
            <Searchbox />
          </div>
          <div className='menu-icon' onClick={handleClick}>
            <i className={click ? 'fas fa-times' : 'fas fa-bars'} />
          </div>
          <ul className={click ? 'nav-menu active' : navMenuClassName}>
            <li className='nav-item'>
              <Link to='/' className='nav-links' onClick={closeMobileMenu}>
                Home
              </Link>
            </li>
            <li className='nav-item'>
              <Link
                to='/restaurants'
                className='nav-links'
                onClick={closeMobileMenu}
              >
                Restaurants
              </Link>
            </li>
            <li>
              <Link
                className='nav-links-mobile'
                onClick={openLogInModal}
              >
                Log In
              </Link>

              <Link
                className='nav-links-mobile'
                onClick={openSignUpModal}
              >
                Sign Up
              </Link>
            </li>
          </ul>
          {button &&
            <Button
              buttonStyle='btn--outline'
              buttonSize='btn--medium'
              onClick={openLogInModal}
              id="login">
                LOG IN
            </Button>}
          <Modal
            showLogInModal={showLogInModal}
            setShowLogInModal={setShowLogInModal}
          />
          {button &&
            <Button
              id="signup"
              buttonSize='btn--medium'
              onClick={openSignUpModal}>
                SIGN UP
            </Button>}
          <Modal
            showSignUpModal={showSignUpModal}
            setShowSignUpModal={setShowSignUpModal}
          />
        </div>
      </nav>
    </>
  );
}

export default Navbar;