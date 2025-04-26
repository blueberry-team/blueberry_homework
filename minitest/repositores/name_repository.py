from .base_repository import BaseRepository
from ..tmp_database import tmp_user_db
from ..models import Name
from django.core.validators import ValidationError # type: ignore

class NameRepository(BaseRepository):
    def __init__(self):
        super().__init__(Name)

    def get_name(self, idx=None):
        if idx is not None:
            idx = int(idx)
            if idx >= len(tmp_user_db):
                raise ValidationError('해당 인덱스에 값이 없습니다')
            return [tmp_user_db[idx]]
        return tmp_user_db
        # return self.model.objects.all()
    
    def create_name(self, user):    
        if self.find_by_name(user.name):
            raise ValidationError('A name with the same value already exists')
        tmp_user_db.append(user)    
        return tmp_user_db
        # self.model.objects.create(name=name) 

    def change_name(self, uuid, name):
        if self.find_by_name(name):
            raise ValidationError('A name with the same value already exists')

        for exist_user in tmp_user_db:
            if str(exist_user.id) == uuid:
                exist_user.name = name
                print('change name: ', exist_user)
                return exist_user 
        raise ValidationError('A user does not exist')

    def find_by_name(self, name: str):
        for user in tmp_user_db:
            if user.name==name:
                return True
        return False
    

    def delete_index(self, idx: int):
        if idx > len(tmp_user_db):
            raise ValidationError('해당 인덱스에 값이 없습니다')
        tmp_user_db.pop(idx)
        return tmp_user_db
    
    def delete_name(self, name: str):
        tmp_user_db[:] = [user for user in tmp_user_db if user.name != name]
        return tmp_user_db