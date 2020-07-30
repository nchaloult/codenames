import { GameState, GameActionTypes, SET_GAME_ID } from './types';

const initialState: GameState = {
  id: '',
};

export default function gameReducer(
  state = initialState,
  action: GameActionTypes,
): GameState {
  switch (action.type) {
    case SET_GAME_ID:
      return {
        ...state,
        id: action.payload,
      };
    default:
      return state;
  }
}
