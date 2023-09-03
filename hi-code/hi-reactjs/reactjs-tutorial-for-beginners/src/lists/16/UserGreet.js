import React, { Component } from 'react';

class UserGreet extends Component {
  constructor(props) {
    super(props);

    this.state = {
      isLoggedIn: true
    };
  }

  render() {
    //First
    // return this.state.isLoggedIn && <div>Welcome Anigkus</div>;
    //Second
    // let message;
    // if (this.state.isLoggedIn) {
    //   message = <div>Welcome Anigkus</div>;
    // } else {
    //   message = <div>Welcome Guest</div>;
    // }
    // return message;
    //Third
    // if (this.state.isLoggedIn) {
    //   return <div>Welcome Anigkus</div>;
    // } else {
    //   return <div>Welcome Guest</div>;
    // }
    //Fourth
    return this.state.isLoggedIn ? (
      <div>Welcome Anigkus</div>
    ) : (
      <div>Welcome Guest</div>
    );
  }
}

export default UserGreet;
