import React, { useState } from 'react';
import './App.css';
import Home from '../pages/Home';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Restaurants from '../pages/Restaurants';
import ScrollToTop from '../components/ScrollToTop';
import Footer from '../components/footer/Footer';
import { UserContext } from '../UserContext';

function App() {
  const [pragueCollegePath, setPragueCollegePath] = useState(false);
  const [clickedDistrict, setClickedDistrict] = useState(false);
  const [clickedSuggestion, setClickedSuggestion] = useState(false);
  const [checkedDistance, setCheckedDistance] = useState("1000");
  const [chosenRestaurant, setChosenRestaurant] = useState(false);
  const [generalSearchPath, setGeneralSearchPath] = useState(false);
  const [clickOnScreen, setClickOnScreen] = useState(false)

  return (
    <>
      <Router>
        <ScrollToTop />
        <Switch>
          <UserContext.Provider value={{
            pragueCollegePath, setPragueCollegePath,
            clickedDistrict, setClickedDistrict,
            clickedSuggestion, setClickedSuggestion,
            checkedDistance, setCheckedDistance,
            chosenRestaurant, setChosenRestaurant,
            generalSearchPath, setGeneralSearchPath,
            clickOnScreen, setClickOnScreen,
          }}>
            <Route path='/' exact component={Home} />
            <Route path='/restaurants' component={Restaurants} />
          </UserContext.Provider>
        </Switch>
        <Footer />
      </Router>
    </>
  );
}

export default App;