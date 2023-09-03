import React, { useState } from 'react';
import { ChildrenThree, MomoChildrenThree } from './ChildrenThree';

const ParentThree = ({ children }) => {
  const [count, setCount] = useState(0);
  console.log('ParentTwo Render');
  return (
    <div>
      ParentThree Count: {count}
      <button onClick={() => setCount(count + 1)}>Click</button>
      {children}
    </div>
  );
};

export default ParentThree;
