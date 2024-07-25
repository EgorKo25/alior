#!/bin/sh

python -m grpc_tools.protoc -I./app/api/proto --python_out=./app/api/proto --grpc_python_out=./app/api/proto ./app/api/proto/tn.proto
python -m grpc_tools.protoc -I./app/api/proto --python_out=./app/api/proto --grpc_python_out=./app/api/proto ./app/api/proto/cms.proto

sed -i 's/import cms_pb2 as cms__pb2/from . import cms_pb2 as cms__pb2/' ./app/api/proto/cms_pb2_grpc.py

echo "Proto compilation completed"