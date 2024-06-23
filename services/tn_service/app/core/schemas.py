from pydantic import BaseModel

class CallBackCreate(BaseModel):
    name: str
    date: str
    number: str

class CallBack(BaseModel):
    id: int
    name: str
    date: str
    number: str

    class Config:
        orm_mode = True