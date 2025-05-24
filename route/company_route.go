package route

import (
	"blueberry_homework/internal/handler"

	"github.com/go-chi/chi/v5"
)

func CompanyRouter(h *handler.CompanyHandler) chi.Router {
	r := chi.NewRouter()

	r.Post("/create-company", h.CreateCompany)
	r.Get("/get-user-company", h.GetUserCompany)
	r.Put("/change-company", h.ChangeCompany)
	r.Delete("/delete-company", h.DeleteCompany)

	return r
}
