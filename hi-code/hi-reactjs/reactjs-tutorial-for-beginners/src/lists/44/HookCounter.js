import React, { useState } from 'react';

export default function HookCounter() {
  const [count, setCount] = useState(0);
  return (
    <div>
      <button onClick={() => setCount(count + 1)}>Click: {count}</button>
    </div>
  );
}
