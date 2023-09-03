import React from 'react';

export const ChildrenTwo = () => {
  console.log('ChildrenTwo Render');
  return <div>Children</div>;
};

export const MomoChildrenTwo = React.memo(ChildrenTwo);
