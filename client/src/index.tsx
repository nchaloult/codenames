import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';

import { BrowserRouter, Switch, Route, Redirect } from 'react-router-dom';
import { Provider } from 'react-redux';
import { configureStore } from './store';
import LandingScreen from './screens/LandingScreen';
import GameScreen from './screens/GameScreen';

const store = configureStore();

const App: React.FC = () => (
  <React.StrictMode>
    <Provider store={store}>
      <BrowserRouter>
        <Switch>
          <Route exact path="/:gameID" component={GameScreen} />
          <Route exact path="/" component={LandingScreen} />
          <Redirect to="/" />
        </Switch>
      </BrowserRouter>
    </Provider>
  </React.StrictMode>
);

ReactDOM.render(<App />, document.getElementById('root'));
