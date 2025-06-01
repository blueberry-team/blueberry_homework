# REST API 미니 테스트 PART.1

이 리포지토리는 다양한 프레임워크를 사용한 간단한 REST API 구현을 위한 미니 테스트 프로젝트입니다.

## 프레임워크

각 참여자는 다음 프레임워크를 사용합니다:

- 해린 - Python (Django)
- 정우 - Python (FastAPI)
- 상화 - Go (Chi)
- 승재 - Go (Gin)
- 한규 - .NET
- 재혁 - Rust (Axum)
- 상일 - Kotlin + Spring

## 규칙

### 사용 가능한 함수명
- `createName`
- `getName`

### 제약 사항
- DB 사용 금지
- 데이터는 빈 배열(`[]`)에 저장하여 관리
- Repository 구조를 사용하여 구현
- Domain 생성 금지

## 구현 목표

간단한 테스트 구조를 생성하는 것이 목표이며, 각 프레임워크에서 동일한 기능을 구현하는 방식을 비교합니다.

## 추가 구현 사항

- 시간이 여유롭다면 유효성 검사(Validation)와 같은 기능들을 추가 구현하셔도 좋습니다.
- 프론트엔드는 원하는 웹 애플리케이션을 사용하셔도 좋고, curl을 통한 API 테스트도 가능합니다.

---

# REST API 미니 테스트 PART.2

## 데이터 구조 변경

- 이전에는 빈 배열(`[]`)에 데이터를 저장하였으나, PART.2에서는 Map 형태로 저장합니다.
- 배열 안에 저장될 내용은 `[{name: "NAME"}, {name: "KING"}]` 형식의 객체 배열로 변경합니다.
- GET 요청 응답도 `[{name: "KING"}]` 형태로 반환되어야 합니다.

## 응답 형식

- API 요청 유형에 따라 응답 형식이 달라집니다:

### POST 요청 응답 (생성/수정/삭제)
- 성공 시 `message` 필드만 포함됩니다:
  ```json
  {
    "message": "success"
  }
  ```

### GET 요청 응답 (조회)
- 성공 시 `message`와 `data` 필드가 모두 포함됩니다:
  ```json
  {
    "message": "success",
    "data": [/* 결과 데이터 */]
  }
  ```
- 예시:
  ```json
  {
    "message": "success",
    "data": [
      {
        "id": "550e8400-e29b-41d4-a716-446655440000",
        "name": "NAME",
        "createdAt": "2023-05-15T14:30:45Z",
        "updatedAt": "2023-05-15T14:30:45Z"
      }
    ]
  }
  ```

### 실패 응답
- 모든 요청 유형에서 실패 시 동일한 형식으로 응답합니다:
  ```json
  {
    "message": "error",
    "error": "오류 메시지"
  }
  ```

## 유효성 검증 (Validation)

- 모든 요청에서 `name` 값에 대한 유효성 검증을 수행해야 합니다.
- `name` 값은 최소 1자 이상, 최대 50자 이하여야 합니다.
- 유효성 검증에 실패할 경우 적절한 오류 메시지를 반환해야 합니다:
  ```json
  {
    "message": "error",
    "error": "name must be between 1 and 50 characters"
  }
  ```

## 추가 기능 구현

### DELETE 요청 구현
- Postman을 사용하여 삭제할 항목의 인덱스를 요청합니다.
- Swagger이 편하신 분들은 사용하셔도 됩니다.
- 해당 인덱스의 항목이 삭제된 후 배열이 재정렬되어야 합니다.
- 예: 인덱스 1을 삭제 요청하면 해당 위치의 항목이 제거되고 뒤의 항목들이 앞으로 이동해야 합니다.

### 사용 가능한 함수명
- `deleteName` - 특정 인덱스의 이름 항목을 삭제하는 기능

## PR 제출 시 요구사항

- PR을 올릴 때 모든 기능이 정상 동작하는지 확인할 수 있도록 Postman이나 Swagger에서 모든 API 응답 구조를 캡처하여 첨부해야 합니다.
- 다음 항목에 대한 캡처를 포함해야 합니다:
  - 데이터 생성(Create) 응답
  - 데이터 조회(GET) 응답
  - 데이터 삭제(DELETE) 응답
  - 유효성 검증 실패 시 오류 응답
