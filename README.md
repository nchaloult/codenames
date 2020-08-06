# codenames

[![server-tests](https://github.com/nchaloult/codenames/workflows/server-tests/badge.svg)](https://github.com/nchaloult/codenames/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/nchaloult/codenames)](https://goreportcard.com/report/github.com/nchaloult/codenames)

Bringing a popular board game online that you can play in real time with your friends

## Developing Locally

After cloning this repo, get everything ready to go:

    $ cd codenames/client && yarn

Then in one shell, spin up the server:

    $ cd server
    $ go run *.go

Or if you're hungry for speed, you can:

    $ go build && ./codenames

And in another shell, spin up the React app:

    $ cd client
    $ yarn start

## Deploying to Production

The following steps are reminders for myself more than anything. If you'd like to deploy your own version of this project, though, it shouldn't be difficult to adapt these steps to better fit your situation.

### Server

The server is "containerized," so you may deploy it through whichever cloud service provider that you're comfortable with, pretty much. I've chosen Heroku because of it's great free-tier plan.

If you'd like, you may test out how your changes to the server behave in a production environment by running:

    $ cd server
    $ docker build -t codenames .
    $ docker run --rm --name codenames -p 6969:6969 -d codenames

Deploy your changes to Heroku. These steps assume that you've already installed the Heroku CLI on your machine, logged into your Heroku account through the CLI, and created a Heroku app named "codenames".

    $ heroku container:push web -a codenames
    $ heroku container:release web -a codenames

To check on the status of things, you can view your production deployment's logs.

    $ heroku logs -t -a codenames

### Client

I've deployed the front-end web application with Firebase Hosting. The steps below assume that you've already installed the Firebase CLI on your machine, logged into your Google account through the CLI, created a new Firebase project, and tied this repo to it with `firebase init`.

Invoke the `deploy` script that's defined in in `package.json` to push up your changes/additions by running:

    $ cd client
    $ yarn deploy
