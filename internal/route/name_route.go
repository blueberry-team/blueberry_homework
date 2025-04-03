package route

import (
	"github.com/go-chi/chi/v5"

    "blueberry_homework/internal/handler"
)

func NameRouter(h *handler.NameHandler) chi.Router {
    r := chi.NewRouter()

    r.Post("/create-name", h.CreateName)
    r.Get("/get-names", h.GetNames)

    return r
}
