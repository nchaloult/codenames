import {
  LobbyState,
  LobbyActionTypes,
  ADD_RED_TEAM_PLAYER,
  ADD_BLUE_TEAM_PLAYER,
  REMOVE_RED_TEAM_PLAYER,
  REMOVE_BLUE_TEAM_PLAYER,
} from './types';

const initialState: LobbyState = {
  redTeam: [],
  blueTeam: [],
};

export default function lobbyReducer(
  state = initialState,
  action: LobbyActionTypes,
): LobbyState {
  switch (action.type) {
    case ADD_RED_TEAM_PLAYER:
      return {
        ...state,
        redTeam: [...state.redTeam, action.payload],
      };
    case ADD_BLUE_TEAM_PLAYER:
      return {
        ...state,
        blueTeam: [...state.blueTeam, action.payload],
      };
    case REMOVE_RED_TEAM_PLAYER:
      return {
        ...state,
        redTeam: state.redTeam.filter(
          (player) => player.displayName !== action.payload,
        ),
      };
    case REMOVE_BLUE_TEAM_PLAYER:
      return {
        ...state,
        blueTeam: state.blueTeam.filter(
          (player) => player.displayName !== action.payload,
        ),
      };
    default:
      return state;
  }
}
