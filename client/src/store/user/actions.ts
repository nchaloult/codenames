import { UserActionTypes, SET_DISPLAY_NAME } from './types';

// Action creator for the SET_DISPLAY_NAME action type.
export default function setDisplayName(displayName: string): UserActionTypes {
  return {
    type: SET_DISPLAY_NAME,
    payload: displayName,
  };
}
