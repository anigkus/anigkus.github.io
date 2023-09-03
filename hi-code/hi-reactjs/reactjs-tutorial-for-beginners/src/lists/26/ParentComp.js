import React, { Component, PureComponent } from 'react';
import MemoComp from '../27/MemoComp';
import PureComp from './PureComp';
import RegularComp from './RegularComp';
class ParentComp extends PureComponent {
  //class ParentComp extends Component {
  constructor(props) {
    super(props);

    this.state = {
      name: 'ParentComp'
    };
  }
  asyncSetState() {
    //PureComponent:Call Mult(ParentComp、MemoComp)
    //Component:Call Mult(ParentComp、MemoComp)
    this.setState(
      (prevState, props) => ({
        name: 'Change ParentComp :' + Date.now()
      }),
      () => {
        console.log(this.state.name);
      }
    );
  }
  syncSetState() {
    //PureComponent:Only call once(ParentComp)
    //Component:Call Mult(ParentComp)
    this.setState({
      name: 'Change ParentComp'
    });
  }
  componentDidMount() {
    setInterval(() => {
      this.asyncSetState();
      //this.syncSetState();
    }, 2000);
  }

  render() {
    console.log('ParentComp render');
    return (
      <div>
        <h1>{this.state.name}</h1>
        {/* <RegularComp name="RegularComp" />
        <PureComp name="PureComp" /> */}
        <MemoComp name={this.state.name} />
      </div>
    );
  }
}

export default ParentComp;
