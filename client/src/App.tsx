import React from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import LandingScreen from './screens/LandingScreen';

const App: React.FC = () => (
  <BrowserRouter>
    <Switch>
      <Route path="/">
        <LandingScreen />
      </Route>
    </Switch>
  </BrowserRouter>
);

export default App;
