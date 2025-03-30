# 이름 입력 모델
from pydantic import BaseModel, Field


class NameModel(BaseModel):
    name: str = Field(
        ..., description="이름을 입력해주세요", min_length=1, max_length=50
    )
