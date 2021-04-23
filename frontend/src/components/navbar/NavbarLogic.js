import { useState } from 'react'

const NavbarLogic = () => {
    let navMenuClassName = setNavMenuClassName();
    let searchbox = addSearchboxToNavbar();

    const [showSignUpModal, setShowSignUpModal] = useState(false);
    const [showLogInModal, setShowLogInModal] = useState(false);

    const openSignUpModal = () => {
        setShowSignUpModal(prev=> !prev);
    }

    const openLogInModal = () => {
        setShowLogInModal(!showLogInModal);
    }

    function addSearchboxToNavbar () {
        switch(window.location.pathname){
            case '/restaurants':
                return "nav-links hover mobile";
            default:
                return "hidden";
        }
    }

    function setNavMenuClassName () {
        switch(window.location.pathname){
            case '/restaurants':
                return "nav-menu-updated";
            default:
                return "nav-menu";
        }
    }

    return {
        navMenuClassName, searchbox, showLogInModal,
        showSignUpModal, openLogInModal, openSignUpModal,
        setShowLogInModal, setShowSignUpModal
    }
}

export default NavbarLogic;