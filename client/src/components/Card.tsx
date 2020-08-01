import React from 'react';
import './Card.css';
import { connect, ConnectedProps } from 'react-redux';
import { RootState } from '../store';

// React business.

const mapState = (state: RootState) => ({
  board: state.board.board,
});
const connector = connect(mapState);
type PropsFromRedux = ConnectedProps<typeof connector>;

// Component.

interface ComponentProps {
  index: number;
}
type Props = PropsFromRedux & ComponentProps;

const Card: React.FC<Props> = (props: Props) => {
  const curCard = props.board[props.index];

  return (
    <div id="card-container">
      <span id="card-word">{curCard.word.toUpperCase()}</span>
    </div>
  );
};

export default connector(Card);
