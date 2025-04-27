package main

import (
	"blueberry_homework/app"
	"blueberry_homework/internal/data/repository"
	"blueberry_homework/internal/domain/usecase"
	"blueberry_homework/internal/handler"
	"blueberry_homework/route"
	"fmt"
	"net/http"
)

func main() {
    // name
    nameRepo := repository.NewNameRepository()
    nameUsecase := usecase.NewNameUsecase(nameRepo)
    nameHandler := handler.NewNameHandler(nameUsecase)

    // company
    companyRepo := repository.NewCompanyRepository()
    createCompanyUsecase := usecase.NewCreateCompanyUsecase(nameRepo, companyRepo)
    companyUsecase := usecase.NewCompanyUsecase(companyRepo)
    companyHandler := handler.NewCompanyHandler(createCompanyUsecase, companyUsecase)

    application := app.NewApp()
    fmt.Println("app start!", application)

    application.Router.Mount("/names", route.NameRouter(nameHandler))
    application.Router.Mount("/companies", route.CompanyRouter(companyHandler))
    fmt.Println("route set up done!!")

    http.ListenAndServe(":3000", application.Router)
}
