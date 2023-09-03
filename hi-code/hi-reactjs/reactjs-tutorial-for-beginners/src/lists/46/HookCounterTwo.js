import React, { useState } from 'react';

export default function HookCounterTwo() {
  const [count, setCount] = useState(0);
  const fiveIncrement = () => {
    for (let i = 1; i <= 5; i++) {
      
      setCount((prevCount) => prevCount + 1);
    }
  };
  return ( 
    <div>
      <div>{count}</div>
      <button onClick={() => setCount(0)}>Reset</button>
      <button onClick={() => setCount(count + 1)}>Increment</button>
      <button onClick={() => setCount(count - 1)}>Decrement</button>
      <button onClick={fiveIncrement}>FiveIncrement</button>
    </div>
  );
}
