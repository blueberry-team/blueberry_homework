package route

import (
	"blueberry_homework/internal/handler/user_handler"
	"blueberry_homework/middleware"

	"github.com/go-chi/chi/v5"
)

func UserRouter(h *user_handler.UserHandler) chi.Router {
	r := chi.NewRouter()

	// 토큰 인증 필요 없는 라우트
	r.Post("/sign-up", h.SignUp)
	r.Post("/log-in", h.Login)

	// 토큰 인증이 필요한 라우트 (protected route)
	r.Group(func(pr chi.Router) {
		pr.Use(middleware.VerifyToken)
		pr.Get("/get-user", h.GetUser)
		pr.Put("/change-user", h.ChangeUser)
	})
	return r
}
