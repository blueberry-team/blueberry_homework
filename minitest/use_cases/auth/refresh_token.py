from django.core.validators import ValidationError
from ...repositores.name_repository import NameRepository
from ...utils.jwt.generate_token import create_jwt_token

def refresh_user_token(user_id):
    """
    사용자 ID로 최신 정보를 조회하고 새 토큰 생성

    """
    name_repo = NameRepository()
    
    try:
        user = name_repo.model.objects.get(id=user_id)
        new_token = create_jwt_token(user)
        
        return {
            "success": True,
            "token": new_token,
            "user": {
                "id": str(user.id),
                "email": user.email,
                "name": user.name
            }
        }
        
    except name_repo.model.DoesNotExist:
        raise ValidationError("사용자를 찾을 수 없습니다.")
    except Exception as e:
        raise ValidationError(f"토큰 갱신 중 오류 발생: {str(e)}")