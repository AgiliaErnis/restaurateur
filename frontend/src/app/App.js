import React from 'react';
import Navbar from '../components/navbar/Navbar';
import './App.css';
import Home from '../pages/Home';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Restaurants from '../pages/Restaurants';
import ScrollToTop from '../components/ScrollToTop';

function App() {
  return (
    <>
      <Router>
        <ScrollToTop />
        <Navbar />
        <Switch>
          <Route path='/' exact component={Home} />
          <Route path='/restaurants' component={Restaurants} />
        </Switch>
      </Router>
    </>
  );
}

export default App;
