[![Backend CI/CD](https://github.com/gwkline/full-stack-skeleton/actions/workflows/backend.yml/badge.svg)](https://github.com/gwkline/full-stack-skeleton/actions/workflows/backend.yml)
[![Frontend CI/CD](https://github.com/gwkline/full-stack-skeleton/actions/workflows/frontend.yml/badge.svg)](https://github.com/gwkline/full-stack-skeleton/actions/workflows/frontend.yml)
[![Backend Coverage](https://codecov.io/gh/gwkline/full-stack-skeleton/branch/main/graph/badge.svg?token=FQGXXYYJT1)](https://codecov.io/gh/gwkline/full-stack-skeleton)

To start all containers (development):
`docker-compose -f docker-compose.dev.yaml up -d --build`

To build just the BE + DB:
`docker-compose -f docker-compose.dev.yaml up -d --build backend`

To stop all containers:
`docker-compose down`

To turn on file watching:
`docker-compose alpha watch`
