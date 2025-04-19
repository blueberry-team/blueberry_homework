from django.db import models
from django.core.validators import MinLengthValidator, MaxLengthValidator

# 이름 입력, 조회 모델
class Name(models.Model):
    # validator들은 실제 데이터베이스 모델과 연동되어 있을 때 자동으로 적용되며, 지금은 임시 리스트 사용중이므로 실행되지 않음
    name = models.TextField(validators=[MinLengthValidator(1, '1자 이상 적어주세요'), MaxLengthValidator(50, '50자 이하 적어주세요')])
    # datetime
    
    def __str__(self):
        return self.name