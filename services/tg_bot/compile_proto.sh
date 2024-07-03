#!/bin/sh

python -m grpc_tools.protoc -I./app/api/proto --python_out=./app/api/proto --grpc_python_out=./app/api/proto ./app/api/proto/tn.proto

echo "Proto compilation completed"