import React,{useContext} from 'react';
import {UserContextFour} from './UserContextFour'
export const UserFour= () => {
     const context = useContext(UserContextFour);
    const loginHandler = ()=> {
context.setUser({
            uname:"Anigkus",
            email:"Anigkus@gmail.com"
        })
    }
    const logoutHandler = ()=> {
        context.setUser({
            uname:"",
            email:""
        })
    }
     return <div>
         <button onClick={loginHandler}>Login</button>
      <button onClick={logoutHandler}>Logout</button>
        uname:{context.user.uname},email:{context.user.email}
        
     </div>
}