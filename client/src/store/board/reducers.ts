import { BoardState, BoardActionTypes, SET_BOARD } from './types';

const initialState: BoardState = {
  board: [
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
    // {
    //   word: 'foo',
    //   isRevealed: false,
    //   classification: 0,
    // },
  ],
};

export default function boardReducer(
  state = initialState,
  action: BoardActionTypes,
): BoardState {
  switch (action.type) {
    case SET_BOARD:
      return {
        ...state,
        board: action.payload,
      };
    default:
      return state;
  }
}
