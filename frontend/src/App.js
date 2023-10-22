import logo from './logo.svg';
import './App.css';
import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import DataTable from './DataTable';


function MainPage(){
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}


function App() {
  return (
    <div className="App">
      <Router>
        <Switch>
          <Route path="/datatable" component={DataTable} />
          <Route path="/" component={MainPage} />
          {/* Other routes if needed */}
        </Switch>
      </Router>
    </div>
  );
}

export default App;
