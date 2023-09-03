import React, { useState } from 'react';

const UseStateCounter = () => {
  const [count, setCount] = useState(0);
  console.log('UseStateCounter Render');
  return (
    <div>
      Count: {count}
      <button onClick={() => setCount(count + 1)}>Click</button>
      <button onClick={() => setCount(0)}>SetCount(0)</button>
      <button onClick={() => setCount(5)}>SetCount(5)</button>
    </div>
  );
};

export default UseStateCounter;
