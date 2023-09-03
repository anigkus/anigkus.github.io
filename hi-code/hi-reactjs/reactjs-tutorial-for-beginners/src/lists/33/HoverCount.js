import React, { Component } from 'react';
import UpdateWithCount from './withCount';
class HoverCount extends Component {
  constructor(props) {
    super(props);

    this.state = {
      count: 0
    };
  }
  incrementHandler = () => {
    this.setState((proevStat) => {
      return { count: proevStat.count + 1 };
    });
  };

  //   render() {
  //     const { count } = this.state;
  //     return (
  //       <div>
  //         <h1 onMouseOver={this.incrementHandler}>Count {count} X</h1>
  //       </div>
  //     );
  //   }
  // }
  // export default HoverCount;

  render() {
    const { count, incrementHandler } = this.props;
    return (
      <div>
        <h1 onMouseOver={incrementHandler}>Count {count} X</h1>
      </div>
    );
  }
}
export default UpdateWithCount(HoverCount, 10);
