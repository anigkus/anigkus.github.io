import React, { Component } from 'react';
import LifeCycleB from './LifeCycleB';

class LifeCycleA extends Component {
  //Init=>1,
  constructor(props) {
    super(props);

    this.state = {
      text: 'LifeCycleA'
    };
    console.log('LifeCycleA constructor');
  }
  //Init=>2,Change=>1,
  static getDerivedStateFromProps(props, state) {
    console.log('LifeCycleA getDerivedStateFromProps:' + state.text);
    return null;
  }
  //Change=>2,
  shouldComponentUpdate(nextProps, nextState) {
    console.log('LifeCycleA shouldComponentUpdate:' + nextState.text);
    //If need new value for render , requird return nextState
    return nextState;
  }
  //Change=>4,
  getSnapshotBeforeUpdate = (prevProps, prevState) => {
    console.log('LifeCycleA getSnapshotBeforeUpdate:' + prevState.text);
    return null;
  };

  //Change=>5,
  componentDidUpdate(prevProps, prevState) {
    console.log('LifeCycleA componentDidUpdate:' + prevState.text);
    return null;
  }

  //Init=>4,
  componentDidMount() {
    console.log('LifeCycleA componentDidMount');
  }
  changeStatus = () => {
    this.setState(
      (prevState, props) => ({
        text: 'Change LifeCycleA'
      }),
      () => {
        console.log(this.state.text);
      }
    );
  };
  //Init=>3,Change=>3,
  render() {
    console.log('LifeCycleA render');
    return (
      <div>
        <div>{this.state.text}</div>
        <LifeCycleB />
        <button onClick={this.changeStatus}>Change</button>
      </div>
    );
  }
}

export default LifeCycleA;
