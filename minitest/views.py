from rest_framework.views import APIView 
from rest_framework.response import Response 
from rest_framework import status 
from django.core.exceptions import ValidationError 
from rest_framework.exceptions import APIException 
from .serializer import UserSerializer, CompanySerializer
from .use_cases.auth.sign_up import create_user
from .use_cases.auth.sign_in import login_user
from .use_cases.auth.my_page import get_user, change_user
from .use_cases.company import get_company, create_company, delete_company_by_index, change_company
import json

class SignUpAPIView(APIView):
    def post(self, request): 
        name = request.data.get('name')
        email = request.data.get('email')
        password = request.data.get('password')
        role = request.data.get('role')
        try:
            users = create_user(name, email, password, role)
            return Response({
                "message": "success",
                "data": UserSerializer(users, many=True).data
            })
        except ValidationError as e:
            error_response = {
                    "message": "error",
                    "error": e
                }
            return Response(error_response, status=status.HTTP_400_BAD_REQUEST)

class SignInAPIView(APIView):
    '''
    {
    "email":"test1@test.com",
    "password":"test1!"
    }
    '''
    def post(self, request): 
        email = request.data.get('email')
        password = request.data.get('password')
        try:
            users = login_user(email, password)

            return Response({
                "message": "success",
                "data": UserSerializer(users).data
            })
        except ValidationError as e:
            error_response = {
                    "message": "error",
                    "error": e
                }
            return Response(error_response, status=status.HTTP_400_BAD_REQUEST)

class MyPageAPIView(APIView):
    def get(self, request):
        try:
            users = get_user()
            return Response({
                "message": "success",
                "data": UserSerializer(users, many=True).data
            })
        except APIException as e:
            return Response({
                "message": "error",
                "data": str(e)
            })
        
    def put(self, request):
        try:
            data = json.loads(request.body)
            
            uuid = data.get('id')
            name = data.get('name')

            if not uuid or not name:
                error_response = {
                    "message": "error",
                    "error": "필드가 누락되었습니다"
                }
                return Response(error_response, status=status.HTTP_400_BAD_REQUEST)

            user = change_user(uuid, name)
            
            return Response({
                "message": "success",
                "data": UserSerializer(user, many=True).data
            })
        
        except json.JSONDecodeError:
            error_response = {
                    "message": "error",
                    "error": "입력 형식을 맞춰주세요."
                }
            return Response(error_response, status=status.HTTP_400_BAD_REQUEST)
        except ValidationError as e:
            error_response = {
                    "message": "error",
                    "error": e
                }
            return Response(error_response, status=status.HTTP_400_BAD_REQUEST)

class CompanyAPIView(APIView):
    def get(self, request):
        try:
            company = get_company()
            return Response({
                "message": "success",
                "data": CompanySerializer(company, many=True).data
            })
        except APIException as e:
            return Response({
                "message": "error",
                "data": str(e)
            })

    def post(self, request): 
        print("Raw body:", request.body)  
        try:
            data = request.data            
            user_id = data.get('user_id')
            company_name = data.get('company_name')
            company_address = data.get('company_address')
            total_staff = data.get('total_staff')
            
            if not user_id or not company_name or not company_address or total_staff is None:
                error_response = {
                    "message": "error",
                    "error": "필드가 누락되었습니다"
                }
                return Response(error_response, status=status.HTTP_400_BAD_REQUEST)

            company = create_company(user_id=user_id,
                                    company_name=company_name,
                                    company_address=company_address,
                                    total_staff=total_staff)
            return Response({
                "message": "success",
                "data": CompanySerializer(company, many=True).data
            })
        except ValidationError as e:
            error_response = {
                    "message": "error",
                    "error": e
                }
            return Response(error_response, status=status.HTTP_400_BAD_REQUEST)
        
    def put(self, request):
        try:
            data = json.loads(request.body)
            
            uuid = data.get('id')
            company_name = data.get('company_name')
            company_address = data.get('company_address')
            total_staff = data.get('total_staff')

            if not uuid or not company_name or not company_address or total_staff is None:  
                error_response = {
                    "message": "error",
                    "error": "필드가 누락되었습니다"
                }
                return Response(error_response, status=status.HTTP_400_BAD_REQUEST)

            company = change_company(uuid, 
                            company_name=company_name, 
                            company_address=company_address, 
                            total_staff=total_staff)
            
            return Response({
                "message": "success",
                "data": CompanySerializer(company).data
            })
        
        except json.JSONDecodeError:
            error_response = {
                    "message": "error",
                    "error": "입력 형식을 맞춰주세요."
                }
            return Response(error_response, status=status.HTTP_400_BAD_REQUEST)
        except ValidationError as e:
            error_response = {
                    "message": "error",
                    "error": e
                }
            return Response(error_response, status=status.HTTP_400_BAD_REQUEST)

class CompanyDeleteAPIView(APIView):
    def get(self, request, idx):
        try:
            company = get_company(idx)
            return Response({
                "message": "success",
                "data": CompanySerializer(company, many=True).data
            })
        except ValidationError as e:
            error_response = {
                    "message": "error",
                    "error": e
                }
            return Response(error_response, status=status.HTTP_400_BAD_REQUEST)

    def delete(self, request, idx):
        try:
            company = delete_company_by_index(idx)
            return Response({
                "message": "success",
                "data": CompanySerializer(company, many=True).data
            })
        except APIException as e:
            return Response({
                "message": "error",
                "data": str(e)
            })