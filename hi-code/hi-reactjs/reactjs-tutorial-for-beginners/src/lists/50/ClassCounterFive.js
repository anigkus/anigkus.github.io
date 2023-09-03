import React, { Component } from 'react';

class ClassCounterFive extends Component {
  constructor(props) {
    super(props);

    this.state = {
      count: 0,
      name: ''
    };
  }

  componentDidMount() {
    document.title = `You Click ${this.state.count} timer`;
  }
  componentDidUpdate(prevProps, prevState) {
    if (prevState.count !== this.state.count) {
      console.log('*******componentDidUpdate*********');
      document.title = `You Click ${this.state.count} timer`;
    }
  }

  render() {
    const { count, name } = this.state;
    return (
      <div>
        <input
          type="text"
          value={name}
          onChange={(e) => {
            this.setState({
              name: e.target.value
            });
          }}
        />
        {/* <button
          onClick={() => {
            this.setState({
              count: count + 1
            });
          }}
        >
          You Click {count} timer
        </button> */}
        <button
          onClick={() => {
            this.setState((prevState, prevProps) => ({
              count: prevState.count + 1
            }));
          }}
        >
          You Click {count} timer
        </button>
      </div>
    );
  }
}

export default ClassCounterFive;
