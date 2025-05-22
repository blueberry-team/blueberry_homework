from pydantic import BaseModel, Field
from typing import Optional


# 회사 입력 모델
class CompanyReqDTO(BaseModel):
    # 여기에서 벨리데이션 처리
    name: str = Field(..., min_length=1, max_length=50)
    company_name: str = Field(..., min_length=1, max_length=100)


# 회사 생성 요청 DTO
class CreateCompanyReqDTO(BaseModel):
    user_id: str  # 회사 소유자(유저) ID
    company_name: str  # 회사 이름
    company_address: Optional[str] = ""  # 회사 주소
    total_staff: Optional[int] = 0  # 총 직원 수


# 회사 수정 요청 DTO
class ChangeCompanyReqDTO(BaseModel):
    id: str  # 회사 ID
    company_name: Optional[str] = None  # 회사 이름(선택)
    company_address: Optional[str] = None  # 회사 주소(선택)
    total_staff: Optional[int] = None  # 총 직원 수(선택)


# 회사 삭제 요청 DTO
class DeleteCompanyReqDTO(BaseModel):
    id: str  # 회사 ID
