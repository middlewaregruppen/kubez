# Development

## Requirements
- GNU Make
- Go
- Docker

## Building dr. kubez from source
```sh
make frontend
make docker_build
```

## Start the server locally
```sh
go run cmd/kubez/main.go
```

or using Docker

```sh
docker run -p3000:3000 -m50m --cpus=0.5 middlewaregruppen/kubez:latest
```

## Start the UI
`npm run serve --prefix web/frontend`

Access the ui from `http://localhost:8080`

Kubez will listen to port 3000. Vue listen on 8080. Vue will proxy requests to /kubez to kubez.
