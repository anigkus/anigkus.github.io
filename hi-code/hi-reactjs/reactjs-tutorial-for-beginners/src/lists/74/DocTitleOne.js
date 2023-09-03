import React, { useState } from 'react';
import useDocumenTitle from '../../hooks/useDocumenTitle';

export default function DocTitleOne() {
  const [count, setCount] = useState(0);
  //   useEffect(() => {
  //     document.title = `count ${count}`;
  //   }, [count]);
  useDocumenTitle(count);
  return (
    <div>
      Count : {count}
      <button onClick={() => setCount(count + 1)}>Click Count</button>
    </div>
  );
}
