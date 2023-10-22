import './App.css';
import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import DataTable from './pages/DataTable';
import Home from './pages/Home';

function App() {
  return (
    <div className="App">
      <Router>
        <Switch>
          <Route path="/datatable" component={DataTable} />
          <Route path="/" component={Home} />
          {/* Other routes if needed */}
        </Switch>
      </Router>
    </div>
  );
}

export default App;
