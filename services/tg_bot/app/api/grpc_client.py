import grpc
from app.api.proto import tn_pb2_grpc, tn_pb2
from app.config.config import config

class GRPCClient:
    def __init__(self, host=config['grpc_client']['address'], port=config['grpc_client']['port']):
        self.channel = grpc.insecure_channel(f'{host}:{port}')
        self.stub = tn_pb2_grpc.TNStub(self.channel)

    def create_callback(self, callback):
        request = tn_pb2.RequestCreateCallback(callback=callback)
        return self.stub.CreateCallBack(request)

    def get_callback(self, id):
        request = tn_pb2.RequestGetCallBack(id=id)
        return self.stub.GetCallBack(request)

    def get_all_callbacks(self, number):
        request = tn_pb2.RequestGetAllCallBacks(Number=number)
        return self.stub.GetAllCallBacks(request)

    def delete_callback(self, id):
        request = tn_pb2.RequestDeleteCallBack(id=id)
        return self.stub.DeleteCallBack(request)