import React from 'react';
import {User} from './User.type'

// const User = (props:User) => {
//   return (
//     <div>
//       Name:{props.name},gender: {props.gender}
//     </div>
//   );
// }
const UserOne = ({name,gender}:User) => {
  return (
    <div>
      Name:{name},gender: {gender}
    </div>
  );
}

export default UserOne;
