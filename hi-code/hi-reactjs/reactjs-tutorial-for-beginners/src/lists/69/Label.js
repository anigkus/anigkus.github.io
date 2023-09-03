import React from 'react';

function Label({ text, count }) {
  console.log('Label' + '-' + text + '-' + count);
  return (
    <label>
      Label: {text}- {count}
    </label>
  );
}
export default React.memo(Label);
