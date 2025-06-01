from .base_repository import BaseRepository
from ..models import Name
from django.core.validators import ValidationError 
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

    def change_name(self, uuid, name):
        if self.find_by_name(name):
            raise ValidationError('A name with the same value already exists')

        try:
            user = self.model.objects.get(id=uuid)
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
        
        user_to_delete = all_users[idx]
        user_to_delete.delete()
        
        return list(self.model.objects.all())
    
    def delete_name(self, name: str):
        print(f"Trying to delete users with name: {name}")
        exists = self.model.objects.filter(name=name).exists()
        print(f"Users with name '{name}' exist: {exists}")
        
        deleted_count = self.model.objects.filter(name=name).delete()
        print(f"Deleted count: {deleted_count}")
        
        return list(self.model.objects.all())