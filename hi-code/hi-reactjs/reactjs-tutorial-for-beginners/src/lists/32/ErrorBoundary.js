import React, { Component } from 'react';

class ErrorBoundary extends Component {
  constructor(props) {
    super(props);

    this.state = {
      isError: false
    };
  }
  static getDerivedStateFromError(error) {
    return {
      isError: true
    };
  }
  componentDidCatch(error, info) {
    console.log(error);
    console.log(info);
  }

  render() {
    const style = {
      color: 'Red'
    };
    if (this.state.isError) {
      return <h1 style={style}>Something went error</h1>;
    }
    return this.props.children;
  }
}

export default ErrorBoundary;
