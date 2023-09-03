import React,{createContext} from 'react';
import {ThemeOne} from './ThemeOne';

type ThemeContxtProviderProps = {
    children : React.ReactNode
}
export const ThemeContext = createContext(ThemeOne)

export const ThemeContextProvider = ({children}:ThemeContxtProviderProps) =>{

    return <ThemeContext.Provider value={ThemeOne}>
        {children}
    </ThemeContext.Provider>
}
