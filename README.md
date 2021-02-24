# Go Wasm Demo

## Build

```shell script
GOOS=js GOARCH=wasm go build -o main.wasm
```

## Run

```shell script
go run cmd/server/main.go
open http://localhost:8080
```
