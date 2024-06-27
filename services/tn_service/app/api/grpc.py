import grpc

from app.api.proto import tn_pb2


async def run_client():
    async with grpc.aio.insecure_channel('localhost:50051') as channel:
        stub = tn_pb2.TNStub(channel)
        print(stub)
