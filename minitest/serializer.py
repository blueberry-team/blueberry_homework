from rest_framework import serializers
from .models import Name, Company

class UserSerializer(serializers.ModelSerializer):
    created_at = serializers.DateTimeField(format="%Y-%m-%d %H:%M:%S", read_only=True)
    updated_at = serializers.DateTimeField(format="%Y-%m-%d %H:%M:%S", read_only=True)
    
    class Meta:
        model = Name
        fields = ['id', 'name', 'email', 'token', 'password', 'role', 'created_at', 'updated_at']
        # extra_kwargs = {
        #     'password': {'write_only': True}  # 보안상 password는 응답에서 제외 가능함
        # }

class CompanySerializer(serializers.ModelSerializer):
    user_id = serializers.CharField(source='user_id.id', read_only=True)
    created_at = serializers.DateTimeField(format="%Y-%m-%d %H:%M:%S", read_only=True)
    updated_at = serializers.DateTimeField(format="%Y-%m-%d %H:%M:%S", read_only=True)
    
    class Meta:
        model = Company
        fields = ['id', 'user_id', 'company_name', 'company_address', 'total_staff', 'created_at', 'updated_at']
