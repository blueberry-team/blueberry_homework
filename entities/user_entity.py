from datetime import datetime
from pydantic import BaseModel


# 사용자 엔티티
class UserEntity(BaseModel):
    id: str
    name: str
    email: str
    password: str
    address: str
    role: str
    updated_at: datetime
    created_at: datetime

    def __str__(self):
        return f"UserEntity(id={self.id}, name={self.name}, email={self.email}, role={self.role}, created_at={self.created_at}, updated_at={self.updated_at})"
