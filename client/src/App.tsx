import React, { useState } from 'react';

const App: React.FC = () => {
  const [gameID, setGameID] = useState('game-id');

  const handleGameIDFormSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    alert(`Temporary! Sending gameID: ${gameID} logic goes here soon :)`);
  };

  return (
    <div className="container centered-container">
      <div className="card card-sm">
        <h1>Codenames</h1>
        <p>
          Play the popular{' '}
          <a href="https://www.youtube.com/watch?v=zQVHkl8oQEU">Codenames</a>{' '}
          board game online with your friends.
        </p>
        <p>Enter a game ID to join an existing game or to create a new one.</p>
        <form className="form-sm" onSubmit={(e) => handleGameIDFormSubmit(e)}>
          <input
            type="text"
            value={gameID}
            onChange={(e) => setGameID(e.target.value)}
          />
          <button type="submit">Play</button>
        </form>
      </div>
    </div>
  );
};

export default App;
