import React, { useState, useEffect } from 'react';
import { connect, ConnectedProps } from 'react-redux';
import { RouteComponentProps } from 'react-router-dom';
import { RootState } from '../store';
import { setIsCreated, setGameID, setIsJoined } from '../store/game/actions';
import Loading from '../components/Loading';
import CreateGameScreen from './CreateGameScreen';
import BoardScreen from './BoardScreen';
import JoinGameScreen from './JoinGameScreen';
import { setUserID } from '../store/user/actions';
import setSocket from '../store/websocket/actions';
import { establishWSConnection } from '../realtime/ws';

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
  setUserID: (id: string) => setUserID(id),
  setGameID: (id: string) => setGameID(id),
  setIsCreated: (isCreated: boolean) => setIsCreated(isCreated),
  setIsJoined: (isJoined: boolean) => setIsJoined(isJoined),
  setSocket: (socket: WebSocket) => setSocket(socket),
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

    // Establish a Websocket connection with the server.
    const socket = establishWSConnection(gameIDFromURL);
    props.setSocket(socket);

    // For now, just stub this out & pretend that the server told us that a game
    // with the provided ID doesn't exist, therefore we'll be creating a new
    // one.
    props.setIsCreated(false);
    props.setIsJoined(false);

    // Once a connection is established, the server will create a new game with
    // the provided gameID if one doens't already exist. Once a game has either
    // been found or a new one has been created, the server will then add a new
    // Player to that game. That Player will have its own unique ID, which the
    // client needs to store. Listen for that new Player ID here, as well.

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
    return <JoinGameScreen />;
  }
  return <CreateGameScreen />;
};

export default connector(GameScreen);
