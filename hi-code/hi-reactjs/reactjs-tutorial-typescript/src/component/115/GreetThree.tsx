type GreetThreeProps = {
  name: string;
  messageCount?: number;
  isLoggin: boolean;
};

const GreetThree = (props: GreetThreeProps) => {
  const { messageCount = 0 } = props;
  return (
    <div>
      {props.isLoggin
        ? `Welcome ${props.name},you have ${messageCount} unrend message`
        : 'Welcome Guest'}
    </div>
  );
};

export default GreetThree;
