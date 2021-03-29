# HttpServer

## Description

- This is a simple Http server written in Go

## Settings

- Default port is 8080
- You can change the port by setting the environment variable: ```TCP_PORT=XXXX (ie. TCP_PORT=9090```)

## Usage

```bash
kubectl run frontend --image=am8850/httpserver --env=TCP_PORT=9090 --restart=Never --port=9090 --expose
```
