export interface GameState {
  id: string;
}

export const SET_GAME_ID = 'SET_GAME_ID';
interface SetGameIDAction {
  type: typeof SET_GAME_ID;
  payload: string;
}
export type GameActionTypes = SetGameIDAction;
