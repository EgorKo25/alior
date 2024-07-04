import grpc

from app.api.proto import tn_pb2_grpc
from app.core.services.callback_service import CallbackService
from app.utils.database import SessionLocal

from app.config.config import config


class TNService(tn_pb2_grpc.TNServicer):

    def __init__(self, db_session):
        self.callback_service = CallbackService(db_session)

    def CreateCallback(self, request, context):
        return self.callback_service.create_callback(request, context)

    def GetCallback(self, request, context):
        return self.callback_service.get_callback(request, context)

    def GetAllCallbacks(self, request, context):
        return self.callback_service.get_all_callbacks(request, context)

    def GetCallbacksPaginated(self, request, context):
        return self.callback_service.get_callbacks_paginated(request, context)

    def DeleteCallbackById(self, request, context):
        return self.callback_service.delete_callback_by_id(request, context)

    def DeleteCallbackByNumber(self, request, context):
        return self.callback_service.delete_callback_by_number(request, context)


async def serve():
    server = grpc.aio.server()
    tn_pb2_grpc.add_TNServicer_to_server(TNService(SessionLocal), server)
    server.add_insecure_port(f"[::]:{config['app']['port']}")
    await server.start()
    await server.wait_for_termination()
