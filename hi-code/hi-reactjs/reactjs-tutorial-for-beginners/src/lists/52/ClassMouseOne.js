import React, { Component } from 'react';

class ClassMouseOne extends Component {
  constructor(props) {
    super(props);

    this.state = {
      x: 0,
      y: 0
    };
  }
  mouseoverHander = (e) => {
    console.log('*********mouseover***********');
    this.setState({
      x: e.clientX,
      y: e.clientY
    });
  };
  componentDidMount() {
    console.log('*********componentDidMount***********');
    window.addEventListener('mousemove', this.mouseoverHander);
  }
  componentWillUnmount() {
    window.removeEventListener('mousemove', this.mouseoverHander);
  }
  render() {
    return (
      <div>
        Class Mouse X:{this.state.x},Y:{this.state.y}
      </div>
    );
  }
}

export default ClassMouseOne;
