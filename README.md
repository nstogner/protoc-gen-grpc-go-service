# Golang Service Struct Generator (gRPC)

Use this protoc plugin to generate a go struct that satisfies the corresponding generated gRPC interface.

```sh
# Install
go get github.com/nstogner/protoc-gen-grpc-goservice

# How example/grpcd/service.go was generated...
cd $GOPATH/github.com/nstogner/protoc-gen-grpc-goservice/example

protoc --plugin=protoc-gen-grpc-goservice=$GOPATH/bin/protoc-gen-grpc-goservice --grpc-goservice_out=./grpcd example.proto
```
