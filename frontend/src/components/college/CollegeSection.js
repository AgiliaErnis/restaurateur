import React, { useContext, useEffect, useState } from 'react';
import './CollegeSection.css';
import { Link } from 'react-router-dom';
import { Button } from '../button/Button';
import Aos from 'aos';
import 'aos/dist/aos.css';
import { UserContext } from '../../UserContext';

function CollegeSection() {
  const { setPragueCollegePath } = useContext(UserContext)
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

 useEffect(() => {
  Aos.init({ duration: 1000 })
  }, [])

  window.addEventListener('resize', showImage);
  return (
    <>
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
                <Link to='/restaurants' onClick={() => setPragueCollegePath(true)}>
                  <Button buttonSize='btn--large'
                          buttonStyle='btn--search'
                          className="btn-see-more"
                  >
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
    </>
  );
}

export default CollegeSection;