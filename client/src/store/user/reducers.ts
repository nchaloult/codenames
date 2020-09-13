import {
  UserState,
  UserActionTypes,
  SET_DISPLAY_NAME,
  SET_IS_SETTING_DISPLAY_NAME,
  SET_USER_ID,
  SET_IS_ON_RED_TEAM,
} from './types';

const initialState: UserState = {
  id: '',
  displayName: '',
  isSettingDisplayName: true,
  isOnRedTeam: true,
};

export default function userReducer(
  state = initialState,
  action: UserActionTypes,
): UserState {
  switch (action.type) {
    case SET_USER_ID:
      return {
        ...state,
        id: action.payload,
      };
    case SET_DISPLAY_NAME:
      return {
        ...state,
        displayName: action.payload,
      };
    case SET_IS_SETTING_DISPLAY_NAME:
      return {
        ...state,
        isSettingDisplayName: action.payload,
      };
    case SET_IS_ON_RED_TEAM:
      return {
        ...state,
        isOnRedTeam: action.payload,
      };
    default:
      return state;
  }
}
