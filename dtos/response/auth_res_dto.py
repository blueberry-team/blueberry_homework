from pydantic import BaseModel, EmailStr
from typing import Optional
from datetime import datetime


# 인증 관련 응답 DTO (회원가입/로그인/수정 등)
class AuthResDTO(BaseModel):
    message: str  # 결과 메시지
    user_id: Optional[str] = None  # 유저 ID(선택)


# 유저 정보 응답 DTO
class UserResDTO(BaseModel):
    user_id: str  # 유저 ID
    name: str  # 이름
    email: EmailStr  # 이메일
    address: Optional[str] = ""  # 주소(선택)
    role: Optional[str] = "worker"  # 역할(기본값: worker)
    created_at: Optional[datetime] = None  # 생성일(선택)
    updated_at: Optional[datetime] = None  # 수정일(선택)
