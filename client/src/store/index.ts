import { combineReducers, createStore } from 'redux';
import gameReducer from './game/reducers';

const rootReducer = combineReducers({
  game: gameReducer,
});

export type RootState = ReturnType<typeof rootReducer>;

export function configureStore() {
  return createStore(rootReducer);
}
