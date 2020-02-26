BINARY=kubez
GOARCH=amd64
VERSION=dev
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GOVERSION=$(shell go version | awk -F\go '{print $$3}' | awk '{print $$1}')
GITHUB_USERNAME=middlewaregruppen
BUILD_DIR=${GOPATH}/src/github.com/${GITHUB_USERNAME}/${BINARY}
PKG_LIST=$$(go list ./... | grep -v /vendor/)
# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH} -X main.GOVERSION=${GOVERSION}"

# Build the project
all: build

test:
	cd ${BUILD_DIR}; \
	go test -v; \
	cd - >/dev/null

fmt:
	cd ${BUILD_DIR}; \
	go fmt ${PKG_LIST} ; \
	cd - >/dev/null

dep:
	go get -v -d ./... ;

frontend:
	npm install --prefix web/frontend
	NODE_ENV=production npm run build --prefix web/frontend
	

linux: dep
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o out/kubez cmd/kubez/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o out/loader cmd/loader/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o out/api-endpoint cmd/api-endpoint/main.go


docker_build: linux
	#docker run --rm -v "${PWD}":/go/src/github.com/${GITHUB_USERNAME}/${BINARY}  -w /go/src/github.com/${GITHUB_USERNAME}/${BINARY}  make fmt test
	docker build -t docker.io/${GITHUB_USERNAME}/${BINARY}:${VERSION} .
	docker tag docker.io/${GITHUB_USERNAME}/${BINARY}:${VERSION}  docker.io/${GITHUB_USERNAME}/${BINARY}:latest

docker_push:
	docker push docker.io/${GITHUB_USERNAME}/${BINARY}:${VERSION}
	docker push docker.io/${GITHUB_USERNAME}/${BINARY}:latest

docker: docker_build docker_push

build: linux frontend docker

clean:
	-rm -rf ${BUILD_DIR}/out/

.PHONY: linux test fmt clean
