import { SERVER_URL } from '../constants';

// establishWSConnection attempts to establish a persistent Websocket connection
// with the server.
export default function establishWSConnection(gameID: string): WebSocket {
  const protocol = window.location.protocol === 'https:' ? 'wss://' : 'ws://';
  const socketURL = `${protocol}${SERVER_URL}/ws?gameID=${gameID}`;
  const socket = new WebSocket(socketURL);

  // TODO: better error handling.
  socket.onerror = () => {
    alert('Something went wrong with the Websocket connection to the server');
    socket.close(1006);
  };

  return socket;
}
