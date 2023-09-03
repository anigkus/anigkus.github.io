import React from 'react';

export const ChildrenThree = ({ name }) => {
  console.log('ChildrenThree Render');
  return <div>Children: parent Name:{name}</div>;
};

export const MomoChildrenThree = React.memo(ChildrenThree);
