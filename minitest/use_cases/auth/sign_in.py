import bcrypt

from ...utils.jwt.generate_token import create_jwt_token
from ...repositores.name_repository import NameRepository
from django.core.validators import ValidationError

name_repo = NameRepository()
def verify_password(plain_password, hashed_password):
    return bcrypt.checkpw(plain_password.encode('utf-8'), hashed_password.encode('utf-8'))

def login_user(email, password):
    try:
        user = name_repo.model.objects.get(email=email)
    except name_repo.model.DoesNotExist:
        raise ValidationError('해당 이메일이 존재하지 않습니다')
    
    if not verify_password(password, user.password):
        raise ValidationError('비밀번호가 틀립니다')
    
    try:
        user = create_login_token(user)
    except Exception as e:
        raise ValidationError(f'로그인 토큰 생성 중 오류 발생: {str(e)}')
    
    return user

def create_login_token(user):
    token = create_jwt_token(user)
    setattr(user, 'token', token)
    
    return user
