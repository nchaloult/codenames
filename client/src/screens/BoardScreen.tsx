import React from 'react';
import './BoardScreen.css';
import Board from '../components/Board';

const BoardScreen: React.FC = () => (
  <div className="container centered-container">
    <div id="board-screen-card" className="card">
      <div id="small-col">
        <button className="secondary-btn" type="button">
          Placeholder
        </button>
      </div>
      <div id="big-col">
        <Board />
      </div>
    </div>
  </div>
);

export default BoardScreen;
