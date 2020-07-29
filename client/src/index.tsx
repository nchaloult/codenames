import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';

import { BrowserRouter, Switch, Route } from 'react-router-dom';
import LandingScreen from './screens/LandingScreen';

const App: React.FC = () => (
  <React.StrictMode>
    <BrowserRouter>
      <Switch>
        <Route path="/">
          <LandingScreen />
        </Route>
      </Switch>
    </BrowserRouter>
  </React.StrictMode>
);

ReactDOM.render(<App />, document.getElementById('root'));
