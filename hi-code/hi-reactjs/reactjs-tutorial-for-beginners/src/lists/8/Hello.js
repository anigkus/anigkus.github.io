import React from 'react';

const Hello = () => {
  //   return (
  //     <div>
  //       <h1>Hello Anigkus</h1>
  //     </div>
  //   );

  return React.createElement(
    'div',
    { id: 'id_hello' },
    React.createElement('h1', null, 'Hello Anigkus')
  );
};

export default Hello;
