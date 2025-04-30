package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/HomeProjectSandbox/basicDockerServices/webserverpostgress/mydb"
	_ "github.com/lib/pq" // register driver to database/sql package

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	connStr := "postgres://myuser:mypw@db:5432/myuser?sslmode=disable"
	//connStr := "postgres://myuser:mypw@localhost:8082/myuser?sslmode=disable"

	m, err := migrate.New(
		"file://database/migrations",
		connStr,
	)

	if err != nil {
		log.Fatalf("Failed to initialize migrations: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully.")

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	ctx := context.Background()
	queries := mydb.New(db)

	//list all
	authors, err := queries.ListAuthors(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println(authors)

	author := mydb.CreateAuthorParams{
		Name: "John",
		Bio: sql.NullString{
			String: "dummy bio",
			Valid:  true,
		},
	}
	a, err := queries.CreateAuthor(ctx, author)
	if err != nil {
		panic(err)
	}
	fmt.Println("author added", a)
	//fmt.Println(authors[0].Bio.String)

}
