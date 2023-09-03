import React, { Component } from 'react';

class ClassCounter extends Component {
  constructor(props) {
    super(props);

    this.state = {
      count: 0
    };
  }
  countHandler = () => {
    this.setState({
      count: this.state.count + 1
    });
  };
  render() {
    const { count } = this.state;
    return (
      <div>
        <button onClick={this.countHandler}>Click: {count}</button>
      </div>
    );
  }
}

export default ClassCounter;
