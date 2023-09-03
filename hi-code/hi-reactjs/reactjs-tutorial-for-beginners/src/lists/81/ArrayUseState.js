import React, { useState } from 'react';

const initialState = ['Bruce', 'Clark'];
const ArrayUseState = () => {
  const [persons, setPersons] = useState(initialState);

  const clickHander = () => {
    //setPersons(() => [...persons, 'Anigkus']);

    const newPersion = [...persons];
    newPersion.push('Anigkus');
    setPersons(newPersion);
  };
  return (
    <div>
      <button onClick={clickHander}>Click</button>
      {persons.map((person, index) => {
        return (
          <div key={index} id={index}>
            {person}
          </div>
        );
      })}
    </div>
  );
};

export default ArrayUseState;
