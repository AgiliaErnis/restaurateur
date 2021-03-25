import React from 'react';
import Cards from '../components/cards/Cards';
import HeroSection from '../components/hero/HeroSection';
import Footer from '../components/footer/Footer';
import Districts from '../components/districts/Districts';

function Home() {
  return (
    <>
      <HeroSection />
      <Cards />
      <Districts />
      <Footer />
    </>
  );
}

export default Home;