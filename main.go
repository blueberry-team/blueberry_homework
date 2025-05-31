package main

import (
	"blueberry_homework/app"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 애플리케이션 생성 및 초기화
	application, err := app.Init()
	if err != nil {
		log.Fatalf("애플리케이션 초기화 실패: %v", err)
	}
	defer application.Session.Close()

	// 서버 시작
	serverAddr := fmt.Sprintf(":%s", application.Config.ServerPort)
	fmt.Printf("서버가 시작되었습니다: http://localhost%s\n", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, application.Router))
}
