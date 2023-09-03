import logo from './logo.svg';
import './App.css';
//import Greet from './lists/5/Greent';
import { Greet } from './lists/5/Greent';
import Welcome from './lists/6/Welcome';
import Hello from './lists/8/Hello';
import ShowGreet from './lists/9/ShowGreet';
import Message from './lists/10/Message';
import Counter from './lists/11/Counter';
import GreetHello from './lists/12/GreetHello';
import FunctionClick from './lists/13/FunctionClick';
import ClassClick from './lists/13/ClassClick';
import EventBind from './lists/14/EventBind';
import ParentComponent from './lists/15/ParentComponent';
import UserGreet from './lists/16/UserGreet';
import NamsList from './lists/17/NamsList';
import StyleSheet from './lists/20/StyleSheet';
import Inline from './lists/20/Inline';
import app from './lists/20/appStyle.css';
import style from './lists/20/appStyle.module.css';
import Form from './lists/21/Form';
import LifeCycleA from './lists/23/LifeCycleA';
import FragmentDemo from './lists/25/FragmentDemo';
import ParentComp from './lists/26/ParentComp';
import RefsDemo from './lists/28/RefsDemo';
import InputFocus from './lists/29/InputFocus';
import FRParentInput from './lists/30/FRParentInput';
import PortalComponent from './lists/31/PortalComponent';
import Hero from './lists/32/Hero';
import ErrorBoundary from './lists/32/ErrorBoundary';
import ClickCount from './lists/33/ClickCount';
import HoverCount from './lists/33/HoverCount';
import ClickCountTwo from './lists/37/ClickCountTwo';
import HoverCountTwo from './lists/37/HoverCountTwo';
import User from './lists/37/User';
import Count from './lists/37/Count';
import { UserProvide } from './lists/39/userContent';
import ComponentC from './lists/39/ComponentC';
import Posts from './lists/41/Posts';
import PostForm from './lists/41/PostForm';
import ClassCounter from './lists/44/ClassCounter';
import HookCounter from './lists/44/HookCounter';
import HookCounterTwo from './lists/46/HookCounterTwo';
import HookCounterThree from './lists/47/HookCounterThree';
import HookCounterFour from './lists/48/HookCounterFour';
import HookCounterFive from './lists/50/HookCounterFive';
import ClassCounterFive from './lists/50/ClassCounterFive';
import ClassMouseOne from './lists/52/ClassMouseOne';
import HookMouseOne from './lists/52/HookMouseOne';
import MouseContainer from './lists/53/MouseContainer';
import IntervalClassCounter from './lists/54/IntervalClassCounter';
import IntervalHookCounter from './lists/54/IntervalHookCounter';
import DataFetchHook from './lists/55/DataFetchHook';
import CompenontX from './lists/59/CompenontX';
import CounterOne from './lists/62/CounterOne';
import CounterTwo from './lists/62/CounterTwo';
import CounterThree from './lists/62/CounterThree';
import CounterApp from './lists/65/CounterApp';
import FetchDataOne from './lists/66/FetchDataOne';
import FetchDataTwo from './lists/66/FetchDataTwo';
import ParentComponentTwo from './lists/69/ParentComponentTwo';
import CounterA1 from './lists/70/CounterA1';
import InputFocusTwo from './lists/71/InputFocusTwo';
import ClassTimer from './lists/72/ClassTimer';
import HookTimer from './lists/72/HookTimer';
import DocTitleOne from './lists/74/DocTitleOne';
import DocTitleTwo from './lists/74/DocTitleTwo';
import CounterA2 from './lists/75/CounterC1';
import CounterC1 from './lists/75/CounterC1';
import CounterC2 from './lists/75/CounterC2';
import UserForm from './lists/76/UserForm';
import UseStateCounter from './lists/79/UseStateCounter';
import UseReducerCounter from './lists/80/UseReducerCounter';
import ObjectUseState from './lists/81/ObjectUseState';
import ArrayUseState from './lists/81/ArrayUseState';
import ParentTwo from './lists/82/ParentTwo';
import ParentThree from './lists/83/ParentThree';
import { ChildrenThree } from './lists/83/ChildrenThree';
import GrandParent from './lists/83/GrandParent';
import ParentFour from './lists/89/ParentFour';
import ContextParent from './lists/91/ContextParent';

