import React, { useEffect, useRef } from 'react';

export default function InputFocusTwo() {
  const nameRef = useRef(null);
  useEffect(() => {
    nameRef.current.focus();
  }, []);
  return (
    <div>
      <input type="text" ref={nameRef} />
    </div>
  );
}
