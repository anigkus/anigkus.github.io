import React, { useState, useCallback, useEffect } from 'react';
import Button from './Button';
import Label from './Label';
import Title from './Title';

export default function ParentComponentTwo() {
  const [age, setAge] = useState(25);
  const [salary, setSalary] = useState(50000);

  // const handlerAge = () => {
  //   setAge((prevAge) => {
  //     return prevAge + 1;
  //   });
  // };
  // const handlerSalary = () => {
  //   setSalary((prevSalry) => prevSalry + 1);
  // };

  const handlerAge = useCallback(() => {
    setAge(age + 1);
  }, [age]);

  const handlerSalary = useCallback(() => {
    setSalary(salary + 1);
  }, [salary]);

  return (
    <div>
      <Title />
      <Label text="Age" count={age} />
      <Button handler={handlerAge}>Age Increment</Button>
      <Label text="Salary" count={salary} />
      <Button handler={handlerSalary}>Salary Increment</Button>
    </div>
  );
}
