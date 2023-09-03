import React, { useContext } from 'react';
import CompenontZ from './CompenontZ';
import { NameContext, ChannerContext } from './CompenontX';

export default function CompenontY() {
  const name = useContext(NameContext);
  const channer = useContext(ChannerContext);

  return (
    <div>
      {/* <CompenontZ/> */} {name},{channer}
    </div>
  );
}
