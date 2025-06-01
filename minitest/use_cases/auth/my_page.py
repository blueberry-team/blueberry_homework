from ...repositores.name_repository import NameRepository
from datetime import datetime
name_repo = NameRepository()

def get_user(idx=None):
    users = name_repo.get_name(idx)
    return users

def get_user_by_uuid(uuid):
    user = name_repo.get_user_by_uuid(uuid)
    return user

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