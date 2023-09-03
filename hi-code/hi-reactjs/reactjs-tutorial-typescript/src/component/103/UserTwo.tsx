import React,{useState} from 'react';
type UserTwoProps = {
    uname:string,
    email:string
}
const UserTwo = () => {
    // const [user, setUser] = useState<UserTwoProps|null>(null);
    const [user, setUser] = useState<UserTwoProps>({} as UserTwoProps)
    const loginHandler = () => {

        setUser({
            uname:"Anigks",
            email:"Anigkus@gmail.com"
        })
    }
    const logoutHandler = ()=> {
        //setUser(null)
        setUser({
            uname:"",
            email:""
        })
    }
  return (
    <div>
      <button onClick={loginHandler}>Login</button>
      <button onClick={logoutHandler}>Logout</button>
      uname:{user.uname},email{user.email}
    </div>
  );
}

export default UserTwo;
