import React, { Component, useState } from 'react';
import ClassMouseOne from '../52/ClassMouseOne';
import HookMouseOne from '../52/HookMouseOne';
function MouseContainer() {
  const [display, setDisplay] = useState(true);
  return (
    <div>
      <button onClick={() => setDisplay(!display)}>ClickToggle</button>
      {display && <HookMouseOne />}
      {/* {display && <ClassMouseOne />} */}
    </div>
  );
}

export default MouseContainer;
