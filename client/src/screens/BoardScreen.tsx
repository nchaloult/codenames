import React from 'react';
import './BoardScreen.css';
import Board from '../components/Board';
import BoardInfo from '../components/BoardInfo';

const BoardScreen: React.FC = () => (
  <div className="container centered-container">
    <div id="board-screen-card" className="card">
      <div id="small-col">
        <BoardInfo />
      </div>
      <div id="big-col">
        <Board />
      </div>
    </div>
  </div>
);

export default BoardScreen;
