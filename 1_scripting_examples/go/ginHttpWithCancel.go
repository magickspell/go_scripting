package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Middleware с таймаутом
	r.Use(func(c *gin.Context) {
		// Создаем контекст с таймаутом
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()

		// Заменяем контекст запроса на новый с таймаутом
		c.Request = c.Request.WithContext(ctx)

		// Передаем управление следующему middleware или обработчику
		c.Next()
	})

	// Обработчик, который может выполняться долго
	r.GET("/long-operation", func(c *gin.Context) {
		// Симуляция долгой операции
		select {
		case <-time.After(5 * time.Second):
			c.JSON(http.StatusOK, gin.H{"message": "Операция завершена"})
		case <-c.Request.Context().Done():
			// Если контекст отменен (например, по таймауту)
			c.JSON(http.StatusRequestTimeout, gin.H{"error": "Операция прервана по таймауту"})
		}
	})

	// Запуск сервера
	r.Run(":8080")
}
