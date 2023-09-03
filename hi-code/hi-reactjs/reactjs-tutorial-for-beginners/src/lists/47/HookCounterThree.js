import React, { useState } from 'react';

export default function HookCounterThree() {
  const [name, setName] = useState({ firstName: '', lastName: '' });
  return (
    <div>
      <input
        value={name.firstName}
        onChange={(e) => setName({ ...name, firstName: e.target.value })}
      />
      <input
        value={name.lastName}
        onChange={(e) => setName({ ...name, lastName: e.target.value })}
      />
      <h2>Your first name is :{name.firstName}</h2>
      <h2>Your last name is :{name.lastName}</h2>
      <h1>{JSON.stringify(name)}</h1>
    </div>
  );
}
