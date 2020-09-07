import {
  LobbyActionTypes,
  ADD_RED_TEAM_PLAYER,
  Player,
  ADD_BLUE_TEAM_PLAYER,
  REMOVE_RED_TEAM_PLAYER,
  REMOVE_BLUE_TEAM_PLAYER,
  CHANGE_DISPLAY_NAME,
} from './types';

// Action creator for the CHANGE_DISPLAY_NAME action type.
export function changeDisplayName(
  id: string,
  newName: string,
): LobbyActionTypes {
  return {
    type: CHANGE_DISPLAY_NAME,
    payload: { id, newName },
  };
}

// Action creator for the ADD_RED_TEAM_PLAYER action type.
export function addRedTeamPlayer(player: Player): LobbyActionTypes {
  return {
    type: ADD_RED_TEAM_PLAYER,
    payload: player,
  };
}

// Action creator for the ADD_BLUE_TEAM_PLAYER action type.
export function addBlueTeamPlayer(player: Player): LobbyActionTypes {
  return {
    type: ADD_BLUE_TEAM_PLAYER,
    payload: player,
  };
}

// Action creator for the REMOVE_RED_TEAM_PLAYER action type.
export function removeRedTeamPlayer(id: string): LobbyActionTypes {
  return {
    type: REMOVE_RED_TEAM_PLAYER,
    payload: id,
  };
}

// Action creator for the REMOVE_BLUE_TEAM_PLAYER action type.
export function removeBlueTeamPlayer(id: string): LobbyActionTypes {
  return {
    type: REMOVE_BLUE_TEAM_PLAYER,
    payload: id,
  };
}
