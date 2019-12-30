# docker build -t ericsgagnon/debug-webserver:latest .

ARG GO_VERSION=1.13

FROM golang:${GO_VERSION} as build

WORKDIR /go/src/debug-webserver

COPY . .

RUN go mod download
RUN go generate
RUN go build

EXPOSE 18080

ENTRYPOINT [ "./debug-webserver" ]