package route

import (
	"blueberry_homework/internal/handler/auth_handler"
	"blueberry_homework/middleware"

	"github.com/go-chi/chi/v5"
)

func AuthRouter(h *auth_handler.AuthHandler) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.VerifyToken)
	r.Post("/refresh-token", h.RefreshToken)

	return r
}
