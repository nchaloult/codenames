import React, { useState } from 'react';
import { connect, ConnectedProps } from 'react-redux';
import { RootState } from '../store';
import setDisplayName from '../store/user/actions';

// Redux business.

const mapState = (state: RootState) => ({
  displayName: state.user.displayName,
});
const mapDispatch = {
  setDisplayName: (displayName: string) => setDisplayName(displayName),
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
    props.setDisplayName(newDisplayName);
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
