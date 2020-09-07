import {
  LobbyState,
  LobbyActionTypes,
  ADD_RED_TEAM_PLAYER,
  ADD_BLUE_TEAM_PLAYER,
  REMOVE_RED_TEAM_PLAYER,
  REMOVE_BLUE_TEAM_PLAYER,
  CHANGE_DISPLAY_NAME,
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
    case CHANGE_DISPLAY_NAME:
      return {
        ...state,
        redTeam: state.redTeam.map((player) => {
          if (player.id === action.payload.id) {
            return {
              id: player.id,
              displayName: action.payload.newName,
            };
          }
          return player;
        }),
        blueTeam: state.blueTeam.map((player) => {
          if (player.id === action.payload.id) {
            return {
              id: player.id,
              displayName: action.payload.newName,
            };
          }
          return player;
        }),
      };
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
        redTeam: state.redTeam.filter((player) => player.id !== action.payload),
      };
    case REMOVE_BLUE_TEAM_PLAYER:
      return {
        ...state,
        blueTeam: state.blueTeam.filter(
          (player) => player.id !== action.payload,
        ),
      };
    default:
      return state;
  }
}
