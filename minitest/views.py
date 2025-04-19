from rest_framework.views import APIView # type: ignore
from rest_framework.response import Response # type: ignore
from rest_framework import status # type: ignore
from django.core.exceptions import ValidationError # type: ignore
from rest_framework.exceptions import APIException # type: ignore
from .use_cases.user import get_user, create_user, delete_user

def _user_to_dict(user):
    # User 객체를 딕셔너리로 변환하는 헬퍼 메소드
    return {
        'name': user.name,
        'created_at': user.created_at.strftime("%Y-%m-%d %H:%M:%S") if user.created_at else None
    }

class NameAPIView(APIView):

    def get(self, request):
        try:
            users = get_user()
            return Response({
                "message": "success",
                "data": [_user_to_dict(user) for user in users]
            })
        except APIException as e:
            return Response({
                "message": "error",
                "data": str(e)
            })

    def post(self, request): 
        name = request.body.decode('utf-8')
        # 리포지토리에서 발생시킨 예외를, 뷰에서 적절한 HTTP 응답으로 변환하여 코드 구조를 분리함 
        # (비즈니스 규칙 에러는 리포지토리에서, HTTP 응답 관리는 뷰에서 처리함)
        try:
            users = create_user(name)
            return Response({
                "message": "success",
                "data": [_user_to_dict(user) for user in users]
            })
        except ValidationError as e:
            error_response = {
                    "message": "error",
                    "error": e
                }
            return Response(error_response, status=status.HTTP_400_BAD_REQUEST)

class NameDeleteAPIView(APIView):

    def get(self, request, idx):
        try:
            users = get_user(idx)
            return Response({
                "message": "success",
                "data": [_user_to_dict(user) for user in users]
            })
        except ValidationError as e:
            error_response = {
                    "message": "error",
                    "error": e
                }
            return Response(error_response, status=status.HTTP_400_BAD_REQUEST)
        
    def delete(self, request, idx):
        try:
            users = delete_user(idx)
            return Response({
                "message": "success",
                "data": [_user_to_dict(user) for user in users]
            })
        except APIException as e:
            return Response({
                "message": "error",
                "data": str(e)
            })
        