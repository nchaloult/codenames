import { UserState, UserActionTypes, SET_DISPLAY_NAME } from './types';

const initialState: UserState = {
  displayName: '',
};

export default function userReducer(
  state = initialState,
  action: UserActionTypes,
): UserState {
  switch (action.type) {
    case SET_DISPLAY_NAME:
      return {
        ...state,
        displayName: action.payload,
      };
    default:
      return state;
  }
}
