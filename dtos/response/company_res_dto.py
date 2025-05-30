from pydantic import BaseModel, Field
from typing import Optional, List
from datetime import datetime


# 회사 단일 응답 DTO
class CompanyResDTO(BaseModel):
    id: str = Field(
        ..., description="회사 ID", example="c1e2d3f4-5678-1234-9abc-1234567890ab"
    )
    user_id: str = Field(
        ...,
        description="소유자(유저) ID",
        example="550e8400-e29b-41d4-a716-446655440000",
    )
    company_name: str = Field(..., description="회사 이름", example="블루베리(주)")
    company_address: Optional[str] = Field(
        "", description="회사 주소", example="서울시 강남구"
    )
    total_staff: Optional[int] = Field(0, description="총 직원 수", example=10)
    created_at: Optional[datetime] = Field(
        None, description="생성일", example="2024-05-20T12:00:00Z"
    )
    updated_at: Optional[datetime] = Field(
        None, description="수정일", example="2024-05-20T12:00:00Z"
    )


# 회사 목록 응답 DTO
class CompanyListResDTO(BaseModel):
    companies: List[CompanyResDTO] = Field(..., description="회사 목록")
