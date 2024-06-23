import grpc
import app.api.proto.tn_pb2_grpc as tn_pb2_grpc
import app.api.proto.tn_pb2 as tn_pb2

async def run_client():
    async with grpc.aio.insecure_channel('localhost:50051') as channel:
        stub = tn_pb2.TNStub(channel)
        pass