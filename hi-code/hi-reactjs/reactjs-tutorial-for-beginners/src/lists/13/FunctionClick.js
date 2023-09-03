import React from 'react';

export default function FunctionClick() {
  function clickHander() {
    console.log('Console FunctionClick');
  }
  return (
    <div>
      <button onClick={clickHander}>FunctionClick</button>
    </div>
  );
}
