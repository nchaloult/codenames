import React from 'react';
import './TeamMembersList.css';
import { connect, ConnectedProps } from 'react-redux';
import { RootState } from '../store';

// Redux business.

const mapState = (state: RootState) => ({
  id: state.user.id,
  redTeam: state.lobby.redTeam,
  blueTeam: state.lobby.blueTeam,
});
const connector = connect(mapState);
type PropsFromRedux = ConnectedProps<typeof connector>;

// Component.

const RED_TEAM_TITLE = 'Red Team';
const BLUE_TEAM_TITLE = 'Blue Team';

interface ComponentProps {
  isRedTeam: boolean;
}
type Props = PropsFromRedux & ComponentProps;

const TeamMembersList: React.FC<Props> = (props: Props) => {
  const title = props.isRedTeam ? RED_TEAM_TITLE : BLUE_TEAM_TITLE;
  const teamMemberList = props.isRedTeam ? props.redTeam : props.blueTeam;

  return (
    <div id="indented-col">
      <h2>{title}</h2>
      <div id="team-list">
        {teamMemberList.map((member) => {
          if (member.id === props.id) {
            return (
              <span
                id="highlighted-team-list-item"
                className="team-list-item"
                key={member.id}>
                {member.displayName}
              </span>
            );
          }
          return (
            <span className="team-list-item" key={member.id}>
              {member.displayName}
            </span>
          );
        })}
      </div>
    </div>
  );
};

export default connector(TeamMembersList);
