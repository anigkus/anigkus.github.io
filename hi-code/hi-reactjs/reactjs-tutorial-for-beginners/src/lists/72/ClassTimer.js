import React, { Component } from 'react';

class ClassTimer extends Component {
  interval;
  constructor(props) {
    super(props);

    this.state = {
      count: 0
    };
    //const xx = () => {};
  }
  //   xx() {}

  componentDidMount() {
    //this.xx();
    this.interval = setInterval(() => {
      this.setState((prevStat) => ({ count: prevStat.count + 1 }));
    }, 1000);
  }
  componentWillUnmount() {
    //Required call
    clearInterval(this.interval);
  }

  render() {
    return (
      <div>
        Class Count: {this.state.count}
        <button onClick={() => clearInterval(this.interval)}>Reset</button>
      </div>
    );
  }
}

export default ClassTimer;
