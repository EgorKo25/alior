#!/bin/sh

python -m grpc_tools.protoc -I./app/api/proto --python_out=./app/api/proto --grpc_python_out=./app/api/proto ./app/api/proto/tn.proto
sed -i 's/import tn_pb2 as tn__pb2/from . import tn_pb2 as tn__pb2/' ./app/api/proto/tn_pb2_grpc.py

echo "Proto compilation completed"