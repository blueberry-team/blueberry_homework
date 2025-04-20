from datetime import datetime
from pydantic import BaseModel

class NameEntity(BaseModel):
    name: str
    created_at: datetime

    def __str__(self):
        return f"NameEntity(name={self.name}, created_at={self.created_at})"
