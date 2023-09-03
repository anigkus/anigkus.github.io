import React, { useState } from 'react';

const initialState = {
  firstName: 'Anig',
  LastName: 'Kus'
};
const ObjectUseState = () => {
  const [person, setPerson] = useState(initialState);

  const clickHander = () => {
    // setPerson(() => ({
    //   firstName: 'Hello',
    //   LastName: 'World'
    // }));
    const newPerson = { ...person };
    newPerson.firstName = 'Hello';
    newPerson.LastName = 'World';
    setPerson(newPerson);
  };
  console.log('ObjectUseState Render');
  return (
    <div>
      {person.firstName} {person.LastName}
      {/* <button onClick={() => clickHander()}>Click</button> */}
      <button onClick={clickHander}>Click</button>
    </div>
  );
};

export default ObjectUseState;
