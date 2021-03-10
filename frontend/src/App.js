import React from 'react';
import Navbar from './components/navbar/Navbar';
import './App.css';
import Home from './pages/Home';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Restaurants from './pages/Restaurants';
import Services from './pages/Services';

function App() {
  return (
    <>
      <Router>
        <Navbar />
        <Switch>
          <Route path='/' exact component={Home} />
          <Route path='/services' component={Services} />
          <Route path='/restaurants' component={Restaurants} />
        </Switch>
      </Router>
    </>
  );
}

export default App;
