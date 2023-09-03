import React, { useReducer, useEffect } from 'react';
import axios from 'axios';

const initialState = {
  loading: true,
  post: {},
  error: ''
};
const reducer = (state, action) => {
  switch (action.type) {
    case 'SUCCESS':
      return {
        loading: action.loading,
        post: action.post
      };
    case 'ERROR':
      return {
        loading: action.loading,
        post: {},
        error: 'useEffect jsonplaceholder Error'
      };
    default:
      return state;
  }
};

export default function FetchDataTwo() {
  const [state, dispatch] = useReducer(reducer, initialState);
  useEffect(() => {
    axios
      .get('https://jsonplaceholder.typicode.com/posts/1')
      .then((Response) => {
        dispatch({ type: 'SUCCESS', post: Response.data, loading: false });
      })
      .catch((error) => {
        dispatch({ type: 'ERROR', loading: false });
      });
  }, []);
  return (
    <div>
      {' '}
      {state.loading ? 'loading' : state.post.title}
      {state.error ? state.error : null}
    </div>
  );
}
