import React from 'react';
import { connect, ConnectedProps } from 'react-redux';
import { RootState } from '../store';
import DisplayName from '../components/DisplayName';
import TeamMembersList from '../components/TeamMembersList';
import SetDisplayNameScreen from './SetDisplayNameScreen';

// Redux business.

const mapState = (state: RootState) => ({
  gameID: state.game.id,
  isSettingDisplayName: state.user.isSettingDisplayName,
});
const connector = connect(mapState);
type PropsFromRedux = ConnectedProps<typeof connector>;

// Component.

const JoinGameScreen: React.FC<PropsFromRedux> = (props: PropsFromRedux) => {
  if (props.isSettingDisplayName) {
    return <SetDisplayNameScreen />;
  }
  return (
    <div className="container centered-container">
      <div className="card">
        <div className="even-columns">
          <div>
            <h1>Join an Ongoing Game</h1>
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
              Join Game
            </button>
          </div>
          <TeamMembersList title="Red Team" />
          <TeamMembersList title="Blue Team" />
        </div>
      </div>
    </div>
  );
};

export default connector(JoinGameScreen);
