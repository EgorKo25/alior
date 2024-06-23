import grpc
from google.protobuf import empty_pb2
from app.core.schemas import CallBackCreate
from app.core.repositories import CallBackRepository
from app.core.models import SessionLocal
from .proto import tn_pb2_grpc
from .proto import tn_pb2

class TNService(tn_pb2_grpc.TNServicer):

    def __init__(self, db_session):
        self.db_session = db_session

    def CreateCallBack(self, request, context):
        callback_data = request.callback
        callback_create = CallBackCreate(name=callback_data.Name, date=callback_data.Date, number=callback_data.Number)
        with self.db_session() as db:
            repository = CallBackRepository(db)
            repository.create_callback(callback_create)
        return empty_pb2.Empty()

async def serve():
    server = grpc.aio.server()
    tn_pb2_grpc.add_TNServicer_to_server(TNService(SessionLocal), server)
    server.add_insecure_port('[::]:50051')
    await server.start()
    await server.wait_for_termination()