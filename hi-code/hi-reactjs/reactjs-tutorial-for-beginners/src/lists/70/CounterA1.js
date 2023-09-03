import React, { useState, useMemo } from 'react';

export default function CounterA1() {
  const [countOne, setCountOne] = useState(0);
  const [countTwo, setCountTwo] = useState(0);

  const countOneHander = () => {
    setCountOne(countOne + 1);
  };
  const countTwoHander = () => {
    setCountTwo(countTwo + 1);
  };
  //   const isEven = () => {
  //     let i = 0;
  //     while (i < 20000000) i++;
  //     return countOne % 2 === 0;
  //   };
  const isEven = useMemo(() => {
    let i = 0;
    while (i < 20000000) i++;
    return countOne % 2 === 0;
  }, [countOne]);

  return (
    <div>
      <div>
        <button onClick={countOneHander}>Count-One: {countOne}</button>
        {/* <span>{isEven() ? 'Even' : 'Odd'}</span> */}
        <span>{isEven ? 'Even' : 'Odd'}</span>
      </div>
      <div>
        <button onClick={countTwoHander}>Count-Two: {countTwo}</button>
      </div>
    </div>
  );
}
