import React, { Component } from 'react';

class Message extends Component {
  constructor() {
    super();
    this.state = {
      message: 'Welcome Anigkus'
    };
  }

  changeMessage() {
    this.setState({ message: 'Thanks for you subscribing' }, () => {
      console.log(this.state.message);
    });
  }
  render() {
    return (
      <div>
        <h1>{this.state.message}</h1>
        <button
          onClick={() => {
            this.changeMessage();
          }}
        >
          Subscribie
        </button>
      </div>
    );
  }
}

export default Message;
