from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

from app.config.config import config

engine = create_engine(config['database']['url'])
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)
