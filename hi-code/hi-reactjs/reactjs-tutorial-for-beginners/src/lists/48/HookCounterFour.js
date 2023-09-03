import React, { useState } from 'react';

export default function HookCounterFour() {
  const [items, setItems] = useState([]);
  const addItemHandler = () => {
    setItems([
      ...items,
      {
        id: items.length,
        value: Math.floor(Math.random() * 10) + 1
      }
    ]);
  };
  return (
    <>
      <button onClick={addItemHandler}>AddItem</button>
      <ul>
        {/* {items.map((item, index) => (
          <li key={item.id}>{item.value}</li>
        ))} */}
        {items.map((item) => {
          return <li key={item.id}>{item.value}</li>;
        })}
      </ul>
      <div>{JSON.stringify(items)}</div>
    </>
  );
}
