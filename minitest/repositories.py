from django.db import models
from .tmp_database import tmp_db
from .models import Name

class BaseRepository: 
    def __init__(self, model: models.Model):
        self.model = model

class NameRepository(BaseRepository):
    def __init__(self):
        super().__init__(Name)

    def get_name(self):
        return {
            "message": "success",
            "data": [{'name': name} for name in tmp_db]
        }
        # return self.model.objects.all()
    
    def create_name(self, name: str):
        tmp_db.append(name)
        return {
            "message": "success",
            "data": [{'name': name} for name in tmp_db]
        }
        # self.model.objects.create(name=name) 