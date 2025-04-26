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

