import React, { Component } from 'react';
import Input from './Input';

class InputFocus extends Component {
  constructor(props) {
    super(props);
    this.inputFocusRef = React.createRef();
  }

  clickHander = () => {
    this.inputFocusRef.current.focusHandler();
  };
  render() {
    return (
      <div>
        <Input ref={this.inputFocusRef} />
        <button onClick={this.clickHander}>Click</button>
      </div>
    );
  }
}

export default InputFocus;
