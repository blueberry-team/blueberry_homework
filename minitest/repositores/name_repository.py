from .base_repository import BaseRepository
from ..tmp_database import tmp_db
from ..models import Name
from django.core.validators import ValidationError # type: ignore

class NameRepository(BaseRepository):
    def __init__(self):
        super().__init__(Name)

    def get_name(self, idx=None):
        if idx is not None:
            idx = int(idx)
            if idx >= len(tmp_db):
                raise ValidationError('해당 인덱스에 값이 없습니다')
            return [tmp_db[idx]]
        return tmp_db
        # return self.model.objects.all()
    
    def create_name(self, user):           
        tmp_db.append(user)    
        return tmp_db
        # self.model.objects.create(name=name) 

    def delete_name(self, idx: int):
        if idx > len(tmp_db):
            raise ValidationError('해당 인덱스에 값이 없습니다')
        tmp_db.pop(idx)
        return tmp_db