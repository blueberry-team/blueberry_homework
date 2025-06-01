import jwt
import json
from datetime import datetime, timedelta
from django.http import JsonResponse
from django.conf import settings
from django.utils.deprecation import MiddlewareMixin

class JWTAuthenticationMiddleware(MiddlewareMixin):
    """
    JWT 토큰 검증 미들웨어
    """
    
    # 인증이 필요하지 않은 경로들
    EXEMPT_PATHS = [
        '/signUp/',
        '/signIn/',
        '/api/auth/refresh-token/',  # 토큰 갱신 경로
        '/admin/',  # Django admin
        '/static/',  # 정적 파일
        '/media/',   # 미디어 파일
    ]
    
    def process_request(self, request):
        """
        요청 처리 전 토큰 검증
        """
        # 인증이 필요하지 않은 경로는 패스
        if self._is_exempt_path(request.path):
            return None
        
        # Authorization 헤더에서 토큰 추출
        auth_header = request.META.get('HTTP_AUTHORIZATION')
        
        if not auth_header:
            return self._unauthorized_response("Authorization 헤더가 필요합니다.")
        
        # Bearer 토큰 형식 확인
        if not auth_header.startswith('Bearer '):
            return self._unauthorized_response("Bearer 토큰 형식이 아닙니다.")
        
        token = auth_header.split(' ')[1]
        
        try:
            # 토큰 검증 및 디코딩
            payload = jwt.decode(
                token, 
                settings.JWT_SECRET_KEY, 
                algorithms=['HS512']  # HS512만 허용 (create_jwt_token에서 사용한 알고리즘과 일치해야 함)
            )
            
            # 토큰에서 사용자 정보 추출
            request.user_id = payload.get('sub')
            request.user_email = payload.get('email')
            request.user_name = payload.get('name')
            request.token_exp = payload.get('exp')
            request.token_iat = payload.get('iat')
            
            return None
            
        except jwt.ExpiredSignatureError:
            return self._unauthorized_response("토큰이 만료되었습니다.")
        except jwt.InvalidTokenError:
            return self._unauthorized_response("유효하지 않은 토큰입니다.")
        except Exception as e:
            return self._unauthorized_response(f"토큰 검증 중 오류가 발생했습니다: {str(e)}")
    
    def process_response(self, request, response):
        """
        응답 처리 시 토큰 갱신 필요 여부 확인
        """
        # 인증이 필요하지 않은 경로는 패스
        if self._is_exempt_path(request.path):
            return response
        
        # 토큰 만료 시간이 있는 경우에만 체크
        if hasattr(request, 'token_exp'):
            current_time = datetime.utcnow().timestamp()
            token_exp_time = request.token_exp
            
            # 토큰 만료까지 1시간 이내인 경우
            if token_exp_time - current_time <= 3600:  # 3600초 = 1시간
                response['X-Token-Refresh-Required'] = 'true'
        
        return response
    
    def _is_exempt_path(self, path):
        """
        인증이 필요하지 않은 경로인지 확인
        """
        return any(path.startswith(exempt_path) for exempt_path in self.EXEMPT_PATHS)
    
    def _unauthorized_response(self, message):
        """
        인증 실패 응답 생성
        """
        return JsonResponse({
            "success": False,
            "message": message
        }, status=401)