import React, { useState, useEffect } from 'react';
import axios from 'axios';

export default function DataFetchHook() {
  //   const [posts, setPosts] = useState([]);
  //   useEffect(() => {
  //     axios
  //       .get('https://jsonplaceholder.typicode.com/posts')
  //       .then((rsp) => {
  //         setPosts(rsp.data);
  //       })
  //       .catch((error) => {
  //         console.log(error);
  //       });
  //   });
  const [post, setPost] = useState({});
  const [id, setId] = useState(1);
  const [useButtonFetch, setUseButtonFetch] = useState(1);
  useEffect(() => {
    console.log('useEffect');
    axios
      .get(`https://jsonplaceholder.typicode.com/posts/${useButtonFetch}`)
      .then((res) => {
        setPost(res.data);
      })
      .catch((err) => {
        console.log(err);
      });
  }, [useButtonFetch]);
  return (
    <div>
      {/* <ul>
        {posts.map((post, index) => {
          let id = `post_${post.id}`;
          return (
            <li key={index} id={id}>
              {post.title}
            </li>
          );
        })}
        
      </ul> */}
      <input type={'text'} value={id} onChange={(e) => setId(e.target.value)} />
      <button onClick={() => setUseButtonFetch(id)}>FetchData</button>
      {post.title}
    </div>
  );
}
