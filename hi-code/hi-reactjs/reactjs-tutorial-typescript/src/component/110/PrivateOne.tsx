import { LoginOne } from "./LoginOne"
import { ProfileOneProps } from "./ProfileOne"

type PrivateOneProps ={
    isLoggedIn:boolean,
    compenont:React.ComponentType<ProfileOneProps>
}
export const PrivateOne = ({isLoggedIn,compenont:Component}: PrivateOneProps)=> {

    // if(isLoggedIn){
    //     return <Component name="Anigkus"></Component>
    // }else{
    //    return  <LoginOne />
    // }

     return <Component name="Anigkus"></Component>
}