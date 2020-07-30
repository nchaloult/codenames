import React, { useEffect } from 'react';
import { connect, ConnectedProps } from 'react-redux';
import { RouteComponentProps } from 'react-router-dom';
import { RootState } from '../store';

// Redux business.

const mapState = (state: RootState) => ({
  gameID: state.game.id,
});
const connector = connect(mapState);
type PropsFromRedux = ConnectedProps<typeof connector>;

// Component.

type Props = PropsFromRedux & RouteComponentProps<any>;

const CreateGameScreen: React.FC<Props> = (props: Props) => {
  useEffect(() => {
    // If the game ID is blank and we've arrived on this page somehow, then
    // redirect users back to the landing page.
    if (props.gameID === '') {
      props.history.push('/');
    }
  }, []);

  return (
    <div className="container centered-container">
      <div className="card">
        <h1>Create a New Game</h1>
        <h3>{props.gameID.toUpperCase()}</h3>
      </div>
    </div>
  );
};

export default connector(CreateGameScreen);
