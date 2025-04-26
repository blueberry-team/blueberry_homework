from ..domain.user import User
from ..repositores.name_repository import NameRepository
from datetime import datetime
from django.core.validators import ValidationError # type: ignore
import uuid

name_repo = NameRepository()

def get_user(idx=None):
    users = name_repo.get_name(idx)
    return users

def create_user(name):
    user = User(name=name, id=uuid.uuid4(), created_at=datetime.now(), updated_at=None)
    try:
        user.validate() 
    except ValueError as e:
        raise ValidationError(str(e))
    users = name_repo.create_name(user)
    return users

def change_user(uuid, name):
    user = name_repo.change_name(uuid, name)
    user.updated_at = datetime.now()
    return user

def delete_user_by_index(idx):
    users = name_repo.delete_index(idx)
    return users

def delete_user_by_name(name):
    users = name_repo.delete_name(name)
    return users