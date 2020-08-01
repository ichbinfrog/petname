ARG VERSION="v0.1.0"
ARG GO_VERSION="1.14.6"

# Builder image
FROM golang:${GO_VERSION} as petname-builder

# Sets workdir
WORKDIR /go/src/app
ADD . /go/src/app

# Installs dependencies
RUN go get -d -v ./...

# Compiles go app
RUN go build -o /go/bin/app


# Distroless execution image
FROM gcr.io/distroless/base
COPY --from=petname-builder /go/bin/app /
COPY .seed.yaml $HOME/.seed.yaml

ENTRYPOINT ["/app"]
