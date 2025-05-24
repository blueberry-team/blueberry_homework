package app

import (
	"fmt"
	"log"

	"blueberry_homework_go_gin/config"
	"blueberry_homework_go_gin/db"
	"blueberry_homework_go_gin/handler"
	"blueberry_homework_go_gin/repository"
	"blueberry_homework_go_gin/usecase"
	"github.com/gin-gonic/gin"
)

// App 애플리케이션 구조체
type App struct {
	Router *gin.Engine
	Config *config.Config
}

// Init 애플리케이션을 초기화하고 반환 (한 줄로 호출 가능)
func Init() (*App, error) {
	log.Println("🚀 애플리케이션 초기화 시작...")

	// 1. 설정 로드
	cfg := config.LoadConfig()
	log.Printf("✅ 설정 로드 완료: %s 환경", cfg.AppEnv)

	// 2. 데이터베이스 초기화
	if err := db.InitMongoDB(cfg); err != nil {
		return nil, fmt.Errorf("데이터베이스 초기화 실패: %v", err)
	}

	// 3. Gin 라우터 생성
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	// 4. 의존성 주입 및 라우터 설정
	setupRoutes(router)

	app := &App{
		Router: router,
		Config: cfg,
	}

	log.Println("✅ 애플리케이션 초기화 완료")
	return app, nil
}

// setupRoutes 라우터와 핸들러를 설정 (의존성 주입)
func setupRoutes(router *gin.Engine) {
	// Repository 계층 초기화
	nameRepo := repository.NewNameRepository()
	companyRepo := repository.NewCompanyRepository()

	// UseCase 계층 초기화 (Repository 의존성 주입)
	nameUseCase := usecase.NewNameUseCase(nameRepo)
	companyUseCase := usecase.NewCompanyUseCase(companyRepo, nameRepo)

	// Handler 계층 초기화 (UseCase 의존성 주입)
	nameHandler := handler.NewNameHandler(nameUseCase)
	companyHandler := handler.NewCompanyHandler(companyUseCase)

	// === 헬스체크 라우트 ===
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "서버가 정상 동작중입니다",
			"database": "MongoDB 연결됨",
		})
	})

	// === 사용자 관련 라우트 ===
	router.POST("/create-name", nameHandler.CreateName)           // 사용자 생성
	router.GET("/get-names", nameHandler.GetNames)               // 모든 사용자 조회
	router.PUT("/change-name", nameHandler.ChangeName)           // 사용자 이름 변경
	router.DELETE("/delete-index", nameHandler.DeleteByIndex)    // 인덱스로 사용자 삭제
	router.DELETE("/delete-name", nameHandler.DeleteByName)      // 이름으로 사용자들 삭제

	// === 회사 관련 라우트 ===
	router.POST("/create-company", companyHandler.CreateCompany)    // 회사 생성
	router.GET("/get-companies", companyHandler.GetAllCompanies)    // 모든 회사 조회

	log.Println("✅ 라우터 설정 완료")
}

// Run 애플리케이션을 실행
func (a *App) Run() error {
	addr := ":" + a.Config.ServerPort
	log.Printf("🌐 서버 시작: http://localhost%s", addr)
	log.Println("📡 API 엔드포인트:")
	log.Println("   GET  /health           - 헬스체크")
	log.Println("   POST /create-name      - 사용자 생성")
	log.Println("   GET  /get-names        - 사용자 목록")
	log.Println("   PUT  /change-name      - 사용자 이름 변경")
	log.Println("   DELETE /delete-index   - 인덱스로 사용자 삭제")
	log.Println("   DELETE /delete-name    - 이름으로 사용자 삭제")
	log.Println("   POST /create-company   - 회사 생성")
	log.Println("   GET  /get-companies    - 회사 목록")
	log.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	return a.Router.Run(addr)
}

// Shutdown 애플리케이션을 종료
func (a *App) Shutdown() error {
	log.Println("🛑 애플리케이션 종료 중...")
	return db.CloseMongoDB()
}
