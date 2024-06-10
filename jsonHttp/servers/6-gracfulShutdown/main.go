package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	http.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second) // чето делаем типа 3 секунды
		_, _ = fmt.Fprint(w, "slow ok")
	})

	server := &http.Server{Addr: "127.0.0.1connectionSimple:8080"}

	go func() {
		// стартуем сервеср в горутине и блокируем горутину (ListenAndServe ввсегда возвращает ошибку и мы ждем ошибку)
		srvErr := server.ListenAndServe()
		if srvErr != http.ErrServerClosed {
			log.Fatalf("http server listen err: %v", srvErr)
		}
		fmt.Println("Server successfully stopped")
	}()

	termChan := make(chan os.Signal)                         // делаем канал и слушаем терминал
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM) // слушаем системыне вызовы
	termSignal := <-termChan                                 // тут блокировка работы горутины (основного потока)
	// т.е мы ждем в канал определенных сигналов и залипаем пока эти сигналы не будут получены

	log.Println("Graceful shutdown starter with signal: ", termSignal)

	// код ниже - это КОГДА ПОЙМАЛИ ЗАВЕРШЕНИЕ ПРОГРАММЫ
	// Создаем контекст завершения программы с таймаутом
	closeCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // время должно быть больше чем отводится на ответ сервера
	defer cancel()
	// КОГДА Просим сервер завершить работу
	err := server.Shutdown(closeCtx)
	if err != nil {
		log.Println("Graceful shutdown err:", err)
	}
	log.Println("Graceful shutdown err complete")
}
