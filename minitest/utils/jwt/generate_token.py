import jwt
from datetime import timedelta
from django.utils import timezone
from django.conf import settings

def create_jwt_token(user):
    """
    사용자 정보를 기반 JWT 토큰 생성 (now: asia/seoul)
    """
    now = timezone.now()
    expiration_time = now + timedelta(hours=5)   # 검증 시 timezone.now() 사용 
    iat_timestamp = int(now.timestamp())
    exp_timestamp = int(expiration_time.timestamp())
    
    payload = {
        "sub": str(user.id), 
        "email": user.email,
        "name": user.name,  
        "exp": exp_timestamp,
        "iat": iat_timestamp
    }
    
    token = jwt.encode(
        payload, 
        settings.JWT_SECRET_KEY, 
        algorithm='HS512' # SHA-512 해시 알고리즘
    )
    
    return token