from ...domain.user import User
from ...repositores.name_repository import NameRepository
from datetime import datetime
from django.core.validators import ValidationError 
import uuid
import bcrypt

name_repo = NameRepository()

def hash_password(password):
    salt = bcrypt.gensalt()
    hashed_password = bcrypt.hashpw(password.encode('utf-8'), salt)
    return hashed_password.decode('utf-8')

def create_user(name, email, password, role):
    hashed_password = hash_password(password)
    user = User(
        name=name, 
        id=uuid.uuid4(), 
        email=email, 
        password=hashed_password, 
        role=role, 
        created_at=datetime.now(), 
        updated_at=None
    )
    try:
        user.validate() 
    except ValueError as e:
        raise ValidationError(str(e))
    users = name_repo.create_name(user)
    return users