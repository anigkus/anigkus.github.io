import React from 'react';

export default function Person({ person, keyIndex }) {
  return (
    <h1 key={person.id} id={keyIndex}>
      I am {person.name}. I am {person.age} year old. I know {person.skil}.
    </h1>
  );
}
