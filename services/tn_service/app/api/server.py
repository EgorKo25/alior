import grpc

from app.api.proto import tn_pb2_grpc
from app.core.services.callback_service import CallbackService
from app.utils.database import SessionLocal

from app.config.config import config


class TNService(tn_pb2_grpc.TNServicer):

    def __init__(self, db_session):
        self.callback_service = CallbackService(db_session)

    def CreateCallBack(self, request, context):
        return self.callback_service.create_callback(request, context)

    def GetCallBack(self, request, context):
        return self.callback_service.get_callback(request, context)

    def GetAllCallBacks(self, request, context):
        return self.callback_service.get_all_callbacks(request, context)

    def DeleteCallBack(self, request, context):
        return self.callback_service.delete_callback(request, context)


async def serve():
    server = grpc.aio.server()
    tn_pb2_grpc.add_TNServicer_to_server(TNService(SessionLocal), server)
    server.add_insecure_port(f"[::]:{config['app']['port']}")
    await server.start()
    await server.wait_for_termination()
