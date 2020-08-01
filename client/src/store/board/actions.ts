import { BoardActionTypes, Card, SET_BOARD } from './types';

// Action creator for the SET_BOARD action type.
export default function setBoard(board: Card[]): BoardActionTypes {
  return {
    type: SET_BOARD,
    payload: board,
  };
}
