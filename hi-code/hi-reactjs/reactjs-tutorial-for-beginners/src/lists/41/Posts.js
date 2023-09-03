import React, { Component } from 'react';
import axis from 'axios';

export class Posts extends Component {
  constructor(props) {
    super(props);

    this.state = {
      posts: [],
      error: ''
    };
  }

  componentDidMount() {
    axis
      .get('https://jsonplaceholder.typicode.com/posts')
      .then((response) => {
        console.log(response);
        this.setState({
          posts: response.data
        });
      })
      .catch((error) => {
        this.setState({
          error: error
        });
      });
  }

  render() {
    const style = {
      color: 'red'
    };
    const { posts, error } = this.state;
    return (
      <div>
        {/* {posts.length
          ? posts.map((post, index) => <div key={post.id}>{post.title}</div>)
          : null} */}
        {posts.length
          ? posts.map((post, index) => {
              let id = `post_${index}`;
              return (
                <div key={post.id} id={id}>
                  {post.title}
                </div>
              );
            })
          : null}
        {error ? <div style={style}>{error}</div> : null}
      </div>
    );
  }
}

export default Posts;
