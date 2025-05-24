# Blueberry Homework Go Gin - PART.5 MongoDB 통합

🚀 Go Gin 프레임워크와 MongoDB를 사용한 완전한 계층형 아키텍처 애플리케이션

## 📁 프로젝트 구조

```
blueberry_homework_go_gin/
├── app/
│   └── app.go                    # 애플리케이션 초기화 (의존성 주입)
├── config/
│   └── config.go                 # 환경 설정 관리
├── db/
│   └── db.go                     # MongoDB 연결 관리
├── entity/
│   ├── user.go                   # 사용자 엔티티
│   └── company.go                # 회사 엔티티
├── repository/
│   ├── user_repository.go        # 사용자 저장소 (MongoDB)
│   └── company_repository.go     # 회사 저장소 (MongoDB)
├── usecase/
│   ├── user_usecase.go           # 사용자 비즈니스 로직
│   └── company_usecase.go        # 회사 비즈니스 로직
├── handler/
│   ├── user_handler.go           # 사용자 HTTP 핸들러
│   └── company_handler.go        # 회사 HTTP 핸들러
├── .env                          # 환경 변수
├── go.mod                        # Go 모듈 정의
├── main.go                       # 애플리케이션 진입점
├── test_api.sh                   # API 테스트 스크립트
└── README.md                     # 이 파일
```

## 🔧 사전 요구사항

- **Go 1.21 이상**
- **MongoDB** (로컬 설치)

## 📦 MongoDB 설치

### macOS (Homebrew)
```bash
brew tap mongodb/brew
brew install mongodb-community@7.0
```

### Ubuntu/Debian
```bash
wget -qO - https://www.mongodb.org/static/pgp/server-7.0.asc | sudo apt-key add -
echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/7.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-7.0.list
sudo apt-get update
sudo apt-get install -y mongodb-org
```

## 🚀 실행 방법

### 1단계: MongoDB 실행
```bash
# 간단한 방법 (터미널 1)
mkdir -p ~/mongodb-data
mongod --dbpath ~/mongodb-data --port 27017 --bind_ip 127.0.0.1
```

### 2단계: 애플리케이션 실행
```bash
# 터미널 2에서
cd blueberry_homework_go_gin

# 의존성 설치
rm go.sum
go mod download
go mod tidy

# 애플리케이션 실행
go run main.go
```

### 성공 로그 확인
```
🚀 애플리케이션 초기화 시작...
✅ 설정 로드 완료: development 환경
✅ MongoDB 연결 성공: mongodb://localhost:27017/blueberry_homework
✅ 컬렉션 초기화 완료
✅ 라우터 설정 완료
✅ 애플리케이션 초기화 완료
🌐 서버 시작: http://localhost:8080
```

## 🧪 API 테스트

### 자동 테스트 스크립트 실행
```bash
chmod +x test_api.sh
./test_api.sh
```

### 수동 테스트

#### 1. 헬스체크
```bash
curl http://localhost:8080/health
```

#### 2. 사용자 생성
```bash
curl -X POST http://localhost:8080/create-name \
  -H "Content-Type: application/json" \
  -d '{"name": "Kim"}'
```

#### 3. 사용자 목록 조회
```bash
curl http://localhost:8080/get-names
```

#### 4. 회사 생성
```bash
curl -X POST http://localhost:8080/create-company \
  -H "Content-Type: application/json" \
  -d '{"name": "Kim", "company_name": "ABC Corp"}'
```

## 📊 MongoDB 데이터 확인

```bash
# MongoDB Shell 접속
mongosh

# 데이터베이스 선택
use blueberry_homework

# 사용자 컬렉션 조회
db.users.find().pretty()

# 회사 컬렉션 조회
db.companies.find().pretty()

# 컬렉션 통계
db.users.countDocuments()
db.companies.countDocuments()
```

## 🔌 API 엔드포인트

| 메서드 | 경로 | 설명 | 요청 본문 | 응답 형식 |
|--------|------|------|-----------|-----------|
| GET | `/health` | 서버 상태 확인 | - | `{message, status, database}` |
| POST | `/create-name` | 사용자 생성 | `{name}` | `{message}` |
| GET | `/get-names` | 사용자 목록 조회 | - | `{message, data}` |
| PUT | `/change-name` | 사용자 이름 변경 | `{id, name}` | `{message}` |
| DELETE | `/delete-index?index=N` | 인덱스로 사용자 삭제 | - | `{message}` |
| DELETE | `/delete-name` | 이름으로 사용자 삭제 | `{name}` | `{message}` |
| POST | `/create-company` | 회사 생성 | `{name, company_name}` | `{message}` |
| GET | `/get-companies` | 회사 목록 조회 | - | `{message, data}` |

## 💡 주요 특징

### ✅ 완전한 계층형 아키텍처
- **Handler** → **UseCase** → **Repository** → **MongoDB**
- 관심사 분리 및 의존성 주입

### ✅ 실제 데이터베이스 통합
- MongoDB를 통한 영구 데이터 저장
- BSON 태그를 통한 자동 직렬화/역직렬화

### ✅ 비즈니스 로직 구현
- 중복 이름 방지
- 사용자당 하나의 회사만 허용
- 완전한 에러 처리

### ✅ 환경 설정 관리
- `.env` 파일을 통한 설정 관리
- 개발/프로덕션 환경 분리

### ✅ 한 줄 초기화
```go
app, err := app.Init() // 모든 의존성이 자동으로 주입됩니다
```

## 🔍 문제 해결

### MongoDB 연결 실패
```bash
# MongoDB 프로세스 확인
ps aux | grep mongod

# 포트 확인
lsof -i :27017

# MongoDB 재시작
mongod --dbpath ~/mongodb-data --port 27017 --bind_ip 127.0.0.1
```

### Go 모듈 문제
```bash
# 모듈 캐시 정리
go clean -modcache
rm go.sum
go mod download
go mod tidy
```

## 🎯 PART.5에서 달성한 목표

- ✅ **MongoDB 실제 데이터베이스 통합**
- ✅ **계층형 아키텍처 유지** (Repository 계층만 수정)
- ✅ **환경 설정 파일** (config, .env)
- ✅ **애플리케이션 초기화** (app.go)
- ✅ **데이터베이스 초기화** (db.go)
- ✅ **한 줄 초기화** (app.Init())
- ✅ **모든 API 기능 정상 동작**
- ✅ **데이터 영속성 보장**

이제 데이터가 메모리가 아닌 실제 MongoDB에 저장되어 애플리케이션을 재시작해도 데이터가 유지됩니다! 🎉
