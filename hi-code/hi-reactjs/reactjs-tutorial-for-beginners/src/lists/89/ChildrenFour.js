import React from 'react';

const ChildrenFour = ({ name, person }) => {
  console.log('ChildrenFour Render');
  return (
    <div>
      Parent Name: {name}
      {/* fName:{person.fname} */}
    </div>
  );
};

export default ChildrenFour;
