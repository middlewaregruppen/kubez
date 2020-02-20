# Welcome to Dr. Kubez

Dr. Kubez is a tool to test, diagnose and load k8s clusters.

To run it in your cluster

``` 
kubectl apply -f kubez.yaml
```

##  Things you can do with it today

### Load
- Load the API Server with gazillions of calls to fill up the etcd database.
- Max out the CPU
- Allocate RAM-memory 

### Diagnostics
- View information about CGroups.
- View HTTP Headers that comes in to the container.
- TCP/IP Connection test

### Test:
- Create services that mock endpoints and introduces errors and delays in them.


## Thing we are about to create:
- Hijack k8s services and introduce random errors and delays to them. 
- DNS Lookup from inside the cluster.




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





