import React from 'react';
import { connect, ConnectedProps } from 'react-redux';
import { RootState } from '../store';

// Redux business.

const mapState = (state: RootState) => ({
  gameID: state.game.id,
});
const connector = connect(mapState);
type PropsFromRedux = ConnectedProps<typeof connector>;

// Component.

const CreateGameScreen: React.FC<PropsFromRedux> = (props: PropsFromRedux) => (
  <div className="container centered-container">
    <div className="card">
      <h1>Create a New Game</h1>
      <h3>{props.gameID.toUpperCase()}</h3>
    </div>
  </div>
);

export default connector(CreateGameScreen);
