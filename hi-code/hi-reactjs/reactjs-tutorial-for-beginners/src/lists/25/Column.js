import React from 'react';

export default function Column({ field }) {
  const Datas = [
    {
      name: 'Bruce',
      age: 30,
      skil: 'Java'
    },
    {
      name: 'Clark',
      age: 31,
      skil: 'Devops'
    },
    {
      name: 'Diana',
      age: 33,
      skil: 'Golang'
    }
  ];
  //   const tds = Datas.map((data, index) => (
  //     <td key={Reflect.get(data, field)}>{Reflect.get(data, field)}</td>
  //   ));
  const tds = Datas.map((data, index) => {
    let value = Reflect.get(data, field);
    return <td key={value}>{value}</td>;
  });
  //   return <React.Fragment>{tds}</React.Fragment>;
  return <>{tds}</>;
}
