from pydantic import BaseModel, Field
from typing import List


# 단일 이름 응답 DTO
class NameResDTO(BaseModel):
    name: str = Field(..., description="이름", example="홍길동")


# 이름 목록 응답 DTO
class NameListResDTO(BaseModel):
    names: List[NameResDTO] = Field(..., description="이름 목록")
