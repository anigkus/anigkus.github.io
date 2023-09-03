import React from 'react';

function Title() {
  console.log('Title');
  return <div>UseCallback Hook</div>;
}
export default React.memo(Title);
