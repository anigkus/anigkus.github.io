import React, { Component } from 'react';

class LifeCycleB extends Component {
  constructor(props) {
    super(props);

    this.state = {
      text: 'LifeCycleB'
    };
    console.log('LifeCycleB constructor');
  }

  static getDerivedStateFromProps(props, state) {
    console.log('LifeCycleB getDerivedStateFromProps');
    return null;
  }
  shouldComponentUpdate(nextProps, nextState) {
    console.log('LifeCycleB shouldComponentUpdate');
    return null;
  }
  getSnapshotBeforeUpdate = (prevProps, prevState) => {
    console.log('LifeCycleB getSnapshotBeforeUpdate');
    return null;
  };
  componentDidUpdate(prevProps, prevState) {
    console.log('LifeCycleB componentDidUpdate');
    return null;
  }

  componentDidMount() {
    console.log('LifeCycleB componentDidMount');
  }

  render() {
    console.log('LifeCycleB render');
    return <div>{this.state.text}</div>;
  }
}

export default LifeCycleB;
