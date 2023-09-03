type PersonListTwoProps = {
    names:{
    fname:string,
    lname:string
   }[]
}
const PersonListTwo = (props: PersonListTwoProps) => {
  return (
    <div>
      {
        props.names.map((name,index) => {
            return (
                <h2 key={index}>{name.fname} {name.lname}</h2>
            )
        })
      }
    </div>
  );
}

export default PersonListTwo;
