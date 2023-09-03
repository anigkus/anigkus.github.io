import React, { Component } from 'react';

class Count extends Component {
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
    // return <div>{this.props.render(this.state.count, this.incrementHandler)}</div>;
    // return <>{this.props.render(this.state.count, this.incrementHandler)}</>;
    return (
      <React.Fragment>
        {this.props.render(this.state.count, this.incrementHandler)}
      </React.Fragment>
    );
  }
}

export default Count;
