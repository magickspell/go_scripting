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

	// SELECT одного элемента
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
	fmt.Println()
	// select IN
	{
		productIDs := []int64{1, 3, 2}

		query := `SELECT p.id, p.name, p.created_at FROM products p WHERE p.id in (:product_ids)`

		query, args, err := sqlx.Named(query, map[string]interface{}{
			"product_ids": productIDs,
		})
		if err != nil {
			log.Fatalf("query err: %v", err)
		}

		// sqlx.In нужен только когда есть IN
		query, args, err = sqlx.In(query, args...)
		if err != nil {
			log.Fatalf("query IN err: %v", err)
		}

		query = db.Rebind(query)

		p := product{}
		err = db.GetContext(ctx, &p, query, args...)
		if err != nil {
			log.Fatalf("GetContext err: %v", err)
		}
		fmt.Printf("%d - %s (%v)\n", p.ID, p.Name, p.CreatedAt.Year())
	}
	fmt.Println()
	// SELECT несколько элементов
	{
		const query = `SELECT p.id, p.name, p.created_at FROM products p WHERE p.category_id = (:category_id)`

		rows, err := db.NamedQueryContext(ctx, query, map[string]interface{}{
			"category_id": 3,
		})
		if err != nil {
			log.Fatalf("NamedQueryContext err: %v", err)
		}
		defer rows.Close()

		var products []product
		for rows.Next() {
			var p product
			err = rows.StructScan(&p) // парсим строку и проверяем что она product
			if err != nil {
				log.Fatalf("StructScan err: %v", err)
			}
			products = append(products, p)
		}
		for _, p := range products {
			fmt.Printf("%d - %s (%v)\n", p.ID, p.Name, p.CreatedAt.Year())
		}
	}
	fmt.Println()
	// SELECT CONTEXT (самый удобный?)
	{
		ctxNew := context.Background()
		const query = `SELECT p.id, p.name, p.created_at FROM products p ORDER BY created_at DESC`

		var products []product
		err := db.SelectContext(ctxNew, &products, query)
		if err != nil {
			log.Fatalf("SelectContext err: %v", err)
		}
		for _, p := range products {
			fmt.Printf("%d - %s (%v)\n", p.ID, p.Name, p.CreatedAt.Year())
		}
	}

}
