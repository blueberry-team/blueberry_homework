# 이름 입력 모델
from pydantic import BaseModel, Field


# 이름 입력 모델
class NameReqDTO(BaseModel):
    # 여기에서 벨리데이션 처리
    name: str = Field(
        ..., description="이름을 입력해주세요", min_length=1, max_length=50
    )
