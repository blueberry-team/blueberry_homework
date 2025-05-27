package main

import (
	"blueberry_homework_go_gin/handler"
	"blueberry_homework_go_gin/repository"
	"blueberry_homework_go_gin/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// 계층별 초기화: Repository -> UseCase -> Handler
	repo := repository.NewNameRepository()
	nameUseCase := usecase.NewNameUseCase(repo)
	nameHandler := handler.NewNameHandler(nameUseCase)

	// Gin 라우터 생성
	r := gin.Default()

	// 라우트 정의
	r.POST("/create-name", nameHandler.CreateName)
	r.GET("/get-names", nameHandler.GetNames)
	r.DELETE("/delete-index", nameHandler.DeleteByIndex) // 인덱스로 삭제 (기존 /delete-name 변경)
	r.DELETE("/delete-name", nameHandler.DeleteByName)   // 이름으로 삭제 (신규)

	// 서버 시작
	r.Run(":8080")
}
