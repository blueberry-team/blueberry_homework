import jwt
from django.http import JsonResponse
from django.conf import settings
from django.utils.deprecation import MiddlewareMixin
from django.utils import timezone

class JWTAuthenticationMiddleware(MiddlewareMixin):
    """
    JWT 토큰 검증 미들웨어 (timezone.now() 사용)
    """
    
    EXEMPT_PATHS = [
        '/signUp/',
        '/signIn/',
        '/admin/',
        '/static/',
        '/media/',
    ]
    
    def process_request(self, request):
        if self._is_exempt_path(request.path):
            return None
        
        auth_header = request.META.get('HTTP_AUTHORIZATION')
        
        if not auth_header or not auth_header.startswith('Bearer '):
            return self._unauthorized_response("Authorization 헤더가 필요합니다.")
        
        token = auth_header.split(' ')[1]
        
        try:
            payload = jwt.decode(
                token, 
                settings.JWT_SECRET_KEY, 
                algorithms=['HS512'],
            )
            
            request.user_id = payload.get('sub')
            request.user_email = payload.get('email')
            request.user_name = payload.get('name')
            request.token_exp = payload.get('exp')
            request.token_iat = payload.get('iat')
            
            return None
            
        except Exception as e:
            return self._unauthorized_response("유효하지 않은 토큰입니다.")
    
    def process_response(self, request, response):
        if self._is_exempt_path(request.path):
            return response
        
        if hasattr(request, 'token_exp'):
            current_time = timezone.now().timestamp()
            token_exp_time = request.token_exp
            time_left = token_exp_time - current_time
            if time_left <= 3600:  # 1시간 이내
                response['X-Token-Refresh-Required'] = 'true'
        
        return response
    
    def _is_exempt_path(self, path):
        return any(path.startswith(exempt_path) for exempt_path in self.EXEMPT_PATHS)
    
    def _unauthorized_response(self, message):
        return JsonResponse({"success": False, "message": message}, status=401)