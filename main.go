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
    nameRepo := repository.NewNameRepository()
    nameUsecase := usecase.NewNameUsecase(nameRepo)
    nameHandler := handler.NewNameHandler(nameUsecase)

    application := app.NewApp()
    fmt.Println("app start!", application)

    application.Router.Mount("/names", route.NameRouter(nameHandler))
    fmt.Println("route set up done!!")

    http.ListenAndServe(":3000", application.Router)
}
