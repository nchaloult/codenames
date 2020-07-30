import React, { useState } from 'react';
import { RouteComponentProps } from 'react-router-dom';
import LandingScreenLinks from '../components/LandingScreenLinks';

type Props = RouteComponentProps<any>;

const LandingScreen: React.FC<Props> = (props: Props) => {
  const [gameID, setGameID] = useState('');

  const handleGameIDFormSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    // Just redirect users to the /:gameID route. Submitting this form is no
    // different than navigating to the /:gameID route manually.
    e.preventDefault();
    props.history.push(`/${gameID}`);
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
              autoFocus
              type="text"
              value={gameID}
              placeholder="game-id"
              onChange={(e) => setGameID(e.target.value)}
            />
            <button className="primary-btn" type="submit">
              Play
            </button>
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
