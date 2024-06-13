package main

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v4/stdlib" // регаем драйвер подключения к БД
	"github.com/jmoiron/sqlx"
	"log"
)

func main() {
	fmt.Println()

	db, err := sqlx.Open("pgx", "postgres://postgres:postgres@localhost:8432/go_db")
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	sb := psql.Select("id", "name").From("products")

	productIDs := []int64{2, 3, 1, 5}
	if len(productIDs) != 0 {
		sb = sb.Where(sq.Eq{"id": productIDs})
	}

	// это просто пример как можно
	_ = sq.Or{ // OR - AND
		sq.Eq{"name": "Audi"},       // name = 'Audi'
		sq.Like{"name": "%cap"},     // name like '%cap'
		sq.NotEq{"created_at": nil}, // crated_at IS NOT NULL
	}

	query, args, err := sb.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	ctx := context.Background()

	type product struct {
		ID   int64  `db:"id"`
		Name string `db:"name"`
	}

	var products []product
	rows, err := db.QueryxContext(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to execute query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p product
		err = rows.StructScan(&p)
		if err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		products = append(products, p)
	}

	for _, p := range products {
		fmt.Printf("%d - %s \n", p.ID, p.Name)
	}
}
