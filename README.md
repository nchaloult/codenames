# codenames

Bringing a popular board game online that you can play in real time with your friends

## Launching with Docker

Build an image that mirrors our production environment, and spin up a container from that image.

    $ docker build -t codenames .
    $ docker run --rm --name codenames -p 3000:6969 -d codenames

Now, test things out by either visiting `localhost:3000` in a browser, or hitting HTTP API endpoints at `localhost:3000/api`. For instance:

    $ curl localhost:3000/api/hello

When you're done, spin everything down.

    $ docker stop codenames

## Deploying to Production on Heroku

Build a new image with your most recent changes/additions.

    $ docker build -t codenames .

Deploy your changes to Heroku. These steps assume that you've already installed the Heroku CLI on your machine, logged into your Heroku account through the CLI, and created a Heroku app named "codenames".

    $ heroku container:push web -a codenames
    $ heroku container:release web -a codenames

To check on the status of things, you can view your production deployment's logs.

    $ heroku logs -t -a codenames
