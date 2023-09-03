import React, { Component } from 'react';

const withCount = (WrappedCount, increment) => {
  class WithCount extends Component {
    constructor(props) {
      super(props);

      this.state = {
        count: 0
      };
    }
    incrementHandler = () => {
      this.setState((proevStat) => {
        return { count: proevStat.count + increment };
      });
    };

    render() {
      return (
        <WrappedCount
          count={this.state.count}
          incrementHandler={this.incrementHandler}
          {...this.props}
        />
      );
    }
  }
  return WithCount;
};

export default withCount;
