import {
  UserActionTypes,
  SET_DISPLAY_NAME,
  SET_IS_SETTING_DISPLAY_NAME,
} from './types';

// Action creator for the SET_DISPLAY_NAME action type.
export function setDisplayName(displayName: string): UserActionTypes {
  return {
    type: SET_DISPLAY_NAME,
    payload: displayName,
  };
}

// Action creator for the SET_IS_SETTING_DISPLAY_NAME action type.
export function setIsSettingDisplayName(
  isSettingDisplayName: boolean,
): UserActionTypes {
  return {
    type: SET_IS_SETTING_DISPLAY_NAME,
    payload: isSettingDisplayName,
  };
}
