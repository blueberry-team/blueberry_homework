from pydantic import BaseModel, EmailStr, Field
from typing import Optional


# 회원가입 요청 DTO
class SignUpReqDTO(BaseModel):
    name: str = Field(..., description="사용자 이름", example="홍길동")
    email: EmailStr = Field(..., description="이메일", example="test@example.com")
    password: str = Field(..., description="비밀번호", example="1234abcd!")
    address: Optional[str] = Field("", description="주소", example="서울시 강남구")
    role: Optional[str] = Field(
        "worker", description="역할(boss/worker)", example="boss"
    )


# 로그인 요청 DTO
class LogInReqDTO(BaseModel):
    email: EmailStr = Field(..., description="이메일", example="test@example.com")
    password: str = Field(..., description="비밀번호", example="1234abcd!")


# 유저정보 수정 요청 DTO
class ChangeUserReqDTO(BaseModel):
    user_id: str = Field(
        ..., description="유저 ID", example="550e8400-e29b-41d4-a716-446655440000"
    )
    name: Optional[str] = Field(None, description="이름(선택)", example="홍길동")
    email: Optional[EmailStr] = Field(
        None, description="이메일(선택)", example="test@example.com"
    )
    password: Optional[str] = Field(
        None, description="비밀번호(선택)", example="1234abcd!"
    )
    address: Optional[str] = Field(
        None, description="주소(선택)", example="서울시 강남구"
    )
    role: Optional[str] = Field(None, description="역할(선택)", example="worker")
