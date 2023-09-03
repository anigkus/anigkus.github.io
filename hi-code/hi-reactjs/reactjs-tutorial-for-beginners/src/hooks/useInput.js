import { useState } from 'react';

function useInput(initialState) {
  const [name, setName] = useState(initialState);

  const bind = {
    value: name,
    onChange: (e) => {
      setName(e.target.value);
    }
  };
  const reset = () => {
    setName(initialState);
  };

  return [name, bind, reset];
}

export default useInput;
