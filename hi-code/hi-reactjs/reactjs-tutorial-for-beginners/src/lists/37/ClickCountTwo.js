import React, { Component } from 'react';

class ClickCountTwo extends Component {
  constructor(props) {
    super(props);

    this.state = {
      count: 0
    };
  }
  incrementHandler = () => {
    this.setState((prevstat) => {
      return { count: prevstat.count + 1 };
    });
  };

  render() {
    // const { count } = this.state;
    const { count, incrementHandler } = this.props;
    return (
      <div>
        {/* <button onClick={this.incrementHandler}>Click {count} Increment</button> */}
        <button onClick={incrementHandler}>Click {count} Increment</button>
      </div>
    );
  }
}

export default ClickCountTwo;
