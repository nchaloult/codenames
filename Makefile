dev: # Default target. Spins up back-end and front-end
	make -j 2 runf runb

# Client targets

runf: # Run front-end
	cd client && yarn start

buildf: # Build front-end for production
	cd client && yarn build

# Server targets

runb: # Run back-end
	cd server && go run *.go

buildb: # Build back-end with docker
	cd server && docker build -t codenames .

testb: buildb # Test back-end in a production environment
	docker run --rm --name codenames -p 6969:6969 -d codenames

stoptestb: # Spin down docker container
	docker stop codenames

# Deploy targets

deployf: buildf # Deploy front-end to Firebase
	cd client && firebase deploy

deployb: # Deploy back-end to Heroku
	heroku container:push web -a codenames && \
	heroku container:release web -a codenames

deployboth: # Deploy both the back-end and the front-end
	deployf && deployb

deploylogs: # Check logs of back-end container in production
	heroku logs -t -a codenames
