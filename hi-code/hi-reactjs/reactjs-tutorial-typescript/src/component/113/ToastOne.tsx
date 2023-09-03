type HorizontalPositionProps = 'left' | 'center' | 'right';

type VerticalPositionProps = 'top' | 'center' | 'bottom';

type ToastAProps = {
  position:
    | Exclude<
        `${HorizontalPositionProps}-${VerticalPositionProps}`,
        'center-center'
      >
    | 'center';
};

type ToastBProps = {
  position:
    | Exclude<
        `${HorizontalPositionProps}-${VerticalPositionProps}`,
        'center-center'
      >[]
    | 'center';
};
export const ToastA = ({ position }: ToastAProps) => {
  return <div>{position}</div>;
};
export const ToastB = ({ position }: ToastBProps) => {
  return <div>{position}</div>;
};
