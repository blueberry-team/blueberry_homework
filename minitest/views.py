from rest_framework.views import APIView
from .repositories import NameRepository
from rest_framework.response import Response
from rest_framework import status


class NameAPIView(APIView):
    name_repo = NameRepository()

    def get(self, request):
        name = self.name_repo.get_name()
        return Response(name)

    def post(self, request): 
        name = request.data.get('name')
        self.name_repo.create_name(name)
        return Response(name, status=status.HTTP_201_CREATED)
 