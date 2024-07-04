from sqlalchemy import Column, Integer, String
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()


class CallbackModel(Base):
    __tablename__ = 'callbacks'

    id = Column(Integer, primary_key=True, index=True)
    name = Column(String)
    date = Column(String)
    number = Column(String)
