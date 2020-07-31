import React from 'react';
import './CreateGameScreen.css';
import { connect, ConnectedProps } from 'react-redux';
import { RootState } from '../store';
import DisplayName from '../components/DisplayName';

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
      <div className="even-columns">
        <div>
          <h1>Create a New Game</h1>
          <h3>{props.gameID.toUpperCase()}</h3>
          <DisplayName />
          <button className="secondary-btn" type="button">
            Join Red Team
          </button>
          <button className="secondary-btn" type="button">
            Join Blue Team
          </button>
          <button className="secondary-btn" type="button">
            Change Some Other Setting
          </button>
          <button className="primary-btn" id="create-game-btn" type="button">
            Create Game
          </button>
        </div>
        <div className="indented-col">
          <h2>Red Team</h2>
        </div>
        <div className="indented-col">
          <h2>Blue Team</h2>
        </div>
      </div>
    </div>
  </div>
);

export default connector(CreateGameScreen);
