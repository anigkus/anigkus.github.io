import React, { Component } from 'react';
import { UserConsumer } from './userContent';

class ComponentF extends Component {
  render() {
    return (
      <UserConsumer>
        {(user) => {
          return (
            <div>
              {/* Hello {user.name},{user.skil} */}
              Hello {user}
            </div>
          );
        }}
      </UserConsumer>
    );
  }
}

export default ComponentF;
