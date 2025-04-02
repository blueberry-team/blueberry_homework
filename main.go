package main

import (
	"blueberry_homework_go_gin/handler"
	"blueberry_homework_go_gin/repository"

	"github.com/gin-gonic/gin" // 패키지? import 하는 것?
)

func main() {
	// Initialize repository and handler
	repo := repository.NewNameRepository() //바로밑에 핸들러의 파라미터로 넣어줌
	nameHandler := handler.NewNameHandler(repo)

	// Create Gin router
	r := gin.Default()

	// Define routes
	r.POST("/names", nameHandler.CreateName)
	r.GET("/names", nameHandler.GetName)//메서드를 함수값으로 넘김? nameHandler.CreateName 이게 마치 CreateName(nameHandler, c) 이렇게 쓰는 것 같음. c는 gin이 전달하는 *gin.Context 타입의 파라미터임

	// Start server
	r.Run(":8080")
}
