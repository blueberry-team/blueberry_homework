from pydantic import BaseModel, EmailStr, Field
from typing import Optional
from datetime import datetime


# 인증 관련 응답 DTO (회원가입/로그인/수정 등)
class AuthResDTO(BaseModel):
    message: str = Field(..., description="결과 메시지", example="회원가입 성공")
    user_id: Optional[str] = Field(
        None, description="유저 ID", example="550e8400-e29b-41d4-a716-446655440000"
    )


# 유저 정보 응답 DTO
class UserResDTO(BaseModel):
    user_id: str = Field(
        ..., description="유저 ID", example="550e8400-e29b-41d4-a716-446655440000"
    )
    name: str = Field(..., description="이름", example="홍길동")
    email: EmailStr = Field(..., description="이메일", example="test@example.com")
    address: Optional[str] = Field("", description="주소", example="서울시 강남구")
    role: Optional[str] = Field("worker", description="역할", example="boss")
    created_at: Optional[datetime] = Field(
        None, description="생성일", example="2024-05-20T12:00:00Z"
    )
    updated_at: Optional[datetime] = Field(
        None, description="수정일", example="2024-05-20T12:00:00Z"
    )
