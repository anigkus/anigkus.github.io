type PersonTwoProps = {
    name:string,
    messageCount?:number,
    isLoggin:boolean,
    flname:{
        fname:string,
        lname:string
    }
}

const PersonTwo = (props: PersonTwoProps) => {
  const {messageCount = 0}=  props
  return (
    <div>
      {props.isLoggin?`Welcome ${props.name}, you first name ${props.flname.fname} and last name ${props.flname.lname},you have ${messageCount} unrend message`:'Welcome Guest'}
    </div>
  );
}

export default PersonTwo;
