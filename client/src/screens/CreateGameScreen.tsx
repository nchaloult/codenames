import React from 'react';
import './CreateGameScreen.css';
import { connect, ConnectedProps } from 'react-redux';
import { RootState } from '../store';
import DisplayName from '../components/DisplayName';
import TeamMembersList from '../components/TeamMembersList';
import SetDisplayNameScreen from './SetDisplayNameScreen';
import { Player } from '../store/lobby/types';
import {
  addRedTeamPlayer,
  addBlueTeamPlayer,
  removeRedTeamPlayer,
  removeBlueTeamPlayer,
} from '../store/lobby/actions';
import { constructAndSendEvent, EventKind } from '../realtime/ws';

// Redux business.

const mapState = (state: RootState) => ({
  playerID: state.user.id,
  gameID: state.game.id,
  displayName: state.user.displayName,
  isSettingDisplayName: state.user.isSettingDisplayName,
  socket: state.websocket.socket,
});
const mapDispatch = {
  addRedTeamPlayer: (player: Player) => addRedTeamPlayer(player),
  addBlueTeamPlayer: (player: Player) => addBlueTeamPlayer(player),
  removeRedTeamPlayer: (id: string) => removeRedTeamPlayer(id),
  removeBlueTeamPlayer: (id: string) => removeBlueTeamPlayer(id),
};
const connector = connect(mapState, mapDispatch);
type PropsFromRedux = ConnectedProps<typeof connector>;

// Component.

const CreateGameScreen: React.FC<PropsFromRedux> = (props: PropsFromRedux) => {
  const handleJoinTeamButtonPressed = (isJoiningRedTeam: boolean) => {
    if (!props.socket) {
      // TODO: redirect to some /error route or something to indicate a
      // connection issue.
      return;
    }

    constructAndSendEvent(props.socket, EventKind.changeTeam, isJoiningRedTeam);
    if (isJoiningRedTeam) {
      props.addRedTeamPlayer({
        id: props.playerID,
        displayName: props.displayName,
      });
      props.removeBlueTeamPlayer(props.playerID);
    } else {
      props.addBlueTeamPlayer({
        id: props.playerID,
        displayName: props.displayName,
      });
      props.removeRedTeamPlayer(props.playerID);
    }
  };

  if (props.isSettingDisplayName) {
    return <SetDisplayNameScreen />;
  }
  return (
    <div className="container centered-container">
      <div className="card">
        <div className="even-columns">
          <div>
            <h1>Create a New Game</h1>
            <h3>{props.gameID.toUpperCase()}</h3>
            <DisplayName />
            <button
              className="secondary-btn"
              type="button"
              onClick={() => handleJoinTeamButtonPressed(true)}>
              Join Red Team
            </button>
            <button
              className="secondary-btn"
              type="button"
              onClick={() => handleJoinTeamButtonPressed(false)}>
              Join Blue Team
            </button>
            <button className="secondary-btn" type="button">
              Change Some Other Setting
            </button>
            <button className="primary-btn" id="create-game-btn" type="button">
              Create Game
            </button>
          </div>
          <TeamMembersList isRedTeam={true} />
          <TeamMembersList isRedTeam={false} />
        </div>
      </div>
    </div>
  );
};

export default connector(CreateGameScreen);
