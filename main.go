package main

import (
	"blueberry_homework_go_gin/handler"
	"blueberry_homework_go_gin/repository"
	"blueberry_homework_go_gin/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// 레포지토리 계층 초기화
	nameRepo := repository.NewNameRepository()
	companyRepo := repository.NewCompanyRepository()

	// 유스케이스 계층 초기화
	nameUseCase := usecase.NewNameUseCase(nameRepo)
	companyUseCase := usecase.NewCompanyUseCase(companyRepo, nameRepo)

	// 핸들러 계층 초기화
	nameHandler := handler.NewNameHandler(nameUseCase)
	companyHandler := handler.NewCompanyHandler(companyUseCase)

	// Gin 라우터 생성
	r := gin.Default()

	// User 관련 라우트 정의
	r.POST("/create-name", nameHandler.CreateName)
	r.GET("/get-names", nameHandler.GetNames)
	r.PUT("/change-name", nameHandler.ChangeName)  // 이름 변경 (신규)
	r.DELETE("/delete-index", nameHandler.DeleteByIndex)  // 인덱스로 삭제
	r.DELETE("/delete-name", nameHandler.DeleteByName)    // 이름으로 삭제

	// Company 관련 라우트 정의
	r.POST("/create-company", companyHandler.CreateCompany)
	r.GET("/get-companies", companyHandler.GetCompanies)

	// 서버 시작
	r.Run(":8080")
}