- 캡처 이미지는 PR 설명에 첨부하여 제출합니다.

---

# REST API 미니 테스트 PART.3

## 아키텍처 변경

PART.3에서는 계층형 아키텍처를 적용하여 코드를 구조화합니다:

```
Handler -> UseCase -> Repository
```

## Entity와 Domain 생성

- 이제 Entity와 Domain을 생성해야 합니다.
- UserEntity 구조:
  ```
  UserEntity {
    name: String,
    createdAt: Time
  }
  ```

## UseCase 계층 추가

- UseCase는 Repository로 가기 전에 Entity를 처리합니다.
- UseCase에서 현재 시간을 기준으로 createdAt 필드를 설정해야 합니다.
- 기존 함수 중 Domain을 거치지 않아도 되는 로직들은:
  - 일관성을 위해 모든 경로에서 UseCase를 거치도록 구현하거나
  - 필요한 부분만 연결해도 됩니다.

## 함수 변경 및 추가

- 기존의 `deleteName` 함수는 `deleteIndex`로 이름을 변경합니다.
- 새로운 `deleteName` 함수를 추가합니다:
  - 이름을 입력받아 동일한 이름을 가진 항목을 찾아 삭제하는 기능
  - 여러 항목이 있을 경우 모두 삭제해야 합니다.
- GET 요청으로 이름을 가져올 때는 생성된 시간(createdAt)도 함께 반환해야 합니다.

## 응답 형식

- 모든 API 응답은 PART.2와 동일한 형식을 유지합니다.
- 성공 시:
  ```json
  {
    "message": "success",
    "data": [/* 결과 데이터 */]
  }
  ```
- GET 요청 응답 예시:
  ```json
  {
    "message": "success",
    "data": [
      {
        "name": "NAME",
        "createdAt": "2023-05-15T14:30:45Z"
      },
      {
        "name": "KING",
        "createdAt": "2023-05-15T15:20:10Z"
      }
    ]
  }
  ```
- 실패 시:
  ```json
  {
    "message": "error",
    "error": "오류 메시지"
  }
  ```

## PR 제출 시 요구사항

- PR을 올릴 때 모든 기능이 정상 동작하는지 확인할 수 있도록 Postman이나 Swagger에서 모든 API 응답 구조를 캡처하여 첨부해야 합니다.
- 다음 항목에 대한 캡처를 포함해야 합니다:
  - 데이터 생성(Create) 응답
  - 데이터 조회(GET) 응답
  - 인덱스로 데이터 삭제(deleteIndex) 응답
  - 이름으로 데이터 삭제(deleteName) 응답
  - 유효성 검증 실패 시 오류 응답
- 캡처 이미지는 PR 설명에 첨부하여 제출합니다.

---

# REST API 미니 테스트 PART.4

## UserEntity 확장

PART.4에서는 UserEntity에 다음 필드를 추가합니다:
  ```
  UserEntity {
    id: String/UUID,
    name: String,
    createdAt: Time,
    updatedAt: Time
  }
  ```

## 중복 이름 제약 추가

- `createName` 함수에서는 중복된 이름이 없는지 반드시 체크해야 합니다.
- UserEntity가 확장됨에 따라 이후 `createCompany` 함수까지 연계되기 때문에 유저는 더 이상 중복된 이름을 가질 수 없습니다.
- 중복된 이름으로 생성 시도 시 다음과 같은 오류를 반환해야 합니다:
  ```json
  {
    "message": "error",
    "error": "A name with the same value already exists"
  }
  ```

## 새로운 함수 추가

### UserRepository 및 UseCase에 추가되는 함수

1. **changeName**
   - 기능: 사용자 이름을 변경합니다.
   - 검색 방식: create에서 받은 유저의 UUID를 통해 검색한 후 이름을 변경해야 합니다.
   - 시간 제약: 이름이 변경될 때 `updatedAt` 시간이 현재 시간으로 업데이트되어야 하며, 반드시 `createdAt`과 `updatedAt` 시간이 달라야 합니다.
   - 중복 확인:
     - 자기 자신: 변경하려는 이름이 기존 이름과 동일하다면(변경되지 않았다면) 다음과 같은 오류를 반환해야 합니다:
       ```json
       {
         "message": "error",
         "error": "A name with the same value already exists."
       }
       ```
     - 다른 유저와의 중복: 변경하려는 이름이 이미 다른 유저가 사용 중인 이름이라면 다음과 같은 오류를 반환해야 합니다:
       ```json
       {
         "message": "error",
         "error": "A name with the same value already exists."
       }
       ```
   - API 경로: 적절한 라우팅 설정 필요 (PUT 또는 PATCH 메서드 권장)
   - 요청 형식 예시:
     ```json
     {
       "id": "550e8400-e29b-41d4-a716-446655440000",
       "name": "NEW_NAME"
     }
     ```

2. **findByName**
   - 기능: 이름을 기준으로 사용자가 존재하는지 확인합니다.
   - 위치: `user_repository`에 구현해야 합니다.
   - 참고: 별도의 라우터 경로는 필요하지 않으며, 내부 로직에서 사용됩니다.

## Company Entity 추가

새로운 Company 엔티티를 생성합니다:
  ```
  CompanyEntity {
    id: String/UUID,
    name: String,
    company_name: String,
    created_at: Time
  }
  ```

## Company 관련 구성 요소

1. **CompanyRepository**
   - 새로운 빈 배열 형태의 DB를 추가하여 회사 데이터를 저장합니다.

2. **CompanyUseCase**
   - **createCompany**
     - 기능: 새로운 회사를 생성합니다.
     - 처리 과정: UserRepository의 `findByName`을 먼저 호출하여 해당 사용자가 존재하는지 확인 후, 조건에 맞는 경우에만 회사를 생성합니다.
     - 제약 조건: 이미 회사를 가지고 있는 사용자는 새로운 회사를 생성할 수 없습니다.
     - 중복 확인: 동일한 사용자 이름으로 회사를 중복 생성하려는 경우 아래와 같은 오류 메시지를 반환해야 합니다:
      ```json
        {
          "message": "error",
          "error": "user already has a company"
        }
        ```
     - 요청 형식예시:
      ```json
        {
          "message": "success",
          "data": [
            {
              "id": "550e8400-e29b-41d4-a716-446655440000",
              "name": "NAME",
              "company_name": "COMPANY",
              "created_at": "2023-05-15T14:30:45Z"
            }
          ]
        }

   - **getCompany**
     - 기능: 회사 정보를 조회합니다.


## 응답 형식

- API 요청 유형에 따라 응답 형식이 달라집니다:

### POST 요청 응답 (생성/수정/삭제)
- 성공 시 `message` 필드만 포함됩니다:
  ```json
  {
    "message": "success"
  }
  ```

### GET 요청 응답 (조회)
- 성공 시 `message`와 `data` 필드가 모두 포함됩니다:
  ```json
  {
    "message": "success",
    "data": [/* 결과 데이터 */]
  }
  ```
- 예시:
  ```

---

# REST API 미니 테스트 PART.5

## 데이터베이스 통합

PART.5에서는 실제 데이터베이스를 추가하여 임시 배열 저장소를 대체합니다:

- 데이터베이스 선택: NoSQL 데이터베이스인 Scylla DB 또는 MongoDB 중 선택하여 사용합니다.
  - MongoDB를 선택하는 경우 스키마를 직접 생성해야 합니다.
  - Scylla DB를 선택하는 경우 CQL 문법을 학습하여 사용해야 합니다.

## 구조 변경

다음과 같은 파일 구조를 추가하여 데이터베이스 연결을 관리합니다:

1. **app 파일 생성**
   - `app.py` / `app.go` 등 언어에 맞는 애플리케이션 초기화 파일을 생성합니다.
   - 필요한 모든 의존성을 이 파일에서 초기화합니다.
   - `app.init()` 형태의 함수를 제공하여 프레임워크 빌드 시 한 줄의 코드로 초기화할 수 있도록 합니다.

2. **환경 설정 관리**
   - `config` 파일을 생성하여 환경 변수를 초기화합니다.
   - `.env` 파일을 사용하여 데이터베이스 연결 정보 등의 설정을 관리합니다.

3. **데이터베이스 초기화**
   - `db.go` / `db.py` 등의 파일을 생성하여 데이터베이스 초기화 로직을 구현합니다.
   - 데이터베이스 연결, 테이블/컬렉션 생성 등의 로직을 포함해야 합니다.

## 구현 요구사항

- 모든 데이터는 이제 실제 데이터베이스에 저장되어야 합니다.
- Repository 계층만 수정하여 실제 데이터베이스와 상호작용하도록 변경합니다.
- Domain, UseCase, Handler 계층은 최소한의 변경만 허용됩니다.
- `main` 함수에서는 데이터베이스 초기화 함수를 호출하는 코드만 추가합니다.
- PART5에선 POSTMAN이 아닌 Database에 내가 추가한 데이터가 기록되어 있는지 보여주시면 됩니다.

## PR 제출 시 요구사항

- PR을 올릴 때 모든 기능이 정상 동작하는지 확인할 수 있도록 Postman이나 Swagger에서 모든 API 응답 구조를 캡처하여 첨부해야 합니다.
- 다음 항목에 대한 캡처를 포함해야 합니다:
  - 데이터베이스 연결 성공 로그
  - 모든 API 기능의 정상 동작 확인 캡처
- 데이터베이스 설정 방법 및 실행 방법에 대한 간단한 문서를 PR 설명에 포함해야 합니다.

---

# REST API 미니 테스트 PART.6

## 필수 항목 추가
- 자신이 구성한 아키텍처의 구조를 Excalidraw로 그려주세요
- 외부 툴을 사용하셔도 좋습니다.

## 인증(Auth) 시스템 도입

- UserEntity는 변경되어야 합니다.
- 다음과 같은 인증 기능이 추가됩니다:
  - 회원가입 `sign-up`
  - 로그인 `log-in`
  - 유저정보 수정 `change-user`
  - 유저정보 획득 `get-user`

## 회원가입 요구사항

- 회원가입 시 필요한 정보:
  - email
  - password
  - name
  - role
  - createdAt
  - updatedAt

- role은 다음 두 가지 유형만 가능합니다:
  - boss
  - worker

## 보안 요구사항

- 비밀번호는 반드시 해싱 처리 및 난독화 처리가 필요합니다.
- Auth 모듈 내에 `findById` 기능이 구현되어야 합니다.

## Company 기능 확장

- Company 기능에 다음 작업이 추가됩니다:
  - create (생성)
  - get (조회)
  - change (수정)
  - delete (삭제)

- 회사 생성(create) 제한:
  - 'boss' role을 가진 사용자만 회사를 등록할 수 있습니다.
  - 회사 정보에는 userId가 중복으로 포함되어야 합니다.

  - 회사를 생성 시 필요한 정보:
  - userId
  - companyName
  - companyAddress
  - totalStaff
  - createdAt
  - updatedAt

- 선택 사항:
  - 색인을 위한 `findByCompany`와 같은 로직 추가

## 데이터 저장

- 모든 정보는 데이터베이스에 저장되어야 합니다.

## PR 제출 시 요구사항r

- Postman 대신 다음 정보를 첨부해야 합니다:
  - 'boss' role을 가진 사용자 생성 내역
  - 'worker' role을 가진 사용자 생성 내역
  - 생성된 회사 정보

---

# REST API 미니 테스트 PART.7

## JWT 인증 시스템 도입

PART.7에서는 JWT(JSON Web Token) 기반 인증 시스템을 구현합니다.

## JWT 토큰 발급

- 로그인 시 만료시간 5시간짜리 JWT 토큰을 발급해야 합니다.
- 토큰 생성은 `log_in_usecase`에서 이루어져야 합니다.
- 토큰 인코딩 알고리즘은 반드시 `HS512`를 사용해 이루어져야 합니다.
- 토큰 디코딩 시 디코딩 알고리즘도 반드시 `HS512`만으로 해제할 수 있어야 합니다.

## JWT 토큰 구조

JWT 토큰에는 다음 정보가 포함되어야 합니다:
```json
{
  "sub": "user_id",
  "email": "user_email",
  "name": "user_name",
  "exp": "expiration_time",
  "iat": "issued_at_time"
}
```

