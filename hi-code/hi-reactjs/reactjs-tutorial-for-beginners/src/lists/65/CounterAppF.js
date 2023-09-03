import React, { useContext } from 'react';
import { CounterAppContext } from './CounterApp';

export default function CounterAppF() {
  const counterAppContext = useContext(CounterAppContext);

  return (
    <div>
      Counter F : {counterAppContext.countStatus}
      <button onClick={() => counterAppContext.counterHander('increment')}>
        Increment
      </button>
      <button onClick={() => counterAppContext.counterHander('decrement')}>
        Decrement
      </button>
      <button onClick={() => counterAppContext.counterHander('reset')}>
        Reset
      </button>
      <p>
        {'-'.repeat(50)}CounterAppContext.Consumer{'-'.repeat(50)}
      </p>
      <CounterAppContext.Consumer>
        {(consume) => {
          return (
            <div>
              CounterAppContext.Consumer count: {consume.countStatus}
              <button onClick={() => consume.counterHander('increment')}>
                Increment
              </button>
              <button onClick={() => consume.counterHander('decrement')}>
                Decrement
              </button>
              <button onClick={() => consume.counterHander('reset')}>
                Reset
              </button>
            </div>
          );
        }}
      </CounterAppContext.Consumer>
    </div>
  );
}
