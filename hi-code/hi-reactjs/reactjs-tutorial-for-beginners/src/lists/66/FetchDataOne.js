import axios from 'axios';
import React, { useState, useEffect } from 'react';

export default function FetchDataOne() {
  const [loading, setLoading] = useState(true);
  const [post, setPost] = useState({});
  const [error, setError] = useState('');

  useEffect(() => {
    axios
      .get('https://jsonplaceholder.typicode.com/posts/1')
      .then((Response) => {
        setLoading(false);
        setPost(Response.data);
      })
      .catch((error) => {
        setLoading(false);
        setPost({});
        setError('useEffect jsonplaceholder Error');
      });
  }, []);

  return (
    <div>
      {loading ? 'loading' : post.title}
      {error ? error : null}
    </div>
  );
}
