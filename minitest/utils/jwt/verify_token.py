import jwt
from datetime import datetime
from django.conf import settings

def verify_jwt_token(token):
    """
    JWT 토큰을 검증하고 payload를 반환
    """
    try:
        payload = jwt.decode(
            token,
            settings.JWT_SECRET_KEY,
            algorithms=['HS512']  # HS512만 허용
        )
        return payload
    except jwt.ExpiredSignatureError:
        raise jwt.ExpiredSignatureError("토큰이 만료되었습니다.")
    except jwt.InvalidTokenError:
        raise jwt.InvalidTokenError("유효하지 않은 토큰입니다.")

def extract_token_from_header(auth_header):
    """
    Authorization 헤더에서 토큰을 추출
    """
    if not auth_header:
        return None
    
    if not auth_header.startswith('Bearer '):
        return None
    
    return auth_header.split(' ')[1]

def is_token_near_expiry(token_exp, hours=1):
    """
    토큰이 만료에 가까운지 확인
    """
    current_time = datetime.now().timestamp()
    return token_exp - current_time <= (hours * 3600)

def get_user_from_token(request):
    """
    요청에서 토큰 정보를 기반으로 사용자 정보를 가져옴
    """
    return {
        'id': getattr(request, 'user_id', None),
        'email': getattr(request, 'user_email', None),
        'name': getattr(request, 'user_name', None),
        'token_exp': getattr(request, 'token_exp', None),
        'token_iat': getattr(request, 'token_iat', None)
    }