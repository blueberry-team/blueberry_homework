package app

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type App struct {
	Router *chi.Mux
}

func NewApp() *App {
	router := setupRouter()

	return &App{
		Router: router,
	}
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
