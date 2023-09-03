import React, { Component } from 'react';

class Counter extends Component {
  constructor(props) {
    super(props);

    this.state = {
      count: 0
    };
  }

  incrementHander() {
    // this.setState(
    //   {
    //     count: this.state.count + 1
    //   },
    //   () => {
    //     console.log('Callback value', this.state.count);
    //   }
    // );
    this.setState(
      (prevState, props) => ({
        count: prevState.count + parseInt(props.increment)
      }),
      () => {
        console.log('Callback value', this.state.count);
      }
    );
  }
  incrementHanderFive() {
    this.incrementHander();
    this.incrementHander();
    this.incrementHander();
    this.incrementHander();
    this.incrementHander();
  }

  render() {
    return (
      <div>
        <div>
          Count: {this.state.count},Increment: {this.props.increment}
        </div>
        <button onClick={() => this.incrementHanderFive()}>Increment</button>
      </div>
    );
  }
}

export default Counter;
