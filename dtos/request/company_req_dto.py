from pydantic import BaseModel, Field


# 회사 입력 모델
class CompanyReqDTO(BaseModel):
    # 여기에서 벨리데이션 처리
    name: str = Field(..., min_length=1, max_length=50)
    company_name: str = Field(..., min_length=1, max_length=100)
