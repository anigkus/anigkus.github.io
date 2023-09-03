type HeaderOneProps = {
 children:string
}

const HeaderOne = (props:HeaderOneProps) => {
  return (
    <div>
      {props.children}
    </div>
  );
}

export default HeaderOne;
