import React from 'react';
import Column from './Column';

export default function Table() {
  return (
    <table>
      <thead>
        <tr>
          <Column field="name" key="name" />
        </tr>
      </thead>
      <tbody>
        <tr>
          <Column field="age" key="age" />
        </tr>
        <tr>
          <Column field="skil" key="skil" />
        </tr>
      </tbody>
    </table>
  );
}
