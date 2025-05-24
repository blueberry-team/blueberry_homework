package app

import (
	"blueberry_homework/config"
	"blueberry_homework/internal/data/repository"
	"blueberry_homework/internal/db"
	"blueberry_homework/internal/domain/usecase/company_usecase"
	"blueberry_homework/internal/domain/usecase/user_usecase"
	"blueberry_homework/internal/handler"
	"blueberry_homework/route"
	"fmt"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/gocql/gocql"
)

type App struct {
	Router  *chi.Mux
	Config  *config.Config
	Session *gocql.Session
}

// NewApp은 애플리케이션을 생성하고 모든 의존성을 초기화합니다
func Init() (*App, error) {
	// 환경 설정 로드
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, fmt.Errorf("설정 로드 실패: %v", err)
	}

	// 라우터 설정
	router := setupRouter()

	// 앱 인스턴스 생성
	app := &App{
		Router: router,
		Config: cfg,
	}

	// Scylla DB 초기화
	session, err := db.InitScylla(cfg)
	if err != nil {
		return nil, fmt.Errorf("DB 초기화 실패: %v", err)
	}
	app.Session = session

	// 레포지토리 초기화
	userRepo := repository.NewUserRepository(session)
	companyRepo := repository.NewCompanyRepository(session)

	// 유스케이스 초기화
	userUsecase := user_usecase.NewUserUsecase(userRepo)
	companyUsecase := company_usecase.NewCompanyUsecase(companyRepo, userRepo)

	// 핸들러 초기화
	userHandler := handler.NewUserHandler(userUsecase)
	companyHandler := handler.NewCompanyHandler(companyUsecase)

	// 라우트 설정
	app.Router.Mount("/users", route.UserRouter(userHandler))
	app.Router.Mount("/companies", route.CompanyRouter(companyHandler))

	fmt.Println("✅ 애플리케이션 초기화 완료!")
	return app, nil
}

// cmd internal config
// cmd - 외부 key 불러와야 할 떄
// config - 환경변수 세팅 config.go
// internal - data / domain/ service / handler
func setupRouter() *chi.Mux {
	r := chi.NewRouter()

	// 미들웨어 설정
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	return r
}
