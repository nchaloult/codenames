import { WebsocketActionTypes, WebsocketState, SET_SOCKET } from './types';

const initialState: WebsocketState = {
  socket: undefined,
};

export default function websocketReducer(
  state = initialState,
  action: WebsocketActionTypes,
): WebsocketState {
  switch (action.type) {
    case SET_SOCKET:
      return {
        ...state,
        socket: action.payload,
      };
    default:
      return state;
  }
}
