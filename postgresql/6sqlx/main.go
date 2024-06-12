package main

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib" // регаем драйвер подключения к БД
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type product struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

func main() {
	fmt.Println()

	db, err := sqlx.Open("pgx", "postgres://postgres:postgres@localhost:8432/go_db")
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	ctx := context.Background()

	fmt.Println("GetContext")

	{
		query := `
			SELECT p.id, p.name, p.created_at
			FROM products p
			WHERE p.id = :product_id
		`

		query, args, err := sqlx.Named(query, map[string]interface{}{
			"product_id": 4,
		})
		if err != nil {
			log.Fatalf("query err: %v", err)
		}

		query = sqlx.Rebind(sqlx.DOLLAR, query)

		p := product{}
		err = db.GetContext(ctx, &p, query, args...)
		if err != nil {
			log.Fatalf("GetContext err: %v", err)
		}
		fmt.Printf("%d - %s (%v)\n", p.ID, p.Name, p.CreatedAt.Year())
	}
}
