package main

import (
	"blueberry_homework/app"
	"blueberry_homework/internal/data/repository"
	"blueberry_homework/internal/db"
	"blueberry_homework/internal/domain/usecase"
	"blueberry_homework/internal/handler"
	"blueberry_homework/route"
	"fmt"
	"log"
	"net/http"
)

func main() {
    session, err := db.InitScylla()
	if err != nil {
		panic(err)
	}

    // name
    nameRepo := repository.NewNameRepository(session)
    nameUsecase := usecase.NewNameUsecase(nameRepo)
    nameHandler := handler.NewNameHandler(nameUsecase)

    // company
    companyRepo := repository.NewCompanyRepository(session)
    createCompanyUsecase := usecase.NewCreateCompanyUsecase(nameRepo, companyRepo)
    companyUsecase := usecase.NewCompanyUsecase(companyRepo)
    companyHandler := handler.NewCompanyHandler(createCompanyUsecase, companyUsecase)

    application := app.NewApp()
    fmt.Println("app start!", application)

    application.Router.Mount("/names", route.NameRouter(nameHandler))
    application.Router.Mount("/companies", route.CompanyRouter(companyHandler))

    fmt.Println("route set up done!!")
    log.Fatal(http.ListenAndServe(":3000", application.Router))
}
