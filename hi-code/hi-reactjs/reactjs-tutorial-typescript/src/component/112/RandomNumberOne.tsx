type NumberProps = {
  value: number;
};
type PositiveProps = NumberProps & {
  isPositive: boolean;
  isNagetive?: never;
  isZero?: never;
};
type NagativeProps = NumberProps & {
  isNagetive: boolean;
  isPositive?: never;
  isZero?: never;
};

type ZeroProps = NumberProps & {
  isZero: boolean;
  isPositive?: never;
  isNagetive?: never;
};

// type RandomNumberOneProps = {
//   value: number;
//   isPositive?: boolean;
//   isNagetive?: boolean;
//   izZero?: boolean;
// };
type RandomNumberOneProps = PositiveProps | NagativeProps | ZeroProps;

export const RandomNumberOne = ({
  value,
  isPositive,
  isNagetive,
  isZero
}: RandomNumberOneProps) => {
  return (
    <div>
      {value},{isPositive && 'Positive'},{isNagetive && 'Nagative'},
      {isZero && 'Zero'}
    </div>
  );
};
