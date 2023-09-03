import React, { useContext, useReducer } from 'react';
import CounterAppA from './CounterAppA';
import CounterAppB from './CounterAppB';
import CounterAppC from './CounterAppC';

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
export const CounterAppContext = React.createContext();
export default function CounterApp() {
  const [count, dispatch] = useReducer(reducer, initializer);
  return (
    <CounterAppContext.Provider
      value={{ countStatus: count, counterHander: dispatch }}
    >
      <div>
        Count:{count}
        <CounterAppA />
        <CounterAppB />
        <CounterAppC />
      </div>
    </CounterAppContext.Provider>
  );
}
