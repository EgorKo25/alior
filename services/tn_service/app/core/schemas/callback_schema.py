from pydantic import BaseModel


class CallBackSchema(BaseModel):
    name: str
    date: str
    number: str
