export interface GameState {
  id: string;
  isCreated: boolean;
  isJoined: boolean;
}

export const SET_GAME_ID = 'SET_GAME_ID';
interface SetGameIDAction {
  type: typeof SET_GAME_ID;
  payload: string;
}
export const SET_IS_CREATED = 'SET_IS_CREATED';
interface SetIsCreatedAction {
  type: typeof SET_IS_CREATED;
  payload: boolean;
}
export const SET_IS_JOINED = 'SET_IS_JOINED';
interface SetIsJoinedAction {
  type: typeof SET_IS_JOINED;
  payload: boolean;
}
export type GameActionTypes =
  | SetGameIDAction
  | SetIsCreatedAction
  | SetIsJoinedAction;
