import { useState, useContext } from 'react';
import { UserContext } from "../../UserContext"

const MobileNavbar = () => {
  const [click, setClick] = useState(false);
  const [button, setButton] = useState(true);
  const { setPragueCollegePath } = useContext(UserContext)
  const { setClickedDistrict, setClickedSuggestion } = useContext(UserContext)
  const { setChosenRestaurant, setGeneralSearchPath } = useContext(UserContext);

  const handleClick = () => setClick(!click);

  const closeMenuOpenRestaurants = () => {
    setClick(false);
    setPragueCollegePath(false);
    setClickedDistrict(false);
    setClickedSuggestion(false);
    setChosenRestaurant(false);
    setGeneralSearchPath(false)
  }
  const closeMenuOpenPCRestaurants = () => {
    setClick(false);
    setPragueCollegePath(true);
    setClickedDistrict(false);
    setClickedSuggestion(false);
    setChosenRestaurant(false);
    setGeneralSearchPath(false);
  }

  const closeMenuDiscardChanges = () => {
    setClick(false);
    setPragueCollegePath(false)
    setClickedDistrict(false)
    setClickedSuggestion(false)
    setChosenRestaurant(false)
    setGeneralSearchPath(false);
   }

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

  return {
    click, button, showButton, handleClick,
    showSearch, closeMenuOpenRestaurants,
    closeMenuOpenPCRestaurants, closeMenuDiscardChanges, setClick
  }

}

export default MobileNavbar;
