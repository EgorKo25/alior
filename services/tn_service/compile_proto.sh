#!/bin/sh

python -m grpc_tools.protoc -I../../docs/api/proto --python_out=./app/api/proto --grpc_python_out=./app/api/proto ../../docs/api/proto/tn.proto

echo "Proto compilation completed"