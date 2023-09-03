import React, { useReducer } from 'react';

const initializer = {
  firstValue: 0,
  secondValue: 0
};
type stateProps = {
    firstValue:number,
    secondValue:number
}
// type DispatchProps = {
//     type:string,
//     value:number
// }

type UpdateAction = {
    type:"increment" | "decrement",
    value:number
}

type ResetAction = {
    type:"reset"
}
type DispatchProps = UpdateAction | ResetAction

const reducer = (state:stateProps, dispatch:DispatchProps) => {
  switch (dispatch.type) {
    case 'increment':
      return { ...state, firstValue: state.firstValue + dispatch.value };
    case 'decrement':
      return {
        ...state,
        firstValue: state.firstValue - dispatch.value
      };
    case 'reset':
      return initializer;
    default:
      return state;
  }
};
export default function CounterOne() {
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
        <button onClick={() => dispatch({ type: 'reset' })}>Reset</button>
      </div>
    </div>
  );
}
