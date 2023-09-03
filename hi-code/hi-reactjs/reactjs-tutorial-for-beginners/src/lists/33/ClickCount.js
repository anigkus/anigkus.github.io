import React, { Component } from 'react';
import UpdateWithCount from './withCount';
class ClickCount extends Component {
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
  //         <button onClick={this.incrementHandler}>Click {count} X</button>
  //       </div>
  //     );
  //   }
  // }
  // export default ClickCount;

  render() {
    const { count, incrementHandler } = this.props;
    return (
      <div>
        <button onClick={incrementHandler}>
          {this.props.name} Click {count} X
        </button>
      </div>
    );
  }
}
export default UpdateWithCount(ClickCount, 5);
