import React, { Component } from 'react';

class RegularComp extends Component {
  render() {
    console.log('RegularComp render');
    return <div>{this.props.name}</div>;
  }
}

export default RegularComp;
