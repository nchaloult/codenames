import { combineReducers, createStore } from 'redux';
import gameReducer from './game/reducers';
import userReducer from './user/reducers';
import boardReducer from './board/reducers';
import lobbyReducer from './lobby/reducers';
import websocketReducer from './websocket/reducers';

const rootReducer = combineReducers({
  game: gameReducer,
  user: userReducer,
  board: boardReducer,
  lobby: lobbyReducer,
  websocket: websocketReducer,
});

export type RootState = ReturnType<typeof rootReducer>;

// configureStore creates and returns a Redux global store object. It's also the
// future home for adding middleware or some kind of logging.
// TODO: figure out how to type the reducer param properly.
function configureStore(reducer: any) {
  return createStore(reducer);
}

export const store = configureStore(rootReducer);
