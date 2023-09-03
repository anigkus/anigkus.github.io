import React, { useContext } from 'react';
import { CountContext } from './ContextParent';

export const ChildrenA = () => {
  console.log('ChildrenA Render');
  return (
    <div>
      ChildrenA
      <ChildrenB />
    </div>
  );
};
export const MomeChildrenA = React.memo(ChildrenA);
export const ChildrenB = () => {
  console.log('ChildrenB Render');
  return (
    <div>
      ChildrenB
      <ChildrenC />
    </div>
  );
};

export const ChildrenC = () => {
  const context = useContext(CountContext);
  console.log('ChildrenC Render');
  return (
    <div>
      ChildrenC
      <br />
      <div>count: {context}</div>
    </div>
  );
};
