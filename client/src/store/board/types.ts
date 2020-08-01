export interface Card {
  word: string;
  isRevealed: boolean;
  classification: number;
}

export interface BoardState {
  board: Card[];
}

export const SET_BOARD = 'SET_BOARD';
interface SetBoardAction {
  type: typeof SET_BOARD;
  payload: Card[];
}
export type BoardActionTypes = SetBoardAction;
