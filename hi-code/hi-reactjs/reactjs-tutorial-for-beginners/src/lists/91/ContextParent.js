import React, { useState } from 'react';
import { ChildrenA, MomeChildrenA } from './ContextChildren';

export const CountContext = React.createContext();
const CountProvider = CountContext.Provider;
const ContextParent = () => {
  const [count, setCount] = useState(0);
  console.log('ContextParent Render');

  return (
    <div>
      Count: {count}
      <br />
      <button onClick={() => setCount((count) => count + 1)}>Click</button>
      {/* <ChildrenA /> */}
      <MomeChildrenA />
      {/* <CountProvider value={count}>
        <MomeChildrenA />
      </CountProvider> */}
    </div>
  );
};

export default ContextParent;
