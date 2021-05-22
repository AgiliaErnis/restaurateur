import React, { useState, useEffect } from 'react';
import { Redirect } from 'react-router-dom';
import './App.css';
import Home from '../pages/Home';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Restaurants from '../pages/Restaurants';
import ScrollToTop from '../components/ScrollToTop';
import UserAccount from '../pages/UserAccount'
import Footer from '../components/footer/Footer';
import { UserContext } from '../UserContext';

function App() {
  const [pragueCollegePath, setPragueCollegePath] = useState(false);
  const [clickedDistrict, setClickedDistrict] = useState(false);
  const [clickedSuggestion, setClickedSuggestion] = useState(false);
  const [checkedDistance, setCheckedDistance] = useState("1000")
  const [restaurants, setRestaurants] = useState([]);
  const [clickedUserMenuItem, setClickedUserMenuItem]
    = useState("saved")
  const [goodPassword, setGoodPassword] = useState(false)
  const [chosenRestaurant, setChosenRestaurant] = useState(false);
  const [generalSearchPath, setGeneralSearchPath] = useState(false);
  const [incorrectPassword, setIncorrectPassword] = useState(false);
  const [successfullLogin, setSuccessfullLogin] = useState(false)
  const [username, setUsername] = useState(false)
  const [incorrectOldPassword, setIncorrectOldPassword] = useState(false)
  const [logout, setLogout] = useState(false);
  const [newUsername, setNewUsername] = useState(false)
  const [incorrectPasswordOnDelete, setIncorrectPasswordOnDelete] = useState(false)
  const [deleteAccount, setDeleteAccount] = useState(false)
  const [savedRestaurants, setSavedRestaurants] = useState([])

  useEffect(() => {
    const userLoggedIn = localStorage.getItem("user-logged-in");
    if (userLoggedIn) {
      setSuccessfullLogin(userLoggedIn);
    }
  }, [])

  useEffect(() => {
    const UserInfo = {
      method: 'GET',
      credentials: 'include',
      headers: {
        'Content-Type': 'application/json'
      }
    }

    async function getUserInfo() {
      if (!incorrectPassword) {
        await fetch('http://localhost:8080/auth/user', UserInfo)
          .then(response => response.json())
          .then(res => {
            if (res.Status === 200) {
              setUsername(res.User.Name);
              setSavedRestaurants(res.User.SavedRestaurants)
            }
          })
      }}
    getUserInfo();

  }, [incorrectPassword, successfullLogin,
    newUsername, setSuccessfullLogin,setSavedRestaurants])

  useEffect(() => {
       localStorage.setItem("user-logged-in", successfullLogin)
  }, [successfullLogin])

  useEffect(() => {
    if (logout)
      <Redirect to="/" />
  }, [logout, successfullLogin])

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
            restaurants, setRestaurants,
            clickedUserMenuItem, setClickedUserMenuItem,
            goodPassword, setGoodPassword,
            chosenRestaurant, setChosenRestaurant,
            generalSearchPath, setGeneralSearchPath,
            incorrectPassword, setIncorrectPassword,
            successfullLogin, setSuccessfullLogin,
            username, setUsername,
            incorrectOldPassword, setIncorrectOldPassword,
            logout, setLogout,
            newUsername, setNewUsername,
            incorrectPasswordOnDelete, setIncorrectPasswordOnDelete,
            deleteAccount, setDeleteAccount,
            savedRestaurants,setSavedRestaurants
          }}>
            <Route path='/' exact component={Home} />
            <Route path='/restaurants' component={Restaurants} />
            { (successfullLogin && !logout && !deleteAccount) ?
              <Route path='/user' component={UserAccount} />
              :
              <>
               <Route path='/user' component={UserAccount} />
              <Redirect to='/' /></>
          }
          </UserContext.Provider>
        </Switch>
        <Footer />
      </Router>
    </>
  );
}

export default App;