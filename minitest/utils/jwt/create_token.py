import jwt
from datetime import datetime, timedelta
from django.conf import settings

def create_jwt_token(user):
    """
    사용자 정보를 기반 JWT 토큰 생성
    """
    now = datetime.utcnow()
    expiration_time = now + timedelta(hours=5)
    
    payload = {
        "sub": str(user.id), 
        "email": user.email,
        "name": user.name,  
        "exp": expiration_time,
        "iat": now
    }
    
    token = jwt.encode(
        payload, 
        settings.JWT_SECRET_KEY, 
        algorithm='HS512' # SHA-512 해시 알고리즘
    )
    
    return token