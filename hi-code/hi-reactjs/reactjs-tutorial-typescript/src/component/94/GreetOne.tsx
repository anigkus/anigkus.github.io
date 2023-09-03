
type GreetOneProps ={
    name:string
}
const GreetOne = (props: GreetOneProps) => {
  return (
    <div>
      Welcome {props.name}
    </div>
  );
}

export default GreetOne;
