from datetime import datetime
from pydantic import BaseModel
from typing import Optional


# 회사 엔티티
class CompanyEntity(BaseModel):
    id: str  # 회사 고유 ID
    user_id: str  # 회사 소유자(유저) ID
    company_name: str  # 회사 이름
    company_address: Optional[str] = ""  # 회사 주소
    total_staff: Optional[int] = 0  # 총 직원 수
    created_at: datetime  # 생성일
    updated_at: Optional[datetime] = None  # 수정일

    def __str__(self):
        return f"CompanyEntity(id={self.id}, user_id={self.user_id}, company_name={self.company_name}, company_address={self.company_address}, total_staff={self.total_staff}, created_at={self.created_at}, updated_at={self.updated_at})"
