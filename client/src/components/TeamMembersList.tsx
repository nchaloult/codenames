import React from 'react';
import './TeamMembersList.css';

const RED_TEAM_TITLE = 'Red Team';
const BLUE_TEAM_TITLE = 'Blue Team';

interface Props {
  isRedTeam: boolean;
}

const TeamMembersList: React.FC<Props> = (props: Props) => (
  <div id="indented-col">
    <h2>{props.isRedTeam ? RED_TEAM_TITLE : BLUE_TEAM_TITLE}</h2>
  </div>
);

export default TeamMembersList;
