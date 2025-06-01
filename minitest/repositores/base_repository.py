from django.db import models

class BaseRepository: 
    def __init__(self, model: models.Model):
        self.model = model
