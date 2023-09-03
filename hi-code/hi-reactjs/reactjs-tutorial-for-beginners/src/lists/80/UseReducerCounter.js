import React, { useReducer } from 'react';

const initializer = 0;
const reducer = (state, dispatch) => {
  switch (dispatch) {
    case 'increment':
      return state + 1;
    case 'decrement':
      return state - 1;
    case 'decrement5':
      return 5;
    case 'reset':
      return initializer;
    default:
      return state;
  }
};

const UseReducerCounter = () => {
  console.log('UseReducerCounter Render');
  const [state, dispatch] = useReducer(reducer, initializer);
  return (
    <div>
      Count: {state}
      <button onClick={() => dispatch('increment')}>Increment</button>
      <button onClick={() => dispatch('decrement')}>Decrement</button>
      <button onClick={() => dispatch('decrement5')}>Decrement5</button>
      <button onClick={() => dispatch('reset')}>Reset</button>
    </div>
  );
};

export default UseReducerCounter;
