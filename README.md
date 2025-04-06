# REST API 미니 테스트

이 리포지토리는 다양한 프레임워크를 사용한 간단한 REST API 구현을 위한 미니 테스트 프로젝트입니다.

## 프레임워크

각 참여자는 다음 프레임워크를 사용합니다:

- 해린 - Python (Django)
- 정우 - Python (FastAPI)
- 상화 - Go (Chi)
- 승재 - Go (Gin)
- 한규 - .NET
- 재혁 - Go (Fiber)

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
