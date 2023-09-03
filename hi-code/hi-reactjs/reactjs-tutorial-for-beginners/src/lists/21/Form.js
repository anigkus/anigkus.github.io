import React, { Component } from 'react';

class Form extends Component {
  constructor(props) {
    super(props);

    this.state = {
      username: '',
      comment: '',
      topic: 'Kubernetes'
    };
  }
  handlerUserNameChange = (event) => {
    this.setState({
      username: event.target.value
    });
  };
  handlerCommentChange = (event) => {
    this.setState({
      comment: event.target.value
    });
  };
  handlerTopicChange = (event) => {
    this.setState({
      topic: event.target.value
    });
  };

  submitForm = (event) => {
    alert(`${this.state.username},${this.state.comment} ,${this.state.topic} `);
    event.preventDefault();
  };
  render() {
    const { username, comment, topic } = this.state;
    return (
      <div>
        <form onSubmit={this.submitForm}>
          <div>
            <label>UserName</label>
            <input
              type="text"
              value={username}
              onChange={this.handlerUserNameChange}
            />
          </div>
          <div>
            <label>Comment</label>
            <textarea value={comment} onChange={this.handlerCommentChange} />
          </div>
          <div>
            <label>Topic</label>
            <select value={topic} onChange={this.handlerTopicChange}>
              <option value="Java">Java</option>
              <option value="Golang">Golang</option>
              <option value="Kubernetes">Kubernetes</option>
            </select>
          </div>
          <div>
            <button type="submit">Submit</button>
          </div>
        </form>
      </div>
    );
  }
}

export default Form;
