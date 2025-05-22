from pydantic import BaseModel, EmailStr, Field
from typing import Optional


# 회원가입 요청 DTO
class SignUpReqDTO(BaseModel):
    name: str  # 사용자 이름
    email: EmailStr  # 이메일
    password: str  # 비밀번호
    address: Optional[str] = ""  # 주소(선택)
    role: Optional[str] = "worker"  # 역할(기본값: worker)


# 로그인 요청 DTO
class LogInReqDTO(BaseModel):
    email: EmailStr  # 이메일
    password: str  # 비밀번호


# 유저정보 수정 요청 DTO
class ChangeUserReqDTO(BaseModel):
    user_id: str  # 유저 ID
    name: Optional[str] = None  # 이름(선택)
    email: Optional[EmailStr] = None  # 이메일(선택)
    password: Optional[str] = None  # 비밀번호(선택)
    address: Optional[str] = None  # 주소(선택)
    role: Optional[str] = None  # 역할(선택)
