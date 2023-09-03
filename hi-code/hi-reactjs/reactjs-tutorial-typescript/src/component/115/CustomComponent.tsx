import GreetThree from './GreetThree';

export const CustomComponentOne = (
  props: React.ComponentProps<typeof GreetThree>
) => {
  return <div>{props.name}</div>;
};
