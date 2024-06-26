import grpc
from google.protobuf import empty_pb2
from app.core.schemas import CallBackSchema
from app.core.repositories import CallBackRepository
from app.core.models import SessionLocal
from app.api.proto import tn_pb2, tn_pb2_grpc

class TNService(tn_pb2_grpc.TNServicer):

    def __init__(self, db_session):
        self.db_session = db_session

    def CreateCallBack(self, request, context):
        callback_data = request.callback
        callback = CallBackSchema(name=callback_data.Name, date=callback_data.Date, number=callback_data.Number)
        with self.db_session() as db:
            repository = CallBackRepository(db)
            repository.create_callback(callback)
        return empty_pb2.Empty()

    def GetCallBack(self, request, context):
        with self.db_session() as db:
            repository = CallBackRepository(db)
            callback = repository.get_callback(request.id)
            if callback is None:
                context.abort(grpc.StatusCode.NOT_FOUND, "Callback not found")
            return tn_pb2.CallBack(
                Name=callback.name,
                Date=callback.date,
                Number=callback.number
            )

    def UpdateCallBack(self, request, context):
        callback_data = CallBackSchema(
            name=request.callback.Name,
            date=request.callback.Date,
            number=request.callback.Number
        )
        with self.db_session() as db:
            repository = CallBackRepository(db)
            callback = repository.update_callback(request.id, callback_data)
            if callback is None:
                context.abort(grpc.StatusCode.NOT_FOUND, "Callback not found")
            return tn_pb2.CallBack(
                Name=callback.name,
                Date=callback.date,
                Number=callback.number
            )

    def DeleteCallBack(self, request, context):
        with self.db_session() as db:
            repository = CallBackRepository(db)
            deleted_id = repository.delete_callback(request.id)
            if deleted_id is None:
                context.abort(grpc.StatusCode.NOT_FOUND, "Callback not found")
            return tn_pb2.ResponseDeleteCallBack(id=deleted_id)

async def serve():
    server = grpc.aio.server()
    tn_pb2_grpc.add_TNServicer_to_server(TNService(SessionLocal), server)
    server.add_insecure_port('[::]:50051')
    await server.start()
    await server.wait_for_termination()