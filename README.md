To start all containers (development):
`docker-compose -f docker-compose.dev.yaml up -d --build`

To build just the BE + DB:
`docker-compose -f docker-compose.dev.yaml up -d --build backend`

To stop all containers:
`docker-compose down`

To turn on file watching:
`docker-compose alpha watch`
