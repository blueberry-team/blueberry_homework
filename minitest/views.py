from rest_framework.views import APIView
from .repositories import NameRepository
from rest_framework.response import Response
from rest_framework import status


class NameAPIView(APIView):
    name_repo = NameRepository()

    def get(self, request):
        result = self.name_repo.get_name()
        return Response(result)

    def post(self, request): 
        name = request.body.decode('utf-8')

        ### 유효성 검증 임시 위치
        if not name or len(name) < 1 or len(name) > 50:
            ### 이러한 validate 로직을 구조적으로 어디서 처리하면 좋을 지 고민필요 
            error_response = {
                "message": "error",
                "error": "name must be between 1 and 50 characters"
            }
            return Response(error_response, status=status.HTTP_400_BAD_REQUEST)
        
        result = self.name_repo.create_name(name)
        return Response(result)
