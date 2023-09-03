import React, { Component } from 'react';

class GreetHello extends Component {
  //First
  // render() {
  //   return (
  //     <div>
  //       This is {this.props.name}, is {this.props.year} years old.
  //     </div>
  //   );
  // }
  //Second
  constructor(props) {
    super(props);
    // const { name, year } = props;
    // this.state = {
    //   name: name,
    //   year: year
    // };
    this.state = {
      name: this.props.name,
      year: this.props.year
    };
  }

  render() {
    return (
      <div>
        This is {this.state.name}, is {this.state.year} years old.
      </div>
    );
  }
  //Third
  // render() {
  //   const { name, year } = this.props;
  //   return (
  //     <div>
  //       This is {name}, is {year} years old.
  //     </div>
  //   );
  // }
}

//Fourth
// const GreetHello = (props) => {
//   return (
//     <div>
//       This is {props.name}, is {props.year} years old.
//     </div>
//   );
// };

//Fifth
// const GreetHello = (props) => {
//   const { name, year } = props;
//   return (
//     <div>
//       This is {name}, is {year} years old.
//     </div>
//   );
// };

//Sixth
// const GreetHello = ({ name, year }) => {
//   return (
//     <div>
//       This is {name}, is {year} years old.
//     </div>
//   );
// };

export default GreetHello;
