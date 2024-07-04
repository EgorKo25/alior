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

    def get_all_callbacks(self, request, context):
        with self.db_session() as db:
            repository = CallBackRepository(db)
            callbacks = repository.get_all_callbacks()
            if not callbacks:
                context.abort(grpc.StatusCode.NOT_FOUND, "Callbacks not found")
            for callback in callbacks:
                tn_callback = tn_pb2.CallBack(
                    Name=callback.name,
                    Date=callback.date,
                    Number=callback.number
                )
                yield tn_callback

    def get_callbacks_paginated(self, request, context):
        with self.db_session() as db:
            repository = CallBackRepository(db)
            callbacks = repository.get_callbacks_paginated(request.limit, request.offset)
            if not callbacks:
                context.abort(grpc.StatusCode.NOT_FOUND, "Callbacks not found")
            for callback in callbacks:
                tn_callback = tn_pb2.CallBack(
                    Name=callback.name,
                    Date=callback.date,
                    Number=callback.number
                )
                yield tn_callback

    def get_callbacks_quantity(self, request, context):
        with self.db_session() as db:
            repository = CallBackRepository(db)
            callbacks_quantity = repository.get_callbacks_quantity()
            return tn_pb2.ResponseGetCallBacksQuantity(quantity=callbacks_quantity)

    def delete_callback(self, request, context):
        with self.db_session() as db:
            repository = CallBackRepository(db)
            deleted_id = repository.delete_callback(request.id)
            if deleted_id is None:
                context.abort(grpc.StatusCode.NOT_FOUND, "Callback not found")
            return tn_pb2.ResponseDeleteCallBack(id=deleted_id)
