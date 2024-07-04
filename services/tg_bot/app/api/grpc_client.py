import grpc
from app.api.proto import tn_pb2_grpc, tn_pb2
from app.config.config import config

from google.protobuf.empty_pb2 import Empty


class GRPCClient:
    def __init__(self, host=config['grpc_client']['address'], port=config['grpc_client']['port']):
        self.channel = grpc.insecure_channel(f'{host}:{port}')
        self.stub = tn_pb2_grpc.TNStub(self.channel)

    def get_callbacks_paginated(self, limit: int, offset: int):
        request = tn_pb2.RequestGetCallbacksPaginated(limit=limit, offset=offset)
        return self.stub.GetCallbacksPaginated(request)

    def get_callbacks_quantity(self):
        request = Empty()
        return self.stub.GetCallbacksQuantity(request)

    def delete_callback(self, id: int):
        request = tn_pb2.RequestDeleteCallback(id=id)
        return self.stub.DeleteCallback(request)
