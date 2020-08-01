import React, { useState, useEffect } from 'react';
import { connect, ConnectedProps } from 'react-redux';
import { RouteComponentProps } from 'react-router-dom';
import { RootState } from '../store';
import { setIsCreated, setGameID, setIsJoined } from '../store/game/actions';
import Loading from '../components/Loading';
import CreateGameScreen from './CreateGameScreen';
import BoardScreen from './BoardScreen';

// In this component, a Websocket connection is established with the server.
// Depending on whether the game with the provided gameID has already been
// created or not, this component then displays either the CreateGameScreen or
// the JoinGameScreen.

// Redux business.

const mapState = (state: RootState) => ({
  isCreated: state.game.isCreated,
  isJoined: state.game.isJoined,
});
const mapDispatch = {
  setGameID: (id: string) => setGameID(id),
  setIsCreated: (isCreated: boolean) => setIsCreated(isCreated),
  setIsJoined: (isJoined: boolean) => setIsJoined(isJoined),
};
const connector = connect(mapState, mapDispatch);
type PropsFromRedux = ConnectedProps<typeof connector>;

// Component.

type Props = PropsFromRedux & RouteComponentProps<any>;

const GameScreen: React.FC<Props> = (props: Props) => {
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const gameIDFromURL: string = props.match.params.gameID;
    props.setGameID(gameIDFromURL);

    // TODO: logic which establishes a Websocket connection with the server goes
    // here.
    //
    // For now, just stub this out & pretend that the server told us that a game
    // with the provided ID doesn't exist, therefore we'll be creating a new
    // one.
    props.setIsCreated(false);
    props.setIsJoined(false);

    setIsLoading(false);
  }, []);

  // Render.

  if (isLoading) {
    return <Loading />;
  }
  if (props.isJoined) {
    return <BoardScreen />;
  }
  if (props.isCreated) {
    return <p>join game screen goes here</p>;
  }
  return <CreateGameScreen />;
};

export default connector(GameScreen);
