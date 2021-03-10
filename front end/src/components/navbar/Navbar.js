import React, { useState, useEffect } from 'react';
import { Button } from '../button/Button';
import { Link } from 'react-router-dom';
import './Navbar.css';
import { Modal } from '../forms/Modal'


function Navbar() {
  const [click, setClick] = useState(false);
  const [button, setButton] = useState(true);
  const [showSignUpModal, setShowSignUpModal] = useState(false);
  const [showLogInModal, setShowLogInModal] = useState(false);

  const openSignUpModal = () => {
    setShowSignUpModal(prev=> !prev);
  }

  const openLogInModal = () =>
    setShowLogInModal (!showLogInModal);


  const handleClick = () => setClick(!click);
  const closeMobileMenu = () => setClick(false);

  const showButton = () => {
    if (window.innerWidth <= 960) {
      setButton(false);
    } else {
      setButton(true);
    }
  };

  useEffect(() => {
    showButton();
  }, []);

  window.addEventListener('resize', showButton);

  return (
    <>
      <nav className={click ? 'navbar active' : 'navbar'}>
        <div className='navbar-container'>
          <Link to='/' className='navbar-logo' onClick={closeMobileMenu}>
          <i class="fas fa-utensils" />
          </Link>
          <div className='menu-icon' onClick={handleClick}>
            <i className={click ? 'fas fa-times' : 'fas fa-bars'} />
          </div>
          <ul className={click ? 'nav-menu active' : 'nav-menu'}>
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
            <li className='nav-item'>
              <Link
                to='/services'
                className='nav-links'
                onClick={closeMobileMenu}
              >
               Services
              </Link></li>
              <li>

              <Link
                to='/'
                className='nav-links-mobile'
                onClick={openLogInModal}
              >
                Log In
              </Link>

              <Link
                to='/'
                className='nav-links-mobile'
                onClick={openSignUpModal}
              >
                Sign Up
              </Link>

            </li>
          </ul>
          {button &&
          <Button buttonStyle='btn--outline' onClick={openLogInModal}
                  id ="login" >LOG IN</Button>}
          <Modal showLogInModal={showLogInModal}
                 setShowLogInModal={setShowLogInModal}/>

          {button && <Button  id ="signup" onClick={openSignUpModal}>SIGN UP</Button>}
          <Modal showSignUpModal={showSignUpModal} setShowSignUpModal={setShowSignUpModal}/>
        </div>
      </nav>
    </>
  );
}

export default Navbar;


