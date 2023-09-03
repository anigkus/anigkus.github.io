import React from 'react';

function ChildComponent(props) {
  return (
    <div>
      <button onClick={() => props.clickHander('Child')}>Click</button>
    </div>
  );
}

export default ChildComponent;
