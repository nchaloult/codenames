// This component exists in case I want to let users change their display names
// after they've already provided one before they join a game. Perhaps a
// "change" button could exist, and when clicked, a text field would be
// presented.

import React from 'react';
import './DisplayName.css';
import { connect, ConnectedProps } from 'react-redux';
import { RootState } from '../store';

// Redux business.
const mapState = (state: RootState) => ({
  displayName: state.user.displayName,
});
const connector = connect(mapState);
type PropsFromRedux = ConnectedProps<typeof connector>;

// Component.

const DisplayName: React.FC<PropsFromRedux> = (props: PropsFromRedux) => {
  return (
    <p>
      Display name: <span id="emph-display-name">{props.displayName}</span>
    </p>
  );
};

export default connector(DisplayName);
