#!/bin/bash

protoc C:\Users\ericj\Documents\GitHub\grpc-test-project\greetsum\greetsumpb\greetsum.proto -I $env:GOPATH/src -I C:\Users\ericj\Documents\GitHub\grpc-test-project\greetsum\greetsumpb -I $env:GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:.
