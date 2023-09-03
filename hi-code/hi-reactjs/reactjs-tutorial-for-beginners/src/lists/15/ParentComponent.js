import React, { Component } from 'react';
import ChildComponent from './ChildComponent';

class ParentComponent extends Component {
  constructor(props) {
    super(props);

    this.state = {
      message: 'Parent'
    };
    this.clickHander = this.clickHander.bind(this);
  }
  clickHander(child) {
    alert(`Hello ${this.state.message} from ${child} `);
  }

  render() {
    return (
      <div>
        <ChildComponent clickHander={this.clickHander} />
      </div>
    );
  }
}

export default ParentComponent;
