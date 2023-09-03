import React, { Component } from 'react';
import FRInput from './FRInput';

class FRParentInput extends Component {
  constructor(props) {
    super(props);

    this.inputRef = React.createRef();
  }

  clickHander = () => {
    this.inputRef.current.focus();
  };
  render() {
    return (
      <div>
        <FRInput ref={this.inputRef} />
        <button onClick={this.clickHander}>Click</button>
      </div>
    );
  }
}

export default FRParentInput;
