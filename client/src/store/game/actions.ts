import {
  GameActionTypes,
  SET_GAME_ID,
  SET_IS_CREATED,
  SET_IS_JOINED,
} from './types';

// Action creator for the SET_GAME_ID action type.
export function setGameID(id: string): GameActionTypes {
  return {
    type: SET_GAME_ID,
    payload: id,
  };
}

// Action creator for the SET_IS_CREATED action type.
export function setIsCreated(isCreated: boolean): GameActionTypes {
  return {
    type: SET_IS_CREATED,
    payload: isCreated,
  };
}

// Action creator for the SET_IS_JOINED action type.
export function setIsJoined(isJoined: boolean): GameActionTypes {
  return {
    type: SET_IS_JOINED,
    payload: isJoined,
  };
}
