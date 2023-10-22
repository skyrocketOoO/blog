<script>
    import { DataHandler } from './datahandler'; // Make sure the path to DataHandler.js is correct
    import data from './data'; // Assuming you have a data.js file with your data
  

    let rowsArray = [];
    const unsubscribe = rows.subscribe(value => {
        rowsArray = value;
    });
    const handler = new DataHandler(data, { rowsPerPage: 10 });
    let rows = handler.getRows();

    onDestroy(() => {
        unsubscribe();
    });
  </script>
  
  <table>
    <thead>
      <tr>
        <th>First name</th>
        <th>Last name</th>
        <th>Email</th>
      </tr>
    </thead>
    <tbody>
      {#each rows as row}
        <tr>
          <td>{row.first_name}</td>
          <td>{row.last_name}</td>
          <td>{row.email}</td>
        </tr>
      {/each}
    </tbody>
  </table>
  
  <style>
    table {
      text-align: left;
      border-collapse: separate;
      border-spacing: 0;
      width: 100%;
    }
    td, th {
      padding: 4px 20px;
      border-bottom: 1px solid #eee;
    }
  </style>
  