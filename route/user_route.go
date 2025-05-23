package route

import (
	"blueberry_homework/internal/handler"

	"github.com/go-chi/chi/v5"
)

func UserRouter(h *handler.UserHandler) chi.Router {
	r := chi.NewRouter()

	// user
	r.Post("/sign-up", h.SignUp)
	r.Post("/log-in", h.Login)
	r.Get("/get-user", h.GetUser)
	r.Post("/change-user", h.ChangeUser)

	return r
}
