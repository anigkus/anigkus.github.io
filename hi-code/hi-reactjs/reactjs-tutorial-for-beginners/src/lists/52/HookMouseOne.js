import React, { useEffect, useState } from 'react';

export default function HookMouseOne() {
  const [x, setX] = useState(0);
  const [y, setY] = useState(0);

  const mouseoverHander = (e) => {
    console.log('*********mouseover***********');
    setX(e.clientX);
    setY(e.clientY);
  };
  useEffect(() => {
    console.log('*********useEffect***********');
    window.addEventListener('mousemove', mouseoverHander);
    return () => {
      console.log('*********useEffect componentWillUnmount***********');
      window.removeEventListener('mousemove', mouseoverHander);
    };
  }, []); //Nothing deps
  return (
    <div>
      Hook Mouse X:{x},Y:{y}
    </div>
  );
}
