package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/HomeProjectSandbox/basicDockerServices/webserverpostgress/config"
	"github.com/HomeProjectSandbox/basicDockerServices/webserverpostgress/mydb"
	_ "github.com/lib/pq" // register driver to database/sql package

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Handler struct {
	queries *mydb.Queries
}

type AddAuthorRequest struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

func main() {
	c := config.GetConfig("/config", "config") //remove this if local development... (not in container)

	fmt.Println(c)
	fmt.Println(c.App.DbName)
	fmt.Println(c.App.DbUser)

	//TODO: need to use the config
	//connStr := "postgres://myuser:mypw@db:5432/myuser?sslmode=disable"
	// Construct the connection string using the loaded configuration
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.App.DbUser, // Username from config
		c.App.DbPw,   // Password from config
		c.App.DbName, // Host from config
		c.App.DbPort, // Port from config
		c.App.DbUser,
	)

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

	h := Handler{
		queries: queries,
	}

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

	http.HandleFunc("/", h.handler)

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		// Optionally, you can serve a favicon here or leave it empty
	})

	http.HandleFunc("/get", h.handlerGet)
	http.HandleFunc("/add", h.handlerAdd)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (h *Handler) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello from handler")
	http.ServeFile(w, r, "./static/index.html")

}

func (h *Handler) handlerGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello from Get")
	// Set the response header to indicate JSON content
	w.Header().Set("Content-Type", "application/json")

	// Create a context for the database query
	ctx := context.Background()

	authors, err := h.queries.ListAuthors(ctx)
	if err != nil {
		// If there's an error, respond with a 500 Internal Server Error
		http.Error(w, "Failed to fetch authors", http.StatusInternalServerError)
		return
	}

	// Encode the authors slice to JSON and write it to the response
	if err := json.NewEncoder(w).Encode(authors); err != nil {
		// If there's an error encoding the JSON, respond with a 500 Internal Server Error
		http.Error(w, "Failed to encode authors to JSON", http.StatusInternalServerError)
		return
	}

}

/*
curl -X POST http://localhost:8081/add \
-H "Content-Type: application/json" \
-d '{"name": "Jane Doe", "bio": "An accomplished writer."}'
*/
func (h *Handler) handlerAdd(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello from add")

	var req AddAuthorRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	arg := mydb.CreateAuthorParams{
		Name: req.Name,
		Bio:  sql.NullString{String: req.Bio, Valid: req.Bio != ""},
	}

	_, err := h.queries.CreateAuthor(context.Background(), arg)
	if err != nil {
		http.Error(w, "Failed to create author", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(req)
}
