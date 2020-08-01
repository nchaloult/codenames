import React from 'react';
import './BoardInfo.css';

// Redux business.
// TODO: wire up corresponding fields in the global store (like game score, for
// instance) once they're added.

// Component.

const BoardInfo: React.FC = () => (
  <div id="board-info-container">
    <div>
      <h1>
        <span id="red-team-score">5</span> - <span id="blue-team-score">4</span>
      </h1>
      <h3>BLUE TEAM'S TURN</h3>
    </div>
    <div>
      <h1>Clue: Foo 2</h1>
    </div>
    <div>
      <button className="primary-btn" type="button">
        Vote to End Turn
      </button>
    </div>
  </div>
);

export default BoardInfo;
