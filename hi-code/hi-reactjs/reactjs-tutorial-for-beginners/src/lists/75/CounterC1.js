import React from 'react';
import useCounter from '../../hooks/useCounter';

function CounterC1() {
  const [count, increament, decreament, reset] = useCounter(0, 1);
  return (
    <div>
      Count - {count}
      {/* <button onClick={() => increament()}>increament</button>
      <button onClick={() => decreament()}>decreament</button>
      <button onClick={() => reset()}>reset</button> */}
      <button onClick={increament}>increament</button>
      <button onClick={decreament}>decreament</button>
      <button onClick={reset}>reset</button>
    </div>
  );
}

export default CounterC1;
