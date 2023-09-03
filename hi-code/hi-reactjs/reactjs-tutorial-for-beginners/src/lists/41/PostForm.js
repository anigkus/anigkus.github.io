import axios from 'axios';
import React, { Component } from 'react';

export class PostForm extends Component {
  constructor(props) {
    super(props);

    this.state = {
      userId: '',
      title: '',
      body: ''
    };
  }
  changeHandler = (e) => {
    this.setState({
      [e.target.name]: e.target.value
    });
  };
  submitHandler = (e) => {
    axios
      .post('https://jsonplaceholder.typicode.com/posts', this.state)
      .then((response) => {
        console.log(response);
      })
      .catch((error) => {
        console.log(error);
      });
    console.log(this.state);
    e.preventDefault();
  };

  render() {
    const { userId, title, body } = this.state;
    return (
      <div>
        <form onSubmit={this.submitHandler}>
          <div>
            <label>userId</label>
            <input
              type="text"
              name="userId"
              value={userId}
              onChange={this.changeHandler}
            />
          </div>
          <div>
            <label>title</label>
            <input
              type="text"
              name="title"
              value={title}
              onChange={this.changeHandler}
            />
          </div>
          <div>
            <label>body</label>
            <input
              type="text"
              name="body"
              value={body}
              onChange={this.changeHandler}
            />
          </div>
          <button type="submit">Submit</button>
        </form>
      </div>
    );
  }
}

export default PostForm;
