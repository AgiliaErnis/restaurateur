import { useState } from 'react';

const MobileNavbar = () => {
  const [click, setClick] = useState(false);
  const [button, setButton] = useState(true);

  const handleClick = () => setClick(!click);
  const closeMobileMenu = () => setClick(false);

  const showButton = () => {
    if (window.innerWidth <= 1120) {
      setButton(false);
    } else {
      setButton(true);
    }
  };

  const showSearch = () => {
    if (window.innerWidth <= 820) {
      setButton(false);
    } else {
      setButton(true);
    }
  };

  return{click, button, showButton, handleClick, closeMobileMenu, showSearch}

}

export default MobileNavbar;
