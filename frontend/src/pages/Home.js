import React from 'react';
import Cards from '../components/cards/Cards';
import HeroSection from '../components/hero/HeroSection';
import Districts from '../components/districts/Districts';
import Navbar from '../components/navbar/Navbar';
import CollegeSection from '../components/college/CollegeSection';




function Home() {
  return (
    <>
      <Navbar />
      <HeroSection />
      <Cards />
      <CollegeSection />
      <Districts />
    </>
  );
}

export default Home;
