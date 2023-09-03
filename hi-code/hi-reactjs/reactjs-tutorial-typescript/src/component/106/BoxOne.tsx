import { useContext } from "react"
import {ThemeContext} from './ThemeContextOne'

export const BoxOne = ()=>{
    const themeContext = useContext(ThemeContext)
    return <div style={{backgroundColor:themeContext.secondary.main,color:themeContext.secondary.text}}>Box Theme</div>
}