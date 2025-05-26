from .base_repository import BaseRepository
from ..tmp_database import tmp_user_db
from ..models import Name
from django.core.validators import ValidationError # type: ignore

class NameRepository(BaseRepository):
    def __init__(self):
        super().__init__(Name)

    def get_name(self, idx=None):
        exist_all_users = list(self.model.objects.all())
        if idx is not None:
            idx = int(idx)
            if idx >= len(exist_all_users):
                raise ValidationError('해당 인덱스에 값이 없습니다')
            return [exist_all_users[idx]]
        return list(self.model.objects.all())
    
    def get_user_by_uuid(self, uuid):
        exist_all_users = list(self.model.objects.all())
        for user in exist_all_users:
            if str(user.id) == uuid:
                return user
        raise ValidationError('해당 UUID의 사용자가 존재하지 않습니다')
    
    def create_name(self, user):    
        if self.find_by_name(user.name):
            raise ValidationError('A name with the same value already exists')
        # tmp_user_db.append(user)  

        # Django 
        user = self.model(
            id=user.id,
            name=user.name,
            email=user.email,
            password=user.password,
            role=user.role,
            created_at=user.created_at,
            updated_at=user.updated_at
        )
        user.save()  
        return list(self.model.objects.all())
        # self.model.objects.create(name=name) 

    def change_name(self, uuid, name):
        if self.find_by_name(name):
            raise ValidationError('A name with the same value already exists')

        try:
            # UUID로 사용자 찾기
            user = self.model.objects.get(id=uuid)
            # 이름 변경 및 저장
            user.name = name
            user.save()
            print('change name: ', user)
            return user
        except self.model.DoesNotExist:
            raise ValidationError('A user does not exist')

    def find_by_name(self, name: str):
        return self.model.objects.filter(name=name).exists()
    

    def delete_index(self, idx: int):
        all_users = list(self.model.objects.all())
        if idx >= len(all_users):
            raise ValidationError('해당 인덱스에 값이 없습니다')
        
        # 해당 인덱스의 사용자 삭제
        user_to_delete = all_users[idx]
        user_to_delete.delete()
        
        # 삭제 후 MongoDB의 모든 사용자 목록 반환
        return list(self.model.objects.all())
    
    def delete_name(self, name: str):
        print(f"Trying to delete users with name: {name}")
        # 삭제 전 해당 이름의 사용자가 존재하는지 확인
        exists = self.model.objects.filter(name=name).exists()
        print(f"Users with name '{name}' exist: {exists}")
        
        # 삭제 작업 수행
        deleted_count = self.model.objects.filter(name=name).delete()
        print(f"Deleted count: {deleted_count}")
        
        # 삭제 후 모든 사용자 목록 반환
        return list(self.model.objects.all())