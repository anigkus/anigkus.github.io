import React, { useReducer } from 'react';

const initializer = 0;
const reducer = (state, dispatch) => {
  switch (dispatch) {
    case 'increment':
      return state + 1;
    case 'decrement':
      return state - 1;
    case 'reset':
      return initializer;
    default:
      return state;
  }
};
export default function CounterOne() {
  const [count, dispatch] = useReducer(reducer, initializer);
  return (
    <div>
      count:{count}
      <div>
        <button onClick={() => dispatch('increment')}>Increment</button>
        <button onClick={() => dispatch('decrement')}>Decrement</button>
        <button onClick={() => dispatch('reset')}>Reset</button>
      </div>
    </div>
  );
}
