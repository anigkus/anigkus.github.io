import React,{useContext} from 'react';
import {UserContextThree} from './UserContextThree'
export const UserThree= () => {
     const context = useContext(UserContextThree);
    const loginHandler = ()=> {
context?.setUser({
            uname:"Anigkus",
            email:"Anigkus@gmail.com"
        })
    }
    const logoutHandler = ()=> {
        context?.setUser(null)
    }
     return <div>
         <button onClick={loginHandler}>Login</button>
      <button onClick={logoutHandler}>Logout</button>
        uname:{context?.user?.uname}, email:{context?.user?.email}
        
     </div>
}