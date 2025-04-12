package main

import (
	"blueberry_homework_go_gin/handler"
	"blueberry_homework_go_gin/repository"

	"github.com/gin-gonic/gin" // 패키지? import 하는 것?
)

func main() {
	// 레포,핸들러 초기화
	repo := repository.NewNameRepository() //바로밑에 핸들러의 파라미터로 넣어줌
	nameHandler := handler.NewNameHandler(repo)

	//진라우터생성
	r := gin.Default()

	// 루트정의
	r.POST("/create-name", nameHandler.CreateName)
	r.GET("/get-names", nameHandler.GetName)//메서드를 함수값으로 넘김? nameHandler.CreateName 이게 마치 CreateName(nameHandler, c) 이렇게 쓰는 것 같음. c는 gin이 전달하는 *gin.Context 타입의 파라미터임
	r.DELETE("/delete-name", nameHandler.DeleteName)

	// 서버시작
	r.Run(":8080")
}
