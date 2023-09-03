import React, { Component } from 'react';

//First
class ShowGreet extends Component {
  render() {
    return (
      <div>
        <h1>
          {this.props.greet} {this.props.name}
        </h1>
        {this.props.children}
      </div>
    );
  }
}

//Second
// const ShowGreet = (props) => {
//   return <h1>{props.greet} {props.name}</h1>;
// };

//Three
// const ShowGreet = ({ name, greet }) => {
//   return (
//     <h1>
//       {greet} {name}
//     </h1>
//   );
// };

export default ShowGreet;
