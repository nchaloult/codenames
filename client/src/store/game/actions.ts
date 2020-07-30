import { GameActionTypes, SET_GAME_ID } from './types';

// Action creator for the SET_GAME_ID action type.
export default function setGameID(id: string): GameActionTypes {
  return {
    type: SET_GAME_ID,
    payload: id,
  };
}
