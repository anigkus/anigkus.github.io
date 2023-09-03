import React,{createContext,useState} from 'react';

export type AuthUser= {
    uname:string,
    email:string
}

type UserContextFourProviderProps = {
    children : React.ReactNode
}
type UserContextFourType = {
   user: AuthUser  ;
   setUser: React.Dispatch<React.SetStateAction<AuthUser>>;
}
export const UserContextFour = createContext<UserContextFourType >({} as UserContextFourType )

export const UserContextFourProvider = ({children}:UserContextFourProviderProps) =>{
    const [user, setUser] = useState<AuthUser>({} as AuthUser);

    return <UserContextFour.Provider value={{user,setUser}}>
        {children}
    </UserContextFour.Provider>
}
