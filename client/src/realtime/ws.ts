import { SERVER_URL } from '../constants';

// establishWSConnection attempts to establish a persistent Websocket connection
// with the server.
export function establishWSConnection(gameID: string): WebSocket {
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

export enum EventKind {
  changeDisplayName = 'CHANGE_DISPLAY_NAME',
}

// Event mirrors the structure of JSON messages that the server sends to clients
// via Websocket connections when something occurrs that all clients in a game
// need to be aware of.
interface Event {
  kind: EventKind;
  body: any;
}

// EventResponse mirrors the structure of JSON messages that the server sends to
// a client via a Websocket connection in response to an event that that client
// sent to the server.
interface EventResponse {
  ok: boolean;
  kind: EventKind;
  body?: any;
}

// constructAndSendEvent builds an Event with the provided fields, serializes it
// into a string, and sends it along the provided Websocket connection.
export function constructAndSendEvent(
  socket: WebSocket,
  kind: EventKind,
  body: any,
): void {
  const event: Event = { kind, body };
  socket.send(JSON.stringify(event));
}

// constructAndSendResponse builds an EventResponse with the provided fields,
// serializes it into a string, and sends it along the provided Websocket
// connection. Body parameter may be omitted.
export function constructAndSendResponse(
  socket: WebSocket,
  kind: EventKind,
  body?: any,
): void {
  const response: EventResponse = { ok: true, kind };
  if (body) {
    response.body = body;
  }
  socket.send(JSON.stringify(response));
}

// constructAndSendErr builds an EventResponse with the provided fields,
// serializes it into a string, and sends it along the provided Websocket
// connection. Body parameter may be omitted.
export function constructAndSendErr(
  socket: WebSocket,
  kind: EventKind,
  body?: any,
): void {
  const response: EventResponse = { ok: false, kind };
  if (body) {
    response.body = body;
  }
  socket.send(JSON.stringify(response));
}

// isResponseOK returns whether an event response object is carrying information
// about an something that went wrong. It's also the future home of logging or
// some kind of universal error handling logic for error responses that indicate
// there's a problem somewhere.
export function isResponseOK(response: any): boolean {
  if (!('ok' in response)) {
    return false;
  }
  return response.ok;
}

// getResponseErr returns the body of an EventResponse that indicated something
// went wrong somewhere. If the provided EventResponse doesn't indicate that
// something went wrong, then null is returned.
export function getResponseErr(response: EventResponse): any {
  if (response.ok) {
    return null;
  }
  return response.body;
}
