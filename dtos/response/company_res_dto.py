from pydantic import BaseModel
from typing import Optional, List
from datetime import datetime


# 회사 단일 응답 DTO
class CompanyResDTO(BaseModel):
    id: str  # 회사 ID
    user_id: str  # 소유자(유저) ID
    company_name: str  # 회사 이름
    company_address: Optional[str] = ""  # 회사 주소
    total_staff: Optional[int] = 0  # 총 직원 수
    created_at: Optional[datetime] = None  # 생성일
    updated_at: Optional[datetime] = None  # 수정일


# 회사 목록 응답 DTO
class CompanyListResDTO(BaseModel):
    companies: List[CompanyResDTO]
