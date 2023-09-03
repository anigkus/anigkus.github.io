
type StyleOneProps ={
    styles:React.CSSProperties
}
const StyleOne = (props:StyleOneProps) => {
  return (
    <div>
      <h1 style={props.styles}>style header</h1>
    </div>
  );
}

export default StyleOne;
