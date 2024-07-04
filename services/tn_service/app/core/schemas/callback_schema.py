from pydantic import BaseModel


class CallbackSchema(BaseModel):
    name: str
    date: str
    number: str
