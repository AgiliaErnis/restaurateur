import React, { useEffect, useState } from 'react';
import './CollegeSection.css';
import { Button } from '../button/Button';
import { Link, BrowserRouter as Router } from 'react-router-dom';
import Aos from 'aos';
import 'aos/dist/aos.css';

function CollegeSection () {
    useEffect(() => {
    Aos.init({ duration: 1000 })
    }, [])
     const [button, setImage] = useState(true);
   const showImage = () => {
    if (window.innerWidth <= 1024) {
      setImage(false);
    } else {
      setImage(true);
    }
  };
    useEffect(() => {
   showImage();
  }, []);

  window.addEventListener('resize', showImage);
  return (
    <>
      <Router>
      <div className= 'college-section'>
        <div className='container' data-aos="fade-up">
          <div
            className='row home__hero-row'
            style={{
              display: 'flex',
              flexDirection: 'row'
            }}
          >
            <div className='col'>
              <div className='text-wrapper'>
                <div className='top-line'>
                  For Prague College Students
                  <img src='images/Home/PC.png'
                       className="prague-college-logo"
                       alt="pc logo" />
                </div>
                <h1 className= 'heading'>
                  Find the best food places near the Prague College
                </h1>
                <p className= 'description'>
                  We would like to help Prague College students to find restaurants
                  for their lunch-breaks as soon as possible. Click on Get Started
                  and find a suitable restaurant based on your own preferences.
                </p>
                <Link to='/restaurants'>
                  <Button buttonSize='btn--large'
                          buttonStyle='btn--search'
                          className="btn-see-more">
                    Get Started
                  </Button>
                </Link>
              </div>
            </div>{button &&
              <div className='col'>
                <div className='img-wrapper'>
                  <img src='images/Home/lnch.jpg'
                       alt='Student'
                       className='student-img' />
                </div>
            </div>}
          </div>
        </div>
      </div>
      </Router>
    </>
  );
}

export default CollegeSection;