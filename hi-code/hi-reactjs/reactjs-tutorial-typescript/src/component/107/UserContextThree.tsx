import React,{createContext,useState} from 'react';

export type AuthUser= {
    uname:string,
    email:string
}

type UserContextThreeProviderProps = {
    children : React.ReactNode
}
type UserContextThreeType = {
   user: AuthUser  | null ;
   setUser: React.Dispatch<React.SetStateAction<AuthUser | null>>;
}
export const UserContextThree = createContext<UserContextThreeType | null >(null )

export const UserContextThreeProvider = ({children}:UserContextThreeProviderProps) =>{
    const [user, setUser] = useState<AuthUser | null>(null);

    return <UserContextThree.Provider value={{user,setUser}}>
        {children}
    </UserContextThree.Provider>
}
