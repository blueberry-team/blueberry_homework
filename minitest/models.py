from django.db import models
import uuid
from django.core.validators import MinLengthValidator, MaxLengthValidator

class Name(models.Model):
    "사용자 모델"
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    email = models.CharField(max_length=50)
    token = models.CharField(max_length=200, blank=True, null=True)
    password = models.CharField(max_length=50)
    name = models.CharField(max_length=50)
    role = models.CharField(max_length=10)
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    def save(self, *args, **kwargs):
        self.clean()
        super().save(*args, **kwargs)

    def __str__(self):
        return self.name

    class Meta:
        db_table = 'users'
        verbose_name = 'User'
        verbose_name_plural = 'Users'


class Company(models.Model):
    """회사 모델"""
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
    # 참조된 Name 객체가 삭제되면 이 회사도 삭제됨 
    user_id = models.ForeignKey(Name, on_delete=models.CASCADE, related_name='companies')
    company_name = models.CharField(max_length=100)
    company_address = models.CharField(max_length=200, blank=True, null=True)
    total_staff = models.IntegerField(default=0, validators=[MinLengthValidator(0), MaxLengthValidator(10000)])
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)

    def save(self, *args, **kwargs):
            self.clean()
            super().save(*args, **kwargs)
            
    def __str__(self):
        return self.company_name

    class Meta:
        db_table = 'companies'
        verbose_name = 'Company'
        verbose_name_plural = 'Companies'