package route

import (
	"blueberry_homework/internal/handler/company_handler"
	"blueberry_homework/middleware"

	"github.com/go-chi/chi/v5"
)

func CompanyRouter(h *company_handler.CompanyHandler) chi.Router {
	r := chi.NewRouter()

	// protected route
	// group 생략해도 되지만 user 와의 통일성을 위해 한번 더 묶어줌
	r.Group(func(pr chi.Router) {
		pr.Use(middleware.VerifyToken)
		pr.Post("/create-company", h.CreateCompany)
		pr.Get("/get-user-company", h.GetUserCompany)
		pr.Put("/change-company", h.ChangeCompany)
		pr.Delete("/delete-company", h.DeleteCompany)
	})

	return r
}
