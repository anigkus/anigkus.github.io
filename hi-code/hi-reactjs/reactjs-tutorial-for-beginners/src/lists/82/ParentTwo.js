import React, { useState } from 'react';
import { ChildrenTwo, MomoChildrenTwo } from './ChildrenTwo';

const ParentTwo = () => {
  const [count, setCount] = useState(0);
  console.log('ParentTwo Render');
  return (
    <div>
      Count: {count}
      <button onClick={() => setCount(count + 1)}>Click</button>
      <MomoChildrenTwo />
    </div>
  );
};

export default ParentTwo;
