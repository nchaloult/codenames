import {
  GameState,
  GameActionTypes,
  SET_GAME_ID,
  SET_IS_CREATED,
} from './types';

const initialState: GameState = {
  id: '',
  isCreated: false,
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
    case SET_IS_CREATED:
      return {
        ...state,
        isCreated: action.payload,
      };
    default:
      return state;
  }
}
