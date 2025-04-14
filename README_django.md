### Django 프로젝트 생성 후 서버 동작
```
django-admin startproject mysite .
python manage.py runserver // default port: 8000
```

### 어플리케이션 (미니 테스트) 생성
```
python manage.py startapp minitest
```

### 미니 테스트 API
- minitest/views.py: 비즈니스 로직
- minitest/repositories.py: 데이터 접근 로직
- minitest/models.py: 데이터 모델 (데이터베이스 활용할 경우, 효과 극대화)
- minitest/tmp_database.py: 임시 데이터베이스

### 미니 테스트 API 동작 캡쳐
- POST
![스크린샷 2025-04-05 오후 5 39 30](https://github.com/user-attachments/assets/3fa295ff-51b2-4bc7-b803-305d1a9ee73f)

- GET
![스크린샷 2025-04-05 오후 5 40 15](https://github.com/user-attachments/assets/4248b0a9-2b4a-49ac-9291-3d29a7136d99)

