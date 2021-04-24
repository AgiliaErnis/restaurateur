import React from 'react';
import './App.css';
import Home from '../pages/Home';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Restaurants from '../pages/Restaurants';
import ScrollToTop from '../components/ScrollToTop';
import Footer from '../components/footer/Footer';

function App() {
  return (
    <>
      <Router>
        <ScrollToTop />
        <Switch>
          <Route path='/' exact component={Home} />
          <Route path='/restaurants' component={Restaurants} />
        </Switch>
        <Footer />
      </Router>
    </>
  );
}

export default App;