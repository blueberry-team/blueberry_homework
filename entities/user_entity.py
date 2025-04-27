from datetime import datetime
from pydantic import BaseModel


class UserEntity(BaseModel):
    id: str
    name: str
    updated_at: datetime
    created_at: datetime

    def __str__(self):
        return f"UserEntity(id={self.id}, name={self.name}, created_at={self.created_at}, updated_at={self.updated_at})"