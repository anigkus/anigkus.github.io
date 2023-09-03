import React, { Component } from 'react';
import ComponentF from './ComponentF';
import { UserContext } from './userContent';

class ComponentE extends Component {
  static contextType = UserContext;
  render() {
    return (
      <>
        ComponentE {this.context}
        <ComponentF />
      </>
    );
  }
}
// ComponentE.contextType = UserContext;
export default ComponentE;
