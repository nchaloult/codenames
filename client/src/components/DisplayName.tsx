// This component exists in case I want to let users change their display names
// after they've already provided one before they join a game. Perhaps a
// "change" button could exist, and when clicked, a text field would be
// presented.

import React from 'react';
import './DisplayName.css';
import { connect, ConnectedProps } from 'react-redux';
import { RootState } from '../store';
import { setIsSettingDisplayName } from '../store/user/actions';

// Redux business.

const mapState = (state: RootState) => ({
  displayName: state.user.displayName,
});
const mapDispatch = {
  setIsSettingDisplayName: (isSettingDisplayName: boolean) =>
    setIsSettingDisplayName(isSettingDisplayName),
};
const connector = connect(mapState, mapDispatch);
type PropsFromRedux = ConnectedProps<typeof connector>;

// Component.

const DisplayName: React.FC<PropsFromRedux> = (props: PropsFromRedux) => {
  return (
    <div id="display-name-container">
      <p>
        Display name: <span id="emph-display-name">{props.displayName}</span>
      </p>
      <button
        className="secondary-btn"
        type="button"
        onClick={() => props.setIsSettingDisplayName(true)}>
        Change Display Name
      </button>
    </div>
  );
};

export default connector(DisplayName);
