package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"blueberry_homework_go_gin/app"
)

func main() {
	// 애플리케이션 초기화 (한 줄로 모든 의존성 초기화)
	application, err := app.Init()
	if err != nil {
		log.Fatalf("❌ 애플리케이션 초기화 실패: %v", err)
	}

	// Graceful shutdown을 위한 시그널 핸들링
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		log.Println("🛑 애플리케이션 종료 신호 받음")
		if err := application.Shutdown(); err != nil {
			log.Printf("❌ 애플리케이션 종료 중 오류: %v", err)
		}
		os.Exit(0)
	}()

	// 서버 시작
	if err := application.Run(); err != nil {
		log.Fatalf("❌ 서버 시작 실패: %v", err)
	}
}
