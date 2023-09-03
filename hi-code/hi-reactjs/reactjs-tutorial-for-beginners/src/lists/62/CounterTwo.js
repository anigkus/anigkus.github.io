import React, { useReducer } from 'react';

const initializer = {
  firstValue: 0,
  secondValue: 0
};
const reducer = (state, dispatch) => {
  switch (dispatch.type) {
    case 'increment':
      return { ...state, firstValue: state.firstValue + dispatch.value };
    case 'decrement':
      return {
        ...state,
        firstValue: state.firstValue - dispatch.value
      };
    case 'increment2':
      return {
        ...state,
        secondValue: state.secondValue + dispatch.value
      };
    case 'decrement2':
      return {
        ...state,
        secondValue: state.secondValue - dispatch.value
      };
    case 'reset':
      return initializer;
    default:
      return state;
  }
};
export default function CounterTwo() {
  const [value, dispatch] = useReducer(reducer, initializer);
  return (
    <div>
      firstValue:{value.firstValue}, secondValue:{value.secondValue}
      <div>
        <button onClick={() => dispatch({ type: 'increment', value: 1 })}>
          Increment firstValue+1
        </button>
        <button onClick={() => dispatch({ type: 'decrement', value: 1 })}>
          Decrement firstValue-1
        </button>
        <button onClick={() => dispatch({ type: 'increment', value: 5 })}>
          Increment firstValue+5
        </button>
        <button onClick={() => dispatch({ type: 'decrement', value: 5 })}>
          Decrement firstValue-5
        </button>
      </div>
      <div>--------------------------------------------------------</div>
      <div>
        <button onClick={() => dispatch({ type: 'increment2', value: 1 })}>
          Increment secondValue+1
        </button>
        <button onClick={() => dispatch({ type: 'decrement2', value: 1 })}>
          Decrement secondValue-1
        </button>
        <button onClick={() => dispatch({ type: 'increment2', value: 5 })}>
          Increment secondValue+5
        </button>
        <button onClick={() => dispatch({ type: 'decrement2', value: 5 })}>
          Decrement secondValue-5
        </button>
      </div>
      <button onClick={() => dispatch({ type: 'reset' })}>Reset</button>
    </div>
  );
}
