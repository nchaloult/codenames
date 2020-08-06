# codenames

[![server-tests](https://github.com/nchaloult/codenames/workflows/server-tests/badge.svg)](https://github.com/nchaloult/codenames/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/nchaloult/codenames)](https://goreportcard.com/report/github.com/nchaloult/codenames)

Bringing a popular board game online that you can play in real time with your friends

## Developing Locally

Spin up both the server and the client with the default target in `Makefile`:

    $ make

If you've just cloned this repo, then you'll first need to install dependencies that the front-end needs:

    $ cd client && yarn

## Deploying to Production

The following steps are reminders for myself more than anything. If you'd like to deploy your own version of this project, though, it shouldn't be difficult to adapt these steps, or make changes in `Makefile`, to better fit your situation.

### Server

The server is "containerized," so you may deploy it through whichever cloud service provider that you're comfortable with, pretty much. I've chosen Heroku because of its great free-tier plan.

If you'd like, you may test out how your changes to the server behave in a production environment by running:

    $ make testb

If everything looks good, deploy your changes to Heroku. This step assumes that you've already installed the Heroku CLI on your machine, logged into your Heroku account through the CLI, and created a Heroku app named "codenames".

    $ make deployb

To check on the status of things, you can view your production deployment's log messages in real time:

    $ make deploylogs

### Client

I've deployed the front-end web application with Firebase Hosting. The steps below assume that you've already installed the Firebase CLI on your machine, logged into your Google account through the CLI, created a new Firebase project, and tied this repo to it with `firebase init`.

Push your changes/additions up to Firebase by running:

    $ make deployf
