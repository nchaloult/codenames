import React, { useState } from 'react';
import { connect, ConnectedProps } from 'react-redux';
import { RootState } from '../store';
import { setDisplayName, setIsSettingDisplayName } from '../store/user/actions';

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

    // Send a "change display name" event to the server.
    const changeDisplayNameEvent = {
      kind: 'changeDisplayName',
      body: newDisplayName,
    };
    props.socket?.send(JSON.stringify(changeDisplayNameEvent));

    props.setDisplayName(newDisplayName);
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
