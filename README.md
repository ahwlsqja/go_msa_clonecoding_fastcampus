# Go MSA E-commerce

Go 기반 마이크로서비스 아키텍처(MSA) 이커머스 프로젝트

## 프로젝트 개요

이 프로젝트는 Go 언어를 사용하여 MSA 구조의 이커머스 시스템을 구축하는 학습 프로젝트입니다.
클린 아키텍처와 레이어드 아키텍처 패턴을 적용하여 확장 가능하고 유지보수가 용이한 구조로 설계되었습니다.

## 기술 스택

- **Language**: Go 1.23.3
- **Web Framework**: Gin
- **Database**: Redis
- **Configuration**: Viper
- **Logging**: Logrus

## 프로젝트 구조

```
go-msa/
├── cmd/
│   └── user/
│       ├── resource/     # 외부 리소스 연결 (Redis, DB)
│       └── repository/    # 데이터 접근 계층
├── config/               # 설정 관리
├── handler/              # HTTP 요청 처리 (Controller)
├── infrastructure/       # 인프라 관련 (로깅 등)
│   └── log/
├── routes/              # 라우팅 설정
├── service/             # 비즈니스 로직
└── main.go             # 애플리케이션 진입점
```

## 아키텍처 패턴

### 레이어드 아키텍처

1. **Handler Layer**: HTTP 요청/응답 처리
2. **Service Layer**: 비즈니스 로직 구현
3. **Repository Layer**: 데이터 접근 및 영속성 관리
4. **Resource Layer**: 외부 리소스 연결 (Redis, Database)

## 설치 및 실행

### 사전 요구사항

- Go 1.23.3 이상
- Redis 서버

### 의존성 설치

```bash
go mod download
go mod vendor
```

### 실행

```bash
go run main.go
```

서버는 기본적으로 8080 포트에서 실행됩니다.

## API 엔드포인트

### Health Check

```
GET /ping
```

응답:
```json
{
    "status": "OK"
}
```

## 환경 설정

`config.yaml` 파일을 통해 애플리케이션 설정을 관리합니다:

- **App**: 포트 설정
- **Redis**: Host, Port, Password 설정
- **Database**: 추후 추가 예정

## 개발 노트

이 프로젝트는 MSA 학습을 위한 클론 코딩 프로젝트입니다.
NestJS와 유사한 구조로 설계되어 있으며, 의존성 주입과 싱글톤 패턴을 적용하였습니다.

### 주요 특징

- 싱글톤 패턴을 통한 리소스 관리
- 레이어 간 명확한 책임 분리
- 설정 외부화를 통한 환경별 관리
- 구조화된 로깅 시스템

## 향후 계획

- PostgreSQL 데이터베이스 연동
- 사용자 인증/인가 기능
- 상품 관리 서비스
- 주문 처리 서비스
- gRPC를 통한 서비스 간 통신
- Docker 컨테이너화
- Kubernetes 배포

## 참고사항

이 프로젝트는 학습 목적으로 제작되었으며, 프로덕션 환경에서 사용하기 위해서는 추가적인 보안 및 최적화 작업이 필요합니다.