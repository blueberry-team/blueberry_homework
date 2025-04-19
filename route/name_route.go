package route

import (
	"github.com/go-chi/chi/v5"

    "blueberry_homework/internal/handler"
)

func NameRouter(h *handler.NameHandler) chi.Router {
    r := chi.NewRouter()

    r.Post("/create-name", h.CreateName)
    r.Get("/get-names", h.GetNames)
    r.Delete("/delete-index", h.DeleteByIndex)
    r.Delete("/delete-name", h.DeleteByName)


    return r
}
