from datetime import datetime
from pydantic import BaseModel


class CompanyEntity(BaseModel):
    id: str
    name: str
    company_name: str
    created_at: datetime

    def __str__(self):
        return f"CompanyEntity(id={self.id}, name={self.name}, company_name={self.company_name}, created_at={self.created_at})"
