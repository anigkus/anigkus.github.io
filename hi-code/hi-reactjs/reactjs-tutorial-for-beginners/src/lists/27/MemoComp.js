import React, { Component } from 'react';

class MemoComp extends Component {
  render() {
    console.log('MemoComp render');
    return <div>{this.props.name}</div>;
  }
}

export default React.memo(MemoComp);
