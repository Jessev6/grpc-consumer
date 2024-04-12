# Project: grpc-consumer

> This project is a part of the event-store project which can be found at [TODO: configure repo with submodules]()

## Development
> This project was fully setup for development in Visual Studio Code. Building and debugging is supported using the VSCode IDE.

#### Requirements

##### Install go
[Official go docs](https://golang.org/doc/install#install)

##### Get the protobuf tools for go
Run following command to get the tools
```
go get -u github.com/golang/protobuf/protoc-gen-go
```

#### Generate the protobuf files

##### VSCode

Run task: "Generate go protos"

##### CLI

Run commands from project root:
```
protoc \
  --go_out=. \
  --go-grpc_out=. \
  protobuf/*.proto
```

#### Debugging

##### VSCode

Launch debug configuration (F5)


#### Build docker image
##### VSCode

Run build task

##### CLI
Run `docker build -t grpc-consumer -f build/docker/Dockerfile .`

#### Testing
> grpc-ui does not support bi-directional streaming so it cannot be used to test this application.

For testing a small example client was put in place and can be found at: [TODO: link to example client repo]()

## Configuration

| Key  | Setting  | Default |
|---|---| ---|
| LOG_LEVEL  | Sets the global logging level for the application, defaults to 3 [INFO]  | 3 |
| PORT | The port to expose the application on | 3000 |
| REDIS_ADDR | The address of the redis server to connect to | localhost:6379 |
| REDIS_PASSWORD | The password to use to connect to redis |  |
| REDIS_TLS | Can be configured to enable tls to secure connections to redis. This setting should be set to `enabled` to use tls | disabled |
| TLS | Can be configured to enable tls to secure connections to this service. This setting should be set to `enabled` to use tls | disabled |
| TLS_CERT_PATH | Path to certificate file |  |
| TLS_KEY_PATH | Path to certificate key file |  |
