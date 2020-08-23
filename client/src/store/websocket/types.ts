export interface WebsocketState {
  socket?: WebSocket;
}

export const SET_SOCKET = 'SET_SOCKET';
interface SetSocketAction {
  type: typeof SET_SOCKET;
  payload: WebSocket;
}
export type WebsocketActionTypes = SetSocketAction;
