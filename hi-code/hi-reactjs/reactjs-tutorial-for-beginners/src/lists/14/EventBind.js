import React, { Component } from 'react';

class EventBind extends Component {
  constructor(props) {
    super(props);

    this.state = {
      message: 'Hello'
    };
  }

  //First
  clickHander() {
    this.setState({
      message: 'GoodBye'
    });
    console.log(this);
  }

  //Second
  //   clickHander = () => {
  //     this.setState({
  //       message: 'GoodBye'
  //     });
  //     console.log(this);
  //   };

  render() {
    return (
      <div>
        <div>{this.state.message}</div>
        {/* <button onClick={this.clickHander.bind(this)}>Click</button> */}
        <button onClick={() => this.clickHander()}>Click</button>
        {/* <button onClick={this.clickHander}>Click</button> */}
      </div>
    );
  }
}

export default EventBind;
