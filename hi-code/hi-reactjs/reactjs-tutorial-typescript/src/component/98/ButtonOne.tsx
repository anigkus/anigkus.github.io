import React from "react";

type ButtonOneProps = {
    handlerOnClick: (event: React.MouseEvent< HTMLButtonElement>,id:number) => void
}
const ButtonOne = (props: ButtonOneProps) => {
  return (
    <div>
      <button onClick={(event) => props.handlerOnClick(event,10)}>Click</button>
    </div>
  );
}

export default ButtonOne;
