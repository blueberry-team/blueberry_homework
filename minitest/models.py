from django.db import models

# 이름 입력, 조회 모델
class Name(models.Model):
    name = models.CharField(max_length=255)

    def __str__(self):
        return self.name