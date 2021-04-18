import React from 'react';
import './App.css';
import Home from '../pages/Home';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Restaurants from '../pages/Restaurants';
import Footer from '../components/footer/Footer';

function App() {
  return (
    <>
      <Router>
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