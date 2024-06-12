package main

import (
	"database/sql"
	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

func main() {
	db, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:8432/go_db")
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer db.Close()

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory("./fixtures"),
		testfixtures.DangerousSkipTestDatabaseCheck(), // allow non-test DB name
	)
	if err != nil {
		log.Fatalf("test fixtures: %v", err)
	}

	err = fixtures.Load()
	if err != nil {
		log.Fatalf("fixtures load: %v", err)
	}
}
