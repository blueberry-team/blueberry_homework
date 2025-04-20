# REST API 미니 테스트 PART.1

이 리포지토리는 다양한 프레임워크를 사용한 간단한 REST API 구현을 위한 미니 테스
트 프로젝트입니다.

## 프레임워크

각 참여자는 다음 프레임워크를 사용합니다:

- 해린 - Python (Django)
- 정우 - Python (FastAPI)
- 상화 - Go (Chi)
- 승재 - Go (Gin)
- 한규 - .NET
- 재혁 - Rust (Axum)
- 상일 - .NET

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

간단한 테스트 구조를 생성하는 것이 목표이며, 각 프레임워크에서 동일한 기능을 구
현하는 방식을 비교합니다.

## 추가 구현 사항

- 시간이 여유롭다면 유효성 검사(Validation)와 같은 기능들을 추가 구현하셔도 좋습
  니다.
- 프론트엔드는 원하는 웹 애플리케이션을 사용하셔도 좋고, curl을 통한 API 테스트
  도 가능합니다.

---

# REST API 미니 테스트 PART.2

## 데이터 구조 변경

- 이전에는 빈 배열(`[]`)에 데이터를 저장하였으나, PART.2에서는 Map 형태로 저장합
  니다.
- 배열 안에 저장될 내용은 `[{name: "NAME"}, {name: "KING"}]` 형식의 객체 배열로
  변경합니다.
- GET 요청 응답도 `[{name: "KING"}]` 형태로 반환되어야 합니다.

## 응답 형식 변경

- Create 요청과 GET 요청 시 응답에 `message: success` 필드가 포함되어야 합니다.
- 예시 응답 형식:
  ```json
  {
    "message": "success",
    "data": [{ "name": "NAME" }, { "name": "KING" }]
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
- 예: 인덱스 1을 삭제 요청하면 해당 위치의 항목이 제거되고 뒤의 항목들이 앞으로
  이동해야 합니다.

### 사용 가능한 함수명

- `deleteName` - 특정 인덱스의 이름 항목을 삭제하는 기능

## PR 제출 시 요구사항

- PR을 올릴 때 모든 기능이 정상 동작하는지 확인할 수 있도록 Postman이나 Swagger
  에서 모든 API 응답 구조를 캡처하여 첨부해야 합니다.
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
    "data": [
      /* 결과 데이터 */
    ]
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

- PR을 올릴 때 모든 기능이 정상 동작하는지 확인할 수 있도록 Postman이나 Swagger
  에서 모든 API 응답 구조를 캡처하여 첨부해야 합니다.
- 다음 항목에 대한 캡처를 포함해야 합니다:
  - 데이터 생성(Create) 응답
  - 데이터 조회(GET) 응답
  - 인덱스로 데이터 삭제(deleteIndex) 응답
  - 이름으로 데이터 삭제(deleteName) 응답
  - 유효성 검증 실패 시 오류 응답
- 캡처 이미지는 PR 설명에 첨부하여 제출합니다.
