package main

import (
	"net/http"
	"fmt"

	"blueberry_homework/internal/app"
	"blueberry_homework/internal/handler"
	"blueberry_homework/internal/repository"
	"blueberry_homework/internal/route"
)

func main() {
    nameRepo := repository.NewNameRepository()
    nameHandler := handler.NewNameHandler(nameRepo)

    application := app.NewApp()
    fmt.Println("app start!", application)

    application.Router.Mount("/names", route.NameRouter(nameHandler))
    fmt.Println("route set up done!!")

    http.ListenAndServe(":3000", application.Router)
}
