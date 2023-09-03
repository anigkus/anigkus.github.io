import React, { Component } from 'react';

class IntervalClassCounter extends Component {
  constructor(props) {
    super(props);

    this.state = {
      count: 1
    };
  }
  componentDidMount() {
    this.interval = setInterval(this.tick, 1000);
  }
  componentWillUnmount() {
    clearInterval(this.interval);
  }
  tick = () => {
    this.setState((prevState, prevProps) => ({
      count: prevState.count + 1
    }));
  };

  render() {
    const { count } = this.state;
    return <div>{count}</div>;
  }
}

export default IntervalClassCounter;
