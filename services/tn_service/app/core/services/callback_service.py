from google.protobuf import empty_pb2
import grpc

from app.api.proto import tn_pb2
from app.core.repositories.callback_repository import CallbackRepository
from app.core.schemas.callback_schema import CallbackSchema


class CallbackService:

    def __init__(self, db_session):
        self.db_session = db_session

    def create_callback(self, request, context):
        callback_data = request.callback
        callback = CallbackSchema(name=callback_data.Name, date=callback_data.Date, number=callback_data.Number)
        with self.db_session() as db:
            repository = CallbackRepository(db)
            repository.create_callback(callback)
        return empty_pb2.Empty()

    def get_callback(self, request, context):
        with self.db_session() as db:
            repository = CallbackRepository(db)
            callback = repository.get_callback(request.id)
            if callback is None:
                context.abort(grpc.StatusCode.NOT_FOUND, "Callback not found")
            return tn_pb2.Callback(
                Name=callback.name,
                Date=callback.date,
                Number=callback.number
            )

    def get_all_callbacks(self, request, context):
        with self.db_session() as db:
            repository = CallbackRepository(db)
            response = repository.get_all_callbacks()
            if not response:
                context.abort(grpc.StatusCode.NOT_FOUND, "Callbacks not found")

            callbacks = [
                {
                    'Name': item.name,
                    'Date': item.date,
                    'Number': item.number
                }
                for item in response['items']
            ]

            items = [tn_pb2.Callback(**callback) for callback in callbacks]

            tn_callback = tn_pb2.ResponseGetAllCallbacks(
                total_items=response['total_items'],
                callbacks=items
            )
            return tn_callback

    def get_callbacks_paginated(self, request, context):
        with self.db_session() as db:
            repository = CallbackRepository(db)
            response = repository.get_callbacks_paginated(request.limit if request.limit else 1, request.offset)
            if not response:
                context.abort(grpc.StatusCode.NOT_FOUND, "Callbacks not found")

            callbacks = [
                {
                    'Name': item.name,
                    'Date': item.date,
                    'Number': item.number
                }
                for item in response['items']
            ]

            items = [tn_pb2.Callback(**callback) for callback in callbacks]

            tn_callback = tn_pb2.ResponseGetAllCallbacks(
                total_items=response['total_items'],
                callbacks=items
            )
            return tn_callback

    def delete_callback_by_id(self, request, context):
        with self.db_session() as db:
            repository = CallbackRepository(db)
            deleted_id = repository.delete_callback_by_id(request.id)
            if deleted_id is None:
                context.abort(grpc.StatusCode.NOT_FOUND, "Callback not found")
            return empty_pb2.Empty()

    def delete_callback_by_number(self, request, context):
        with self.db_session() as db:
            repository = CallbackRepository(db)
            deleted_number = repository.delete_callback_by_number(request.Number)
            if deleted_number is None:
                context.abort(grpc.StatusCode.NOT_FOUND, "Callback not found")
            return empty_pb2.Empty()
