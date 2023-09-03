import { useState } from 'react';

function useCounter(initialCount = 0, value) {
  const [count, setCount] = useState(initialCount);

  const increament = () => {
    setCount((prevcount) => prevcount + value);
  };
  const decreament = () => {
    setCount((prevcount) => prevcount - value);
  };
  const reset = () => {
    setCount(initialCount);
  };
  return [count, increament, decreament, reset];
}

export default useCounter;