function App() {
  // let user = {};
  // user.name = 'Anigkuis';
  // user.skil = 'Kubernetes';
  let user = 'Anigkus';
  return (
    <div className="App">
      {/* <Greet /> */}
      {/* <Welcome /> */}
      {/* <Hello /> */}
      {/* <ShowGreet name="Aid" greet="Hi">
        <p>This is childern prop</p>
      </ShowGreet>
      <ShowGreet name="Bob" greet="Hello">
        <p>This is childern prop</p>
      </ShowGreet>
      <ShowGreet name="Came" greet="Hey">
        <p>This is childern prop</p>
      </ShowGreet> */}
      {/* <Message /> */}
      {/* <Counter increment="1" /> */}
      {/* <GreetHello name="Anigkus" year="22" /> */}
      {/* <FunctionClick />
      <ClassClick /> */}
      {/* <EventBind /> */}
      {/* <ParentComponent /> */}
      {/* <UserGreet /> */}
      {/* <NamsList /> */}
      {/* <StyleSheet primary="true" />
      <Inline />
      <h1 className="error">Error</h1>
      <h1 className={app.error}>Error2</h1>
      <h2 className={style.success}>Success</h2>
      <h2 className="success">Success2</h2> */}
      {/* <Form /> */}
      {/* <LifeCycleA /> */}
      {/* <FragmentDemo /> */}
      {/* <ParentComp /> */}
      {/* <RefsDemo /> */}
      {/* <InputFocus /> */}
      {/* <FRParentInput /> */}
      {/* <PortalComponent /> */}
      {/* <div>
        <ErrorBoundary>
          <Hero name="Bruce"></Hero>
        </ErrorBoundary>
        <ErrorBoundary>
          <Hero name="Clark"></Hero> 
        </ErrorBoundary>
        <ErrorBoundary>
          <Hero name="Diana"></Hero>
        </ErrorBoundary>
      </div> */}
      {/* <ClickCount name="Anigkus" />
      <HoverCount /> */}

      {/* <ClickCountTwo />
      <HoverCountTwo /> */}
      {/* <Count
        render={(count, incrementHandler) => (
          <ClickCountTwo count={count} incrementHandler={incrementHandler} />
        )}
      ></Count>
      <Count
        render={(count, incrementHandler) => (
          <HoverCountTwo count={count} incrementHandler={incrementHandler} />
        )}
      ></Count> */}
      {/* <User name="Anigkus" /> */}
      {/* <User name={() => 'Anigkus'} /> */}
      {/* <User name={(isLoggin) => (isLoggin ? 'Anigkus' : 'Guest')} /> */}

      {/* <UserProvide value={user}>
        <ComponentC />
      </UserProvide> */}
      {/* <ComponentC /> */}
      {/* <Posts /> */}
      {/* <PostForm /> */}
      {/* <ClassCounter /> */}
      {/* <HookCounter /> */}
      {/* <HookCounterTwo /> */}
      {/* <HookCounterThree /> */}
      {/* <HookCounterFour /> */}

      {/* <ClassCounterFive /> */}
      {/* <HookCounterFive /> */}

      {/* <ClassMouseOne /> */}
      {/* <HookMouseOne /> */}
      {/* <MouseContainer /> */}

      {/* <IntervalClassCounter />
      <IntervalHookCounter /> */}

      {/* <DataFetchHook /> */}

      {/* <CompenontX /> */}

      {/* <CounterOne />/ */}
      {/* <CounterTwo /> */}
      {/* <CounterThree /> */}

      {/* <CounterApp /> */}

      {/* <FetchDataOne /> */}
      {/* <FetchDataTwo /> */}

      {/* <ParentComponentTwo /> */}

      {/* <CounterA1 /> */}

      {/* <InputFocusTwo /> */}

      {/* <ClassTimer />
      <HookTimer /> */}

      {/* <DocTitleOne />
      <DocTitleTwo /> */}

      {/* <CounterC1 />
      <CounterC2 /> */}

      {/* <UserForm /> */}

      {/* <UseStateCounter /> */}

      {/* <UseReducerCounter /> */}

      {/* <ObjectUseState />

      <ArrayUseState /> */}

      {/* <ParentTwo /> */}

      {/* <ParentThree>
        <ChildrenThree />
      </ParentThree> */}

      {/* <GrandParent /> */}

      {/* <ParentFour /> */}
      <ContextParent />
    </div>
  );
}

export default App;
