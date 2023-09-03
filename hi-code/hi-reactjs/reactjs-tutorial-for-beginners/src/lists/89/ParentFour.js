import React, { useState, useMemo, useCallback } from 'react';
import ChildrenFour from './ChildrenFour';

const ParentFour = () => {
  const [name, setName] = useState('Anigkus');
  console.log('ParentFour Render');
  const person = {
    fname: 'Anig',
    lname: 'Kus'
  };
  const personHander = () => {};
  //Should not update child component,why
  const MomoPerson = useMemo(() => person, []);
  const memoPersonClick = useCallback(personHander, []);
  return (
    <div>
      Name: {name}
      <br />
      <button onClick={() => setName('Zhang')}>Set Name</button>
      {/* <ChildrenFour name={name} person={person} /> */}
      <ChildrenFour name={name} person={MomoPerson} />
      {/* <ChildrenFour name={name} personClick={memoPersonClick} /> */}
    </div>
  );
};

export default ParentFour;
