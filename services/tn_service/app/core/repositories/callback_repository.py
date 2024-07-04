from sqlalchemy.orm import Session

from app.core.models.callback_model import CallbackModel
from app.core.schemas.callback_schema import CallbackSchema


class CallbackRepository:
    def __init__(self, db_session: Session):
        self.db_session = db_session

    def create_callback(self, callback_create: CallbackSchema):
        callback_db = CallbackModel(name=callback_create.name, number=callback_create.number, date=callback_create.date)
        self.db_session.add(callback_db)
        self.db_session.commit()
        self.db_session.refresh(callback_db)
        return callback_db

    def get_callback(self, callback_id: int):
        return self.db_session.query(CallbackModel).get(callback_id)

    def get_callbacks_paginated(self, limit: int, offset: int):
        total_items = self.db_session.query(CallbackModel).count()
        items = self.db_session.query(CallbackModel).offset(offset).limit(limit).all()
        response = {'total_items': total_items,
                    'items': items}
        return response

    def get_all_callbacks(self):
        total_items = self.db_session.query(CallbackModel).count()
        items = self.db_session.query(CallbackModel).all()
        response = {'total_items': total_items,
                   'items': items}
        return response

    def delete_callback(self, callback_id: int):
        callback_db = self.db_session.query(CallbackModel).get(callback_id)
        if callback_db:
            self.db_session.delete(callback_db)
            self.db_session.commit()
            return callback_id
        return None

