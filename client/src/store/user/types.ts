export interface UserState {
  displayName: string;
  isSettingDisplayName: boolean;
}

export const SET_DISPLAY_NAME = 'SET_DISPLAY_NAME';
interface SetDisplayNameAction {
  type: typeof SET_DISPLAY_NAME;
  payload: string;
}
export const SET_IS_SETTING_DISPLAY_NAME = 'SET_IS_SETTING_DISPLAY_NAME';
interface SetIsSettingDisplayNameAction {
  type: typeof SET_IS_SETTING_DISPLAY_NAME;
  payload: boolean;
}
export type UserActionTypes =
  | SetDisplayNameAction
  | SetIsSettingDisplayNameAction;
