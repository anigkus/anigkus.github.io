import React from 'react';
import useInput from '../../hooks/useInput';

export default function UserForm() {
  const [firstName, firstNameHander, resetFirstName] = useInput('');
  const [lastName, lastNameHander, resetLastName] = useInput('');
  const submitHander = (e) => {
    alert(`firstName: ${firstName} , lastName: ${lastName} `);
    e.preventDefault();
  };
  return (
    <div>
      <form onSubmit={submitHander}>
        <div>
          <label>First Name</label>
          <input type="text" {...firstNameHander} />
        </div>
        <div>
          <label>Last Name</label>
          <input type="text" {...lastNameHander} />
        </div>
        <div>
          <button type="submit">Submit</button>
          <button
            type="button"
            onClick={() => {
              resetFirstName();
              resetLastName();
            }}
          >
            Reset
          </button>
        </div>
      </form>
    </div>
  );
}
