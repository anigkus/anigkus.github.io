import React, { useContext } from 'react';
import { CounterAppContext } from './CounterApp';

export default function CounterAppA() {
  const counterAppContext = useContext(CounterAppContext);
  return (
    <div>
      Counter A : {counterAppContext.countStatus}
      <button onClick={() => counterAppContext.counterHander('increment')}>
        Increment
      </button>
      <button onClick={() => counterAppContext.counterHander('decrement')}>
        Decrement
      </button>
      <button onClick={() => counterAppContext.counterHander('reset')}>
        Reset
      </button>
    </div>
  );
}
