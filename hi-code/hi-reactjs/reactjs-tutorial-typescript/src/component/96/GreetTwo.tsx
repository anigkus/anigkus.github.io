import PersonListTwo from "./PersonListTwo";
import PersonTwo from "./PersonTwo";

type GreetProps ={
    name:string
}
const names = [
  {fname:"Bruce",lname:"A"},
  {fname:"Clark",lname:"B"},
  {fname:"Diana",lname:"C"}
]
const GreetTwo = (props: GreetProps) => {
  return (
    <div>
      Welcome {props.name}<br/>
      <PersonTwo name={props.name} isLoggin={true} flname={{fname:"Anig",lname:"Kus"}} ></PersonTwo>
      <PersonListTwo names={names}></PersonListTwo>
    </div>
  );
}

export default GreetTwo;
