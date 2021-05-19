import React, { useEffect, useContext } from 'react';
import { Button } from '../button/Button';
import MobileNavbar from './MobileNavbar'
import { Link } from 'react-router-dom';
import './Navbar.css';
import { Modal } from '../forms/Modal'
import Searchbox from '../search/Searchbox';
import NavbarLogic from './NavbarLogic';
import { UserContext } from '../../UserContext';
import UserNavbar from '../user/UserNavbar';

function Navbar() {
  const { click, button, showButton,
          handleClick, closeMenuDiscardChanges, closeMenuOpenPCRestaurants }
    = MobileNavbar();
  const { navMenuClassName, searchbox, showLogInModal,
          showSignUpModal, openLogInModal, openSignUpModal,
          setShowLogInModal, setShowSignUpModal }
    = NavbarLogic();
  const { pragueCollegePath,
    successfullLogin, setSuccessfullLogin, logout }
    = useContext(UserContext);

   function setRestaurantsNavLink () {
    switch(window.location.pathname){
      case '/restaurants':
         return (pragueCollegePath === true ?
           "All Restaurants"
           :
           "PC Restaurants")
      default:
        return "All Restaurants";
    }
  }

  useEffect(() => {
   showButton();
  }, [showButton]);

  useEffect(() => {
    if (logout) {
      setShowLogInModal(false);
      setSuccessfullLogin(false)
    }
  }, [logout,setSuccessfullLogin,setShowLogInModal])

  useContext(() => {
    if (localStorage.getItem("user-logged-in") === "true") {
      setShowSignUpModal(false); setShowLogInModal(false)
    }
  },[localStorage.getItem("user-logged-in")])

  window.addEventListener('resize', showButton);

  useEffect(() => {
    setSuccessfullLogin(successfullLogin)
  },[successfullLogin,setSuccessfullLogin])

  return (
    <>
      <nav className={click ? 'navbar active' : 'navbar'}>
        <div className='navbar-container'>
          <Link to='/' className='navbar-logo' onClick={closeMenuDiscardChanges}>
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
              <Link to='/' className='nav-links' onClick={closeMenuDiscardChanges}>
                Home
              </Link>
            </li>
            <li className='nav-item'>
              <Link
                to='/restaurants'
                className='nav-links'
                onClick={setRestaurantsNavLink() === "All Restaurants" ?
                  closeMenuDiscardChanges
                  :
                  closeMenuOpenPCRestaurants}
              >
                {setRestaurantsNavLink()}
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
          {successfullLogin  === true  ?
            <UserNavbar />
            :
            <>
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
          </>}
        </div>
      </nav>
    </>
  );
}

export default Navbar;