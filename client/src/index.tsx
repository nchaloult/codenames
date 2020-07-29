import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';

import { BrowserRouter, Switch, Route, Redirect } from 'react-router-dom';
import LandingScreen from './screens/LandingScreen';

const App: React.FC = () => (
  <React.StrictMode>
    <BrowserRouter>
      <Switch>
        <Route exact path="/" component={LandingScreen} />
        <Redirect to="/" />
      </Switch>
    </BrowserRouter>
  </React.StrictMode>
);

ReactDOM.render(<App />, document.getElementById('root'));
