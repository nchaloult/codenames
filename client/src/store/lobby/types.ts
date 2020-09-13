// Players are different from Users (UserState in users/types.ts): the User is
// the current client, while Players are all other clients that are
// participating in the same game, or are in the same game lobby.
export interface Player {
  id: string;
  displayName?: string;
}

export interface LobbyState {
  redTeam: Player[];
  blueTeam: Player[];
}

export const CHANGE_DISPLAY_NAME = 'CHANGE_DISPLAY_NAME';
interface ChangeDisplayNameAction {
  type: typeof CHANGE_DISPLAY_NAME;
  payload: { id: string; newName: string };
}
export const SOMEONE_ELSE_CHANGE_DISPLAY_NAME =
  'SOMEONE_ELSE_CHANGE_DISPLAY_NAME';
interface SomeoneElseChangeDisplayNameAction {
  type: typeof SOMEONE_ELSE_CHANGE_DISPLAY_NAME;
  payload: { id: string; newName: string };
}
export const ADD_RED_TEAM_PLAYER = 'ADD_RED_TEAM_PLAYER';
interface AddRedTeamPlayerAction {
  type: typeof ADD_RED_TEAM_PLAYER;
  payload: Player;
}
export const ADD_BLUE_TEAM_PLAYER = 'ADD_BLUE_TEAM_PLAYER';
interface AddBlueTeamPlayerAction {
  type: typeof ADD_BLUE_TEAM_PLAYER;
  payload: Player;
}
export const REMOVE_RED_TEAM_PLAYER = 'REMOVE_RED_TEAM_PLAYER';
interface RemoveRedTeamPlayerAction {
  type: typeof REMOVE_RED_TEAM_PLAYER;
  payload: string;
}
export const REMOVE_BLUE_TEAM_PLAYER = 'REMOVE_BLUE_TEAM_PLAYER';
interface RemoveBlueTeamPlayerAction {
  type: typeof REMOVE_BLUE_TEAM_PLAYER;
  payload: string;
}
export type LobbyActionTypes =
  | ChangeDisplayNameAction
  | SomeoneElseChangeDisplayNameAction
  | AddRedTeamPlayerAction
  | AddBlueTeamPlayerAction
  | RemoveRedTeamPlayerAction
  | RemoveBlueTeamPlayerAction;
