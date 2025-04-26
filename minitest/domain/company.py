# Entity
from dataclasses import dataclass
from datetime import datetime

@dataclass()
class Company:
    id: str #uuid
    name: str
    company_name: str
    created_at: datetime