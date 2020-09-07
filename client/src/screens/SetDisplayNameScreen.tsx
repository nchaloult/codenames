import React, { useState } from 'react';
import { connect, ConnectedProps } from 'react-redux';
import { RootState } from '../store';
import { setDisplayName, setIsSettingDisplayName } from '../store/user/actions';
import { constructAndSendEvent, EventKind } from '../realtime/ws';
import { changeDisplayName } from '../store/lobby/actions';

// Redux business.

const mapState = (state: RootState) => ({
  id: state.user.id,
  displayName: state.user.displayName,
  isSettingDisplayName: state.user.isSettingDisplayName,
  socket: state.websocket.socket,
});
const mapDispatch = {
  setLobbyDisplayName: (id: string, displayName: string) =>
    changeDisplayName(id, displayName),
  setClientDisplayName: (displayName: string) => setDisplayName(displayName),
  setIsSettingDisplayName: (isSettingDisplayName: boolean) =>
    setIsSettingDisplayName(isSettingDisplayName),
};
const connector = connect(mapState, mapDispatch);
type PropsFromRedux = ConnectedProps<typeof connector>;

// Component.

const SetDisplayNameScreen: React.FC<PropsFromRedux> = (
  props: PropsFromRedux,
) => {
  const [newDisplayName, setNewDisplayName] = useState(props.displayName);

  const handleDisplayNameFormSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!props.socket) {
      // TODO: redirect to some /error route or something to indicate a
      // connection issue.
      return;
    }

    constructAndSendEvent(
      props.socket,
      EventKind.changeDisplayName,
      newDisplayName,
    );
    props.setLobbyDisplayName(props.id, newDisplayName);
    props.setClientDisplayName(newDisplayName);
    props.setIsSettingDisplayName(false);
  };

  return (
    <div className="container centered-container">
      <div className="card card-sm">
        <h1>Set Display Name</h1>
        <p>
          Your display name is like your username. Other players will see this.
        </p>
        <div id="form-sm-wrapper">
          <form
            className="form-sm"
            onSubmit={(e) => handleDisplayNameFormSubmit(e)}>
            <input
              autoFocus
              type="text"
              value={newDisplayName}
              placeholder="Enter a display name"
              onChange={(e) => setNewDisplayName(e.target.value)}
            />
            <button className="primary-btn" type="submit">
              Set
            </button>
          </form>
        </div>
      </div>
    </div>
  );
};

export default connector(SetDisplayNameScreen);
