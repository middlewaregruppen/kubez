FROM golang:alpine AS go-build
COPY . /build
WORKDIR /build/
RUN apk add --no-cache --update git make ca-certificates \
&&  make

FROM node:10 AS npm-build
COPY web/frontend/ /build
WORKDIR /build/
RUN npm install \
&&  NODE_ENV=production npm run build  

FROM scratch
LABEL maintaner="@middlewaregruppen (github.com/middlewaregruppen)"
COPY --from=go-build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=go-build /build/bin/kubez-linux-amd64 /go/bin/kubez
COPY --from=npm-build /build/dist /web/frontend/dist
ENTRYPOINT ["/go/bin/kubez"]