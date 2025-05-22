from pydantic import BaseModel, Field
from typing import Optional


# 회사 입력 모델
class CompanyReqDTO(BaseModel):
    # 여기에서 벨리데이션 처리
    name: str = Field(..., min_length=1, max_length=50)
    company_name: str = Field(..., min_length=1, max_length=100)


# 회사 생성 요청 DTO
class CreateCompanyReqDTO(BaseModel):
    user_id: str = Field(
        ...,
        description="회사 소유자(유저) ID",
        example="550e8400-e29b-41d4-a716-446655440000",
    )
    company_name: str = Field(..., description="회사 이름", example="블루베리(주)")
    company_address: Optional[str] = Field(
        "", description="회사 주소", example="서울시 강남구"
    )
    total_staff: Optional[int] = Field(0, description="총 직원 수", example=10)


# 회사 수정 요청 DTO
class ChangeCompanyReqDTO(BaseModel):
    id: str = Field(
        ..., description="회사 ID", example="c1e2d3f4-5678-1234-9abc-1234567890ab"
    )
    company_name: Optional[str] = Field(
        None, description="회사 이름(선택)", example="블루베리(주)"
    )
    company_address: Optional[str] = Field(
        None, description="회사 주소(선택)", example="서울시 강남구"
    )
    total_staff: Optional[int] = Field(None, description="총 직원 수(선택)", example=15)


# 회사 삭제 요청 DTO
class DeleteCompanyReqDTO(BaseModel):
    id: str = Field(
        ..., description="회사 ID", example="c1e2d3f4-5678-1234-9abc-1234567890ab"
    )
