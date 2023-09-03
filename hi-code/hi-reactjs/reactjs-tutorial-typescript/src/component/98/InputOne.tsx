import React,{useState} from "react";

type InputOneProps = {
    value:string,
    handlerOnChange: (event:React.ChangeEvent<HTMLInputElement>) => void

}

const InputOne = (props: InputOneProps) => {
    const [value, setValue] = useState(props.value);
    const handlerOnChange = (event:React.ChangeEvent<HTMLInputElement>) => {
        console.log("InputOne",event.currentTarget.value)
       setValue(event.currentTarget.value)
    }
  return (
    <div>
        {/* <input value={props.value} onChange={handlerOnChange} /> */}
        <input value={value} onChange={handlerOnChange} />
    </div>
  );
}

export default InputOne;
