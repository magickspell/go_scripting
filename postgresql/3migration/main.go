package main

import (
	"database/sql"
	"embed"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
	"log"
)

//go:embed migrations/*.sql
var embedMigrationsFS embed.FS

func main() {
	db, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:8432/go_db")
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	goose.SetBaseFS(embedMigrationsFS)

	const cmd = "up"
	err = goose.Run(cmd, db, "migrations")
	if err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
}
