import React, { useState } from 'react';
import { connect, ConnectedProps } from 'react-redux';
import { RootState } from '../store';
import { setDisplayName, setIsSettingDisplayName } from '../store/user/actions';
import {
  constructAndSendEvent,
  EventKind,
  isResponseOK,
  getResponseErr,
} from '../realtime/ws';

// Redux business.

const mapState = (state: RootState) => ({
  displayName: state.user.displayName,
  isSettingDisplayName: state.user.isSettingDisplayName,
  socket: state.websocket.socket,
});
const mapDispatch = {
  setDisplayName: (displayName: string) => setDisplayName(displayName),
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

    // Send a "change display name" event to the server.
    constructAndSendEvent(
      props.socket,
      EventKind.ChangeDisplayName,
      newDisplayName,
    );
    // Listen for an acknowledgement from the server.
    props.socket.onmessage = (event) => {
      const eventResponse = JSON.parse(event.data);
      if (!isResponseOK(eventResponse)) {
        const err = getResponseErr(eventResponse);
        if (err) {
          // TODO: better error handling
          console.error(err);
        }
        return;
      }

      // Response from the server looked good. Commit the display name change
      // client-side.
      props.setDisplayName(newDisplayName);
      props.setIsSettingDisplayName(false);
    };
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
