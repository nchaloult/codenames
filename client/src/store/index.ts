import { combineReducers, createStore } from 'redux';
import gameReducer from './game/reducers';
import userReducer from './user/reducers';

const rootReducer = combineReducers({
  game: gameReducer,
  user: userReducer,
});

export type RootState = ReturnType<typeof rootReducer>;

export function configureStore() {
  return createStore(rootReducer);
}
