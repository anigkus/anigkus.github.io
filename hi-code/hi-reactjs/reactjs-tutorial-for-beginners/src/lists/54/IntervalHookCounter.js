import React, { useState, useEffect, Component } from 'react';

export default function IntervalHookCounter() {
  const [count, setCount] = useState(1);
  //   const tick = () => {
  //     setCount(count + 1);
  //   };
  useEffect(() => {
    function tick() {
      //setCount((prevCount) => prevCount + 1);
      setCount(count + 1);
    }
    const interval = setInterval(tick, 1000);
    return () => {
      clearInterval(interval);
    };
  }, [count]);

  return <div>{count}</div>;
}
