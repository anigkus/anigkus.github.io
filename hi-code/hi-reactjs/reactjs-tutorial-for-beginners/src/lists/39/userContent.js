import React from 'react';
// const UserContext = React.createContext();
const UserContext = React.createContext('Guest');
const UserProvide = UserContext.Provider;
const UserConsumer = UserContext.Consumer;

export { UserContext };
export { UserProvide, UserConsumer };
