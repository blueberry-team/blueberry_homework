# Entity
from dataclasses import dataclass
from datetime import datetime

@dataclass()
class User:
    id: str #uuid
    email: str
    password: str
    name: str
    role: str
    created_at: datetime
    updated_at: datetime

    def validate(self):
        if not self.name or len(self.name) < 1:
            raise ValueError('1자 이상 적어주세요')
        if len(self.name) > 50:
            raise ValueError('50자 이하 적어주세요')
        if not (self.role=="boss" or self.role=="worker"):
            raise ValueError('role은 boss 또는 worker만 가능합니다')
            