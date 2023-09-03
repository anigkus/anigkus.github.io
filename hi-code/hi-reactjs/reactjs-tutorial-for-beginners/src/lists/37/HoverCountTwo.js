import React, { Component } from 'react';

class HoverCountTwo extends Component {
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
        {/* <h1 onMouseOver={this.incrementHandler}>Hover {count} Increment</h1> */}
        <h1 onMouseOver={incrementHandler}>Hover {count} Increment</h1>
      </div>
    );
  }
}

export default HoverCountTwo;
