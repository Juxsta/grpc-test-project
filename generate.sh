#!/bin/bash

cd greet/greetpb

docker run -v $(pwd):/defs protoc -f greet.proto -l go
mv gen/pb-go/greet.pb.go .
rm -rf gen
