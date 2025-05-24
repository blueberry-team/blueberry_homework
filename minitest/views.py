from rest_framework.views import APIView # type: ignore
from rest_framework.response import Response # type: ignore
from rest_framework import status # type: ignore
from django.core.exceptions import ValidationError # type: ignore
from rest_framework.exceptions import APIException # type: ignore
from .use_cases.auth.sign_up import create_user
from .use_cases.auth.sign_in import login_user
from .use_cases.auth.my_page import get_user, change_user, delete_user_by_index, delete_user_by_name 
from .use_cases.company import get_company, create_company, delete_company_by_index, change_company
import json

def _user_to_dict(user):
    # User 객체를 딕셔너리로 변환하는 헬퍼 메소드
    return {
        'id': user.id,
        'name': user.name,
        'email': user.email,
        'password': user.password,
        'role': user.role,  
        'created_at': user.created_at.strftime("%Y-%m-%d %H:%M:%S") if user.created_at else None,
        'updated_at': user.updated_at.strftime("%Y-%m-%d %H:%M:%S") if user.updated_at else None
    }

def _company_to_dict(company):
    # Company 객체를 딕셔너리로 변환하는 헬퍼 메소드
    return {
        'id': company.id,
        # 'name': company.name,
        'user_id': str(company.user_id.id),
        'company_name': company.company_name,
        'company_address': company.company_address,
        'total_staff': company.total_staff,
        'created_at': company.created_at.strftime("%Y-%m-%d %H:%M:%S") if company.created_at else None,
        'updated_at': company.updated_at.strftime("%Y-%m-%d %H:%M:%S") if company.updated_at else None
    }

class SignUpAPIView(APIView):
    # 회원가입 API
    def post(self, request): 
        name = request.data.get('name')
        email = request.data.get('email')
        password = request.data.get('password')
        role = request.data.get('role')
        # 리포지토리에서 발생시킨 예외를, 뷰에서 적절한 HTTP 응답으로 변환하여 코드 구조를 분리함 
        # (비즈니스 규칙 에러는 리포지토리에서, HTTP 응답 관리는 뷰에서 처리함)
        try:
            users = create_user(name, email, password, role)
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

    # def delete(self, request):
    #     try:
    #         name = request.body.decode('utf-8')
    #         users = delete_user_by_name(name)
    #         return Response({
    #             "message": "success",
    #             "data": [_user_to_dict(user) for user in users]
    #         })
    #     except APIException as e:
    #         return Response({
    #             "message": "error",
    #             "data": str(e)
    #         })

# class NameDeleteAPIView(APIView):

#     def get(self, request, idx):
#         try:
#             users = get_user(idx)
#             return Response({
#                 "message": "success",
#                 "data": [_user_to_dict(user) for user in users]
#             })
#         except ValidationError as e:
#             error_response = {
#                     "message": "error",
#                     "error": e
#                 }
#             return Response(error_response, status=status.HTTP_400_BAD_REQUEST)
        
#     def delete(self, request, idx):
#         try:
#             users = delete_user_by_index(idx)
#             return Response({
#                 "message": "success",
#                 "data": [_user_to_dict(user) for user in users]
#             })
#         except APIException as e:
#             return Response({
#                 "message": "error",
#                 "data": str(e)
#             })
        
class SignInAPIView(APIView):
    def post(self, request): 
        email = request.data.get('email')
        password = request.data.get('password')
        # 리포지토리에서 발생시킨 예외를, 뷰에서 적절한 HTTP 응답으로 변환하여 코드 구조를 분리함 
        # (비즈니스 규칙 에러는 리포지토리에서, HTTP 응답 관리는 뷰에서 처리함)
        try:
            users = login_user(email, password)
            print("[hrkim]",users)
            return Response({
                "message": "success",
                "data": _user_to_dict(users)
            })
        except ValidationError as e:
            error_response = {
                    "message": "error",
                    "error": e
                }
            return Response(error_response, status=status.HTTP_400_BAD_REQUEST)

class MyPageAPIView(APIView):
    # 유저 정보 조회 API
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
        
    # 유저 정보 수정 API
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
                "data": [_user_to_dict(user)]
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
    # 회사 정보 조회 및 생성 API
    def get(self, request):
        try:
            company = get_company()
            return Response({
                "message": "success",
                "data": [_company_to_dict(c) for c in company]
            })
        except APIException as e:
            return Response({
                "message": "error",
                "data": str(e)
            })

    def post(self, request): 
        print("Raw body:", request.body)  # 실제 받은 데이터 출력
        try:
            data = request.data
            
            # name = data.get('name')
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
                "data": [_company_to_dict(c) for c in company]
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
                "data": [_company_to_dict(company)]
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
            companies = get_company(idx)
            return Response({
                "message": "success",
                "data": [_company_to_dict(company) for company in companies]
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
                "data": [_company_to_dict(c) for c in company]
            })
        except APIException as e:
            return Response({
                "message": "error",
                "data": str(e)
            })