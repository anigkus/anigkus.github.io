import React, { useEffect, useRef, useState } from 'react';

export default function HookTimer() {
  const [count, setCount] = useState(0);
  const intervalRef = useRef();

  useEffect(() => {
    intervalRef.current = setInterval(() => {
      setCount((prevCount) => prevCount + 1);
      //setCount(count + 1); //[count]
    }, 1000);
    return () => {
      //Required call
      clearInterval(intervalRef.current);
    };
  }, [count]);

  //const test1 = () => {};

  return (
    <div>
      Hook Count: {count}
      <button onClick={() => clearInterval(intervalRef.current)}>Reset</button>
    </div>
  );
}
