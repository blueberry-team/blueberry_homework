# Entity
from dataclasses import dataclass
from datetime import datetime

@dataclass()
class Company:
    id: str #uuid
    user_id: str
    # name: str
    company_name: str
    company_address: str
    total_staff: int
    created_at: datetime
    updated_at: datetime