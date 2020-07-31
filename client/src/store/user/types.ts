export interface UserState {
  displayName: string;
}

export const SET_DISPLAY_NAME = 'SET_DISPLAY_NAME ';
interface SetDisplayNameAction {
  type: typeof SET_DISPLAY_NAME;
  payload: string;
}
export type UserActionTypes = SetDisplayNameAction;
