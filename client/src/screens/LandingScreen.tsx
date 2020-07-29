import React, { useState } from 'react';
import LandingScreenLinks from '../components/LandingScreenLinks';

const LandingScreen: React.FC = () => {
  const [gameID, setGameID] = useState('');

  const handleGameIDFormSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    alert(`Temporary! Sending gameID: ${gameID} logic goes here soon :)`);
  };

  return (
    <div className="container centered-container">
      <div className="card card-sm">
        <h1>Codenames</h1>
        <p>
          Play the{' '}
          <a
            href="https://www.youtube.com/watch?v=zQVHkl8oQEU"
            target="_blank"
            rel="noopener noreferrer">
            popular board game
          </a>{' '}
          online with your friends.
        </p>
        <div id="form-sm-wrapper">
          <form className="form-sm" onSubmit={(e) => handleGameIDFormSubmit(e)}>
            <input
              type="text"
              value={gameID}
              placeholder="game-id"
              onChange={(e) => setGameID(e.target.value)}
            />
            <button type="submit">Play</button>
          </form>
        </div>
        <p className="subtext">
          Enter a game ID to join an existing game or to create a new one.
        </p>
      </div>
      <LandingScreenLinks />
    </div>
  );
};

export default LandingScreen;
