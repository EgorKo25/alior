from google.protobuf import empty_pb2
import grpc

from app.api.proto import tn_pb2
from app.core.repositories.callback_repository import CallBackRepository
from app.core.schemas.callback_schema import CallBackSchema


class CallbackService:

    def __init__(self, db_session):
        self.db_session = db_session

    def create_callback(self, request, context):
        callback_data = request.callback
        callback = CallBackSchema(name=callback_data.Name, date=callback_data.Date, number=callback_data.Number)
        with self.db_session() as db:
            repository = CallBackRepository(db)
            repository.create_callback(callback)
        return empty_pb2.Empty()

    def get_callback(self, request, context):
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

    def update_callback(self, request, context):
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

    def delete_callback(self, request, context):
        with self.db_session() as db:
            repository = CallBackRepository(db)
            deleted_id = repository.delete_callback(request.id)
            if deleted_id is None:
                context.abort(grpc.StatusCode.NOT_FOUND, "Callback not found")
            return tn_pb2.ResponseDeleteCallBack(id=deleted_id)
