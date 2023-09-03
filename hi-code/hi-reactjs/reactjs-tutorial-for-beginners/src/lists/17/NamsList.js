import React, { Component } from 'react';
import Person from './Person';

class NamsList extends Component {
  render() {
    const names = ['Bruce', 'Clark', 'Diana'];
    const Persons = [
      {
        id: 1,
        name: 'Bruce',
        age: 30,
        skil: 'Java'
      },
      { id: 2, name: 'Clark', age: 31, skil: 'Python' },
      {
        id: 3,
        name: 'Diana',
        age: 33,
        skil: 'Go'
      }
    ];

    // const nameList = names.map((name, index) => (
    //   <h2 key={index} id={index}>
    //     {name}
    //   </h2>
    // ));
    // return <div>{nameList}</div>;

    const personList = Persons.map((person, index) => (
      <Person key={index} person={person} keyIndex={index} />
    ));
    return <div>{personList}</div>;

    // return (
    //   <div>
    //     {names.map((name, index) => (
    //       <h2 key={index}>{name}</h2>
    //     ))}
    //   </div>
    // );
  }
}

export default NamsList;
