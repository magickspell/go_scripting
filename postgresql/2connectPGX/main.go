package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
)

func main() {
	fmt.Println()

	// создаем контекс
	ctx := context.Background()
	// и подключаемся к БД
	conn, err := pgx.Connect(ctx, "postgresql://postgres:postgres@localhost:8432/go_db")
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer conn.Close(ctx)

	var pgVersion string
	err = conn.QueryRow(ctx, "select version()").Scan(&pgVersion)
	if err != nil {
		log.Fatalf("failed to query: %v", err)
	}

	fmt.Println("Done QUERY: ", pgVersion)
}
