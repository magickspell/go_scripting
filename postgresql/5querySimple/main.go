package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

func main() {
	fmt.Println()

	db, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:8432/go_db")
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	ctx := context.Background()

	// выполнение запросов INSERT
	{
		const query = `
      INSERT INTO products(name, category_id, created_at)
        VALUES ($1, $2, now())
    `

		result, err := db.ExecContext(ctx, query, "Bert", 3)
		if err != nil {
			log.Fatalf("qury error: %v", err)
		}
		//result.LastInsertId() // не поддерживается библой

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatalf("rowsAffected error: %v", err)
		}
		fmt.Println("rowsAffected", rowsAffected)
	}
	fmt.Println()
	// выполнение запросов SELECT 1 строки
	{
		const query = `
      SELECT p.id, p.name FROM products p WHERE p.id = $1; 
    `

		row := db.QueryRowContext(ctx, query, 2)
		var (
			id   int64
			name string
		)
		err := row.Scan(&id, &name)
		if err == sql.ErrNoRows {
			fmt.Println("row not found")
		}
		if err != nil {
			log.Fatalf("rows err: %v", err)
		}

		fmt.Println("row:", " id - ", id, " name - ", name)
	}
	fmt.Println()
	// выполнение запросов INSERT RETURNING 1 строки
	{
		const query = `
      INSERT INTO products (name, category_id, created_at)
      VALUES($1, $2, now())
      RETURNING id
      ; 
    `

		row := db.QueryRowContext(ctx, query, "Ascot cap", 3)

		var id int64
		err := row.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("new record id: ", id)
	}
	fmt.Println()
	// выполнение запросов SELECT n строк
	{
		const query = `
      SELECT p.id, p.name FROM products p WHERE p.category_id = $1;
      ; 
    `

		rows, err := db.QueryContext(ctx, query, 3)
		if err != nil {
			log.Fatalf("select multi ERR: %v", err)
		}
		defer rows.Close()

		type product struct {
			id   int64
			name string
		}
		var products []product

		for rows.Next() {
			var p product
			err = rows.Scan(&p.id, &p.name)
			if err != nil {
				log.Fatalf("parse err: %v", err)
			}
			products = append(products, p)
		}

		for _, p := range products {
			fmt.Printf("%d - %s\n", p.id, p.name)
		}
	}
	// the end
}
