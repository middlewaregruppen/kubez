# Welcome to Dr. Kubez

## Building dr. kubez
```
make frontend
make docker_build
```

## Local development
#### Start the server 
`go run cmd/kubez/main.go`
 
 or 

`docker run -p3000:3000 -m50m --cpus=0.5 docker.pkg.github.com/middlewaregruppen/kubez:latest`

#### Start vue 

`npm run serve --prefix web/frontend`

Access the ui from `http://localhost:8080`

Kubez will listen to port 3000. Vue listen on 8080. Vue will proxy requests to /kubez to kubez.





