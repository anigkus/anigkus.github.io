import React from 'react';

function Hero({ name }) {
  if (name === 'Diana') {
    throw new Error('Hero Error');
  }
  return <h1>{name}</h1>;
}

export default Hero;
