import React, { Component } from 'react';

class RefsDemo extends Component {
  constructor(props) {
    super(props);

    this.nameRef = React.createRef();
    this.passRef = React.createRef();
    this.nameCbRef = null;
    this.setNameCbRef = (element) => {
      this.nameCbRef = element;
    };
  }
  componentDidMount() {
    //this.nameRef.current.focus();
    if (this.nameCbRef) {
      this.nameCbRef.focus();
    }
  }

  clickHander = () => {
    console.log(this.nameRef);
    console.log(this.setNameCbRef); //
    console.log(this.passRef);
  };
  render() {
    return (
      <div>
        <input type="text" ref={this.nameRef} />
        <input type="text" ref={this.setNameCbRef} />
        <input type="password" ref={this.passRef} />
        <button onClick={this.clickHander}>Click</button>
      </div>
    );
  }
}

export default RefsDemo;
