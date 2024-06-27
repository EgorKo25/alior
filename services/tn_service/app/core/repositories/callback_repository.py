from sqlalchemy.orm import Session

from app.core.models.callback_model import CallBackModel
from app.core.schemas.callback_schema import CallBackSchema


class CallBackRepository:
    def __init__(self, db_session: Session):
        self.db_session = db_session

    def create_callback(self, callback_create: CallBackSchema):
        callback_db = CallBackModel(name=callback_create.name, number=callback_create.number, date=callback_create.date)
        self.db_session.add(callback_db)
        self.db_session.commit()
        self.db_session.refresh(callback_db)
        return callback_db

    def get_callback(self, callback_id: int):
        return self.db_session.query(CallBackModel).get(callback_id)

    def update_callback(self, callback_id: int, callback_update: CallBackSchema):
        callback_db = self.db_session.query(CallBackModel).get(callback_id)
        if callback_db:
            callback_db.name = callback_update.name
            callback_db.date = callback_update.date
            callback_db.number = callback_update.number
            self.db_session.commit()
            self.db_session.refresh(callback_db)
        return callback_db

    def delete_callback(self, callback_id: int):
        callback_db = self.db_session.query(CallBackModel).get(callback_id)
        if callback_db:
            self.db_session.delete(callback_db)
            self.db_session.commit()
            return callback_id
        return None
