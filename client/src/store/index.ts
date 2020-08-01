import { combineReducers, createStore } from 'redux';
import gameReducer from './game/reducers';
import userReducer from './user/reducers';
import boardReducer from './board/reducers';

const rootReducer = combineReducers({
  game: gameReducer,
  user: userReducer,
  board: boardReducer,
});

export type RootState = ReturnType<typeof rootReducer>;

export function configureStore() {
  return createStore(rootReducer);
}
