package main

import (
	"blueberry_homework/app"
	"blueberry_homework/internal/data/repository"
	"blueberry_homework/internal/domain/usecase"
	"blueberry_homework/internal/handler"
	"blueberry_homework/route"
	"fmt"
	"net/http"

	"github.com/gocql/gocql"
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

    // 클러스터 설정 (로컬 Scylla에 연결)
    cluster := gocql.NewCluster("localhost")

    // Keyspace 없이 연결만 테스트할 것이므로 생략
    // cluster.Keyspace = "..."

    // 일관성 설정 (안정성 있는 기본값)
    cluster.Consistency = gocql.Quorum

    // 세션 생성 시도
    session, err := cluster.CreateSession()
    if err != nil {
        panic(fmt.Sprintf("❌ Scylla 연결 실패: %v", err))
    }
    defer session.Close()
    fmt.Println("✅ Scylla 연결 성공!")


    application := app.NewApp()
    fmt.Println("app start!", application)

    application.Router.Mount("/names", route.NameRouter(nameHandler))
    application.Router.Mount("/companies", route.CompanyRouter(companyHandler))
    fmt.Println("route set up done!!")

    http.ListenAndServe(":3000", application.Router)
}
