import React, { Component } from 'react';
import ReactDOM from 'react-dom';
// class PortalComponent extends Component {
//   render() {
//     return <div>Portal Component Web</div>;
//   }
// }

function PortalComponent() {
  return ReactDOM.createPortal(
    <div>Portal Component Web</div>,
    document.getElementById('portal-root')
  );
}

export default PortalComponent;
