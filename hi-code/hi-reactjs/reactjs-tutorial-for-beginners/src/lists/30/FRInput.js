import React, { PureComponent } from 'react';

const FRInput = React.forwardRef((prop, ref) => {
  return (
    <div>
      <input type="text" ref={ref} />
    </div>
  );
});
export default FRInput;
