FROM golang:1.16-alpine AS build

WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/tcpserver

FROM scratch
COPY --from=build /bin/tcpserver /bin/tcpserver
ENTRYPOINT ["/bin/tcpserver"]
