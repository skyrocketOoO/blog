import React from 'react';

const mockData = [
  { id: 1, first_name: 'John', last_name: 'Doe', email: 'john.doe@example.com' },
  { id: 2, first_name: 'Jane', last_name: 'Smith', email: 'jane.smith@example.com' },
  // Add more mock data items as needed
];

const DataTable = () => {
  return (
    <table>
      <thead>
        <tr>
          <th>First name</th>
          <th>Last name</th>
          <th>Email</th>
        </tr>
      </thead>
      <tbody>
        {mockData.map(row => (
          <tr key={row.id}>
            <td>{row.first_name}</td>
            <td>{row.last_name}</td>
            <td>{row.email}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default DataTable;
