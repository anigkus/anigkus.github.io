import React, { useEffect, useRef, useState } from 'react';

export default function MulableRef() {
  const [count, setCount] = useState(0);
  const intervalRef = useRef<number | null>(null);

  useEffect(() => {
    intervalRef.current = window.setInterval(() => {
      setCount(count + 1);
    }, 1000);

    return () => {
      //Required call
      if(intervalRef.current){
         stopClearInterval()
    }
     
    };
  }, [count]);

  const stopClearInterval = () => {
    if(intervalRef.current){
         window.clearInterval(intervalRef.current);
    }
  };

  return (
    <div>
      Hook Count: {count}
      <button onClick={() => stopClearInterval()}>Reset</button>
    </div>
  );
}
