import React, { Component } from 'react';

class ClassClick extends Component {
  clickHander() {
    console.log('Console ClassClick ');
  }
  render() {
    return (
      <div>
        {
          //   <button onClick={() => this.clickHander()}>ClassClick</button>
          <button onClick={this.clickHander}>ClassClick</button>
          // callback for load page
          //   <button onClick={this.clickHander()}>ClassClick</button>
        }
      </div>
    );
  }
}

export default ClassClick;