## JWT 유틸리티 함수

다음 파일들을 생성하여 JWT 관련 기능을 구현해야 합니다:

1. **utils/jwt/generate_token**
   - JWT 토큰 생성, 갱신 기능을 구현합니다.

2. **middleware/verify_token**
   - JWT 토큰 검증 기능을 구현합니다.
   - 각자 프레임워크 별로 middleware 사용법이 있습니다.

## 토큰 갱신 시스템

### 토큰 만료 시간 체크
- 모든 API 요청에서 토큰의 만료 시간을 확인해야 합니다.
- 요청 토큰의 만료시간이 1시간 이내로 남았다면 응답 헤더에 `"X-Token-Refresh-Required"`를 포함하여 전송해야 합니다.

### 토큰 갱신 전용 라우터
- 토큰 갱신을 위한 별도의 라우터를 생성해야 합니다.
- ex: `/api/auth/refresh-token`
- 갱신 요청은 `refresh_usecase`에서 처리되어야 합니다.

### 토큰 갱신 프로세스
- 토큰 갱신 시 원본 데이터의 안전성을 보장할 수 없으므로, 토큰을 만드는데 필요한 유저 정보를 repository에서 다시 획득해야 합니다.
- `refresh_usecase`에서 다음 과정을 거쳐야 합니다:
  1. 기존 토큰에서 user_id 추출
  2. Repository에서 최신 유저 정보 조회
  3. 조회된 최신 정보로 새로운 JWT 토큰 생성 (generate_token 함수 사용)
  4. 갱신된 토큰 반환

## 환경 설정

- `secret_key`는 반드시 `.env` 파일에 생성하고 환경 변수에서 불러와서 사용해야 합니다.
- 예시:
  ```
  JWT_SECRET_KEY=your_secret_key_here
  ```

## 인증이 필요한 API

- `sign_up`과 `login`을 제외한 모든 API 요청에는 JWT 토큰이 필요합니다.
- 요청을 처리하기 위한 사용자 정보는 토큰에서 추출하여 사용해야 합니다.
- Handler에서 토큰 검증 및 갱신 처리가 이루어져야 합니다.

## 구현 요구사항

1. **토큰 발급**: `log_in_usecase`에서 로그인 성공 시 JWT 토큰 생성
2. **토큰 검증**: Handler에서 요청 시마다 토큰 유효성 검증
3. **토큰 갱신 알림**: 토큰 만료 1시간 전 헤더에 갱신 필요 알림
4. **토큰 갱신**: `refresh_usecase`를 통한 토큰 갱신 기능 제공
5. **사용자 정보 추출**: 토큰에서 사용자 정보를 추출하여 비즈니스 로직에 활용

## 보안 고려사항

- JWT secret key는 충분히 복잡하고 안전한 값으로 설정해야 합니다.
- 토큰 만료 시간을 적절히 관리하여 보안을 유지해야 합니다.
- 토큰 검증 실패 시 적절한 오류 응답을 반환해야 합니다.
- 토큰 갱신 시 최신 유저 정보를 사용하여 데이터 일관성을 보장해야 합니다.

## PR 제출 시 요구사항

- 다음 항목에 대한 테스트 결과를 첨부해야 합니다:
  - **로그인 성공 후 토큰 발급**: `log_in` 로직을 성공한 후 토큰이 발급된 부분을 캡처해주세요
  - **토큰 만료 시간 테스트**: 임의로 토큰 생성 시 만료시간을 짧게 설정하고, 토큰을 사용하는 라우터를 실행하여 header에 `X-Token-Refresh-Required` 값이 존재하는지 보여주세요
  - **토큰 갱신 성공**: 토큰 갱신 로직이 성공해서 새로운 토큰을 응답받는 구조도 캡처해주세요
  - **토큰을 사용한 API 요청 성공 케이스**: 임의 경로 하나만 실행하셔도 무관합니다.
  - **토큰 없이 API 요청 시 실패 케이스**
  - **토큰 갱신 전용 라우터(`/api/auth/refresh-token`)를 통한 토큰 갱신 기능 동작 확인**
