import React from 'react';
import './TeamMembersList.css';

interface Props {
  title: string;
}

const TeamMembersList: React.FC<Props> = (props: Props) => (
  <div id="indented-col">
    <h2>{props.title}</h2>
  </div>
);

export default TeamMembersList;
