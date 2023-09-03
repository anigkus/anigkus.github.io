import React, { useState, useEffect } from 'react';

export default function HookCounterFive() {
  const [count, setCount] = useState(0);
  const [name, setName] = useState('');
  useEffect(() => {
    console.log('*******useEffect*********');
    document.title = `You Click ${count} times`;
  }, [count]); // When only count updates
  return (
    <div>
      <input
        type="text"
        value={name}
        onChange={(e) => setName(e.target.value)}
      />
      <button onClick={() => setCount(count + 1)}>
        You Click {count} times
      </button>
    </div>
  );
}
