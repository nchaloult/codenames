import React from 'react';
import './Board.css';
import { connect, ConnectedProps } from 'react-redux';
import { RootState } from '../store';
import { Card } from '../store/board/types';
import CardComponent from './Card';

// Redux business.

const mapState = (state: RootState) => ({
  board: state.board.board,
});
const connector = connect(mapState);
type PropsFromRedux = ConnectedProps<typeof connector>;

// Component.

const Board: React.FC<PropsFromRedux> = (props: PropsFromRedux) => (
  <div id="board-container">
    {props.board.map((card: Card, index) => (
      <CardComponent key={`${card.word}+${index}`} index={index} />
    ))}
  </div>
);

export default connector(Board);
