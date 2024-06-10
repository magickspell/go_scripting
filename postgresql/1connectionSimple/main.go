package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	fmt.Println()

	// есть два варианта подключения
	// "user=postgres password=postgres host=localhost port=55432 database=db sslmode=disable"
	db, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:8432/go_db")
	if err != nil {
		log.Fatalf("sql.Open(): %v", err)
	}

	// приостановка работы программы пока не нажмешь ENTER
	_, _ = fmt.Scanln()
	// НАЖМИ ENTER

	// создаем контекс
	ctx := context.Background()
	// и подключаемся к БД
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf("db.PingContext(): %v", err)
	}

	fmt.Println("Done")
}
