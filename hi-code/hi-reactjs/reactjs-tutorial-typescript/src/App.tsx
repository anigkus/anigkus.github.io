import React from 'react';
import logo from './logo.svg';
import './App.css';
import GreetOne from './component/94/GreetOne';
import GreetTwo from './component/96/GreetTwo';
import StatusOne from './component/97/StatusOne';
import HeaderOne from './component/97/HeadingOne';
import OsCarOne from './component/97/OsCarOne';
import ButtonOne from './component/98/ButtonOne';
import InputOne from './component/98/InputOne';
import StyleOne from './component/99/StyleOne';
import UserOne from './component/100/UserOne';
import UserTwo from './component/103/UserTwo';
import CounterOne from './component/104/CounterOne';
import { ThemeContextProvider } from './component/106/ThemeContextOne';
import { BoxOne } from './component/106/BoxOne';
import { UserContextThreeProvider } from './component/107/UserContextThree';
import { UserThree } from './component/107/UserThree';
import { UserContextFourProvider } from './component/107/UserContextFour';
import { UserFour } from './component/107/UserFour';
import { DomRef } from './component/108/DomRef';
import MulableRef from './component/108/MutableRef';
import ClassCounter from './component/109/ClassCounter';
import { PrivateOne } from './component/110/PrivateOne';
import { ProfileOne } from './component/110/ProfileOne';
import { LoginOne } from './component/110/LoginOne';
import { ListOne, ListThree, ListTwo } from './component/111/List';
import { RandomNumberOne } from './component/112/RandomNumberOne';
import { ToastA, ToastB } from './component/113/ToastOne';
import { CustomComponentOne } from './component/115/CustomComponent';

function App() {
  return (
    <div className="App">
      Hello 93 TypeScript
      <br />
      {/* <GreetOne name='Anigkus'/> */}
      {/* <GreetTwo name='Anigkus' /> */}
      {/* <StatusOne status='success'/> 
      <HeaderOne>Header View</HeaderOne>
      <OsCarOne>
      <HeaderOne>OsCarOne Header View</HeaderOne> 
      </OsCarOne> */}
      {/* <ButtonOne handlerOnClick={(event,id) => {
        console.log(event,id)
      } }></ButtonOne>
      <InputOne value='Anigkus' handlerOnChange={(event) => {
        console.log("App InputOne",event.currentTarget.value)
      }}></InputOne> */}
      {/* <StyleOne styles={{fontSize:"30px",color:"Red",border:"1px solid black",padding:"20px",display:"initial"}}></StyleOne> */}
      {/* <UserOne name='Anigkus' gender='Male'/> */}
      {/* <UserTwo/> */}
      {/* <CounterOne/> */}
      {/* <ThemeContextProvider>
        <BoxOne></BoxOne>
      </ThemeContextProvider> */}
      {/* UserContextThreeProvider
      <UserContextThreeProvider>
        <UserThree></UserThree>
      </UserContextThreeProvider> */}
      {/* UserContextFourProvider
      <UserContextFourProvider>
        <UserFour></UserFour>
      </UserContextFourProvider> */}
      {/* <DomRef></DomRef>
      <MulableRef></MulableRef> */}
      {/* <ClassCounter message='Anigkus get'  /> */}
      {/* <PrivateOne isLoggedIn={true}  compenont = {ProfileOne}/> */}
      {/* <ListOne items={["Bruce","Clarn","Diana"]} onClick={(item) => console.log(item)}></ListOne>
      <ListOne items={[1,2,3]} onClick={(item) => console.log(item)}></ListOne>
      <ListOne items={[true,1,"string"]} onClick={(item) => console.log(item)}></ListOne>
      <ListTwo items={[
        {
          id:1,
          name:"Bruce"
        },
        {
          id:2,
          name:"Clarn"
        },{
          id:3,
          name:"Diana"
        }
      ]}  onClick={(item) => console.log(item)}></ListTwo>
      <ListThree items={["Bruce","Clarn","Diana"]} onClick={(item) => console.log(item)}></ListThree>
      <ListThree items={[1,2,3]} onClick={(item) => console.log(item)}></ListThree> */}
      {/* <RandomNumberOne value={0} isZero></RandomNumberOne> */}
      {/* <ToastA position={'center-bottom'}></ToastA>
      <ToastA position={'center'}></ToastA>
      <ToastA position={'center-top'}></ToastA> */}
      {/* <ToastB position={['center-bottom', 'left-top', 'center']}></ToastB> */}
      {/* <ToastB position={['center-bottom', 'left-top', 'left-top']}></ToastB> */}
      <CustomComponentOne name="Anigkus" isLoggin></CustomComponentOne>
    </div>
  );
}

export default App;
