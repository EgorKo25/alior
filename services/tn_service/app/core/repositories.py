from app.core.models import Base
from sqlalchemy.orm import Session
from .schemas import CallBackCreate
from .models import Base, SessionLocal

class CallBackRepository:
    def __init__(self, db_session: Session):
        self.db_session = db_session

    def create_callback(self, callback_create: CallBackCreate):
        from app.core.models import CallBack
        callback_db = CallBack(**callback_create.dict())
        self.db_session.add(callback_db)
        self.db_session.commit()
        self.db_session.refresh(callback_db)
        return callback_db