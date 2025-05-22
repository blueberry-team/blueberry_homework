# HOW TO: 주요 커맨드 정리

## 1. 도커 및 서버 관련 커맨드

### 도커 전체 실행 (ScyllaDB + API 서버)
```bash
docker-compose up --build
```

### 도커 전체 중지
```bash
docker-compose down
```

### requirements.txt 패키지 설치 (로컬 개발용)
```bash
pip install -r requirements.txt
```

### FastAPI 서버 로컬 실행
```bash
uvicorn main:app --host 0.0.0.0 --port 8000 --reload
```

---

## 2. ScyllaDB 내부 데이터 직접 조회

### ScyllaDB 컨테이너 내부에서 cqlsh 실행
```bash
docker exec -it scylla_db cqlsh
```

### 키스페이스 선택
```sql
USE mykeyspace;
```

### users 테이블 전체 조회
```sql
SELECT * FROM users;
```

### companies 테이블 전체 조회
```sql
SELECT * FROM companies;
```

### cqlsh 종료
```sql
exit;
```

---

## 3. 테이블 구조 변경/초기화

### companies 테이블 삭제 및 재생성 예시
```sql
DROP TABLE IF EXISTS companies;
CREATE TABLE companies (
    id UUID PRIMARY KEY,
    name TEXT,
    company_name TEXT,
    created_at TIMESTAMP
);
```

---

이 문서에 없는 커맨드가 필요하면 언제든 추가 요청해 주세요!
