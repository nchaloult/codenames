import { store, RootState } from '../store/index';
import { SERVER_URL } from '../constants';
import { setIsCreated, setIsJoined } from '../store/game/actions';
import { setUserID } from '../store/user/actions';
import {
  addRedTeamPlayer,
  removeBlueTeamPlayer,
  removeRedTeamPlayer,
  addBlueTeamPlayer,
  changeSomeoneElsesDisplayName,
} from '../store/lobby/actions';
import { UserState } from '../store/user/types';

export enum EventKind {
  lobbyInfo = 'LOBBY_INFO',
  newPlayerID = 'NEW_PLAYER_ID',
  newPlayerJoined = 'NEW_PLAYER_JOINED',
  changeDisplayName = 'CHANGE_DISPLAY_NAME',
  someoneElseChangeDisplayName = 'SOMEONE_ELSE_CHANGE_DISPLAY_NAME',
  changeTeam = 'CHANGE_TEAM',
  someoneElseChangeTeam = 'SOMEONE_ELSE_CHANGE_TEAM',
  notAnEvent = 'NOT_AN_EVENT',
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

// getEventKind makes sure that the provided event object is an Event, and
// returns the EventKind of that event. It's also the future home of logging or
// some kind of universal error handling logic.
export function getEventKind(event: any): EventKind {
  if (!('kind' in event)) {
    return EventKind.notAnEvent;
  }
  return event.kind;
}

// isResponseOK returns whether an event response object is carrying information
// about an something that went wrong. It's also the future home of logging or
// some kind of universal error handling logic for error responses that indicate
// there's a problem somewhere.
export function isResponseOK(response: EventResponse): boolean {
  return response.ok;
}

// getResponseBody returns the body of an EventResponse that indicated something
// was successful. If the provided EventResponse doesn't indicate that something
// was successful, then null is returned.
export function getResponseBody(response: EventResponse): any {
  if (!isResponseOK(response)) {
    return null;
  }
  return response.body;
}
// getResponseErr returns the body of an EventResponse that indicated something
// went wrong somewhere. If the provided EventResponse doesn't indicate that
// something went wrong, then null is returned.
export function getResponseErr(response: EventResponse): any {
  if (isResponseOK(response)) {
    return null;
  }
  return response.body;
}

// notifyOfUnrecognizedEvent is called when we receive a Websocket event from
// the server with an EventKind that we don't recongize.
function notifyOfUnrecognizedEvent(event: Event) {
  // TODO: better error handling.
  alert('Unrecognized Websocket event from the server. Check the console.');
  console.log(event);
}

// handleMsgFromServer parses a Websocket message from the server, and acts
// appropriately depending on that event's kind and the client's state (have we
// joined a game? Is the game we're trying to join created already?)
function handleMsgFromServer(msg: MessageEvent) {
  const globalState = store.getState();
  const event: Event = JSON.parse(msg.data);

  if (globalState.game.isJoined) {
    switch (event.kind) {
      default:
        notifyOfUnrecognizedEvent(event);
        break;
    }
  } else {
    switch (event.kind) {
      case EventKind.newPlayerID:
        // When a client first connects to a game, the server creates a new
        // Player object for them with a UUID. The client needs to include this
        // UUID in the body of select events that it sends.
        store.dispatch(setUserID(event.body.id));
        // When new players join a lobby, the server puts them on red team by
        // default. Show this on the client, as well.
        store.dispatch(addRedTeamPlayer({ id: event.body.id }));
        break;
      case EventKind.newPlayerJoined:
        // When a client joins a game lobby and sets their display name, they'll
        // broadcast a newPlayerJoined event. Add this new Player to the red
        // team.
        store.dispatch(
          addRedTeamPlayer({
            id: event.body.id,
            displayName: event.body.displayName,
          }),
        );
        break;
      case EventKind.lobbyInfo:
        // When a client first visits a /:gameID URL and a Websocket connection
        // with the server is established, the server will send down a lobbyInfo
        // event to tell the client about the game that they're either creating
        // or joining, like what other Players are on which team, for instance.
        store.dispatch(setIsCreated(event.body.isCreated));
        store.dispatch(setIsJoined(false));
        event.body.redTeam.forEach((player: UserState) => {
          if (player.id === store.getState().user.id) {
            return;
          }
          store.dispatch(
            addRedTeamPlayer({
              id: player.id,
              displayName: player.displayName,
            }),
          );
        });
        event.body.blueTeam.forEach((player: UserState) => {
          if (player.id === store.getState().user.id) {
            return;
          }
          store.dispatch(
            addBlueTeamPlayer({
              id: player.id,
              displayName: player.displayName,
            }),
          );
        });
        break;
      case EventKind.someoneElseChangeDisplayName:
        // When another player in a game lobby changes their display name, this
        // event is broadcasted to all other players in that lobby.
        store.dispatch(
          changeSomeoneElsesDisplayName(event.body.id, event.body.displayName),
        );
        break;
      case EventKind.someoneElseChangeTeam:
        // When another player in a game lobby swaps teams, this event is
        // broadcasted to all other players in that lobby.
        if (event.body.isOnRedTeam) {
          store.dispatch(removeBlueTeamPlayer(event.body.id));
          store.dispatch(
            addRedTeamPlayer({
              id: event.body.id,
              displayName: event.body.displayName,
            }),
          );
        } else {
          store.dispatch(removeRedTeamPlayer(event.body.id));
          store.dispatch(
            addBlueTeamPlayer({
              id: event.body.id,
              displayName: event.body.displayName,
            }),
          );
        }
        break;
      default:
        notifyOfUnrecognizedEvent(event);
        break;
    }
  }
}

// establishWSConnection attempts to establish a persistent Websocket connection
// with the server.
export function establishWSConnection(gameID: string): WebSocket {
  const protocol = window.location.protocol === 'https:' ? 'wss://' : 'ws://';
  const socketURL = `${protocol}${SERVER_URL}/ws?gameID=${gameID}`;
  const socket = new WebSocket(socketURL);

  socket.onmessage = (msg) => handleMsgFromServer(msg);
  // TODO: better error handling.
  socket.onerror = (event) => {
    alert(
      'Something went wrong with the Websocket connection to the server. Check the console.',
    );
    console.log(event);
    socket.close(1006);
  };

  return socket;
}
