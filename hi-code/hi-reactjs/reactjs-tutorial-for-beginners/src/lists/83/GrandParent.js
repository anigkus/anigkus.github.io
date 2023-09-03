import React, { useState } from 'react';
import { ChildrenThree, MomoChildrenThree } from './ChildrenThree';
import ParentThree from './ParentThree';

const GrandParent = () => {
  console.log('GrandParent Render');
  const [count, setCount] = useState(0);
  const [name, setName] = useState('Anigkus');
  return (
    <div>
      GrandParent Count: {count}
      <br />
      GrandParent Name: {name}
      <br />
      <button onClick={() => setCount(count + 1)}>Grand Click Count</button>
      <button onClick={() => setName('Zhang')}>Grand Click Name</button>
      <ParentThree>
        <MomoChildrenThree name={name} />
      </ParentThree>
    </div>
  );
};

export default GrandParent;
