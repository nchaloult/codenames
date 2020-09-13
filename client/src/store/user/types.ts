// Users are different from Players (in lobby/types.ts): the User is the current
// client, while Players are all other clients that are participating in the
// same game, or are in the same game lobby.
export interface UserState {
  id: string;
  displayName: string;
  isSettingDisplayName: boolean;
  isOnRedTeam: boolean;
}

export const SET_USER_ID = 'SET_USER_ID';
interface SetIDAction {
  type: typeof SET_USER_ID;
  payload: string;
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
export const SET_IS_ON_RED_TEAM = 'SET_IS_ON_RED_TEAM';
interface SetIsOnRedTeamAction {
  type: typeof SET_IS_ON_RED_TEAM;
  payload: boolean;
}
export type UserActionTypes =
  | SetIDAction
  | SetDisplayNameAction
  | SetIsSettingDisplayNameAction
  | SetIsOnRedTeamAction;
