# 프로젝트 아키텍처 플로우차트

```mermaid
flowchart TD
    subgraph API Layer
        A1[FastAPI Router]
        A2[Handler]
    end
    subgraph Application Layer
        B1[UseCase]
    end
    subgraph Data Layer
        C1[Repository]
        C2[ScyllaDB]
    end

    A1 --> A2
    A2 --> B1
    B1 --> C1
    C1 --> C2

    %% 인증 플로우
    A1_Auth[Auth Router] --> A2_Auth[Auth Handler] --> B1_Auth[Auth UseCase] --> C1_Auth[UserRepository] --> C2

    %% 이름 관리 플로우
    A1_Name[Name Router] --> A2_Name[Name Handler] --> B1_Name[Name UseCase] --> C1_Name[NameRepository] --> C2

    %% 회사 관리 플로우
    A1_Company[Company Router] --> A2_Company[Company Handler] --> B1_Company[Company UseCase] --> C1_Company[CompanyRepository] --> C2

    %% DB 내부
    C2 -.->|users, companies 테이블| C2
```

---

- **API Layer**: FastAPI 라우터와 핸들러가 요청을 받아 처리합니다.
- **Application Layer**: UseCase에서 비즈니스 로직을 담당합니다.
- **Data Layer**: Repository가 ScyllaDB와 직접 통신합니다.
- 인증, 이름, 회사 등 모든 주요 기능이 동일한 계층 구조를 따릅니다.
- DB에는 users, companies 등 실제 테이블이 존재합니다. 