import { WebsocketActionTypes, SET_SOCKET } from './types';

// Action creator for the SET_SOCKET action type.
export default function setSocket(socket: WebSocket): WebsocketActionTypes {
  return {
    type: SET_SOCKET,
    payload: socket,
  };
}
