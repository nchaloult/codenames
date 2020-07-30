import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';

import { BrowserRouter, Switch, Route, Redirect } from 'react-router-dom';
import { Provider } from 'react-redux';
import LandingScreen from './screens/LandingScreen';
import { configureStore } from './store';

const store = configureStore();

const App: React.FC = () => (
  <React.StrictMode>
    <Provider store={store}>
      <BrowserRouter>
        <Switch>
          <Route exact path="/" component={LandingScreen} />
          <Redirect to="/" />
        </Switch>
      </BrowserRouter>
    </Provider>
  </React.StrictMode>
);

ReactDOM.render(<App />, document.getElementById('root'));
