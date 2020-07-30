import React from 'react';
import { connect, ConnectedProps } from 'react-redux';
import LandingScreenLinks from '../components/LandingScreenLinks';
import { RootState } from '../store';
import setGameID from '../store/game/actions';

// Redux business.

const mapState = (state: RootState) => ({
  gameID: state.game.id,
});
const mapDispatch = {
  setGameID: (id: string) => setGameID(id),
};
const connector = connect(mapState, mapDispatch);
type PropsFromRedux = ConnectedProps<typeof connector>;

// Component.

const LandingScreen: React.FC<PropsFromRedux> = (props: PropsFromRedux) => {
  const handleGameIDFormSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    alert(`Temporary! Sending gameID: ${props.gameID} logic goes here soon :)`);
  };

  return (
    <div className="container centered-container">
      <div className="card card-sm">
        <h1>Codenames</h1>
        <p>
          Play the{' '}
          <a
            href="https://www.youtube.com/watch?v=zQVHkl8oQEU"
            target="_blank"
            rel="noopener noreferrer">
            popular board game
          </a>{' '}
          online with your friends.
        </p>
        <div id="form-sm-wrapper">
          <form className="form-sm" onSubmit={(e) => handleGameIDFormSubmit(e)}>
            <input
              autoFocus
              type="text"
              value={props.gameID}
              placeholder="game-id"
              onChange={(e) => props.setGameID(e.target.value)}
            />
            <button type="submit">Play</button>
          </form>
        </div>
        <p className="subtext">
          Enter a game ID to join an existing game or to create a new one.
        </p>
      </div>
      <LandingScreenLinks />
    </div>
  );
};

export default connector(LandingScreen);
