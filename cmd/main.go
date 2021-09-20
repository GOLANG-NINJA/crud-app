package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/GOLANG-NINJA/crud-app/internal/repository/psql"
	"github.com/GOLANG-NINJA/crud-app/internal/service"
	"github.com/GOLANG-NINJA/crud-app/internal/transport/rest"

	_ "github.com/lib/pq"
)

func main() {
	// init db
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres sslmode=disable password=qwerty123")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// init deps
	booksRepo := psql.NewBooks(db)
	booksService := service.NewBooks(booksRepo)
	handler := rest.NewHandler(booksService)

	// init & run server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler.InitRouter(),
	}

	log.Println("SERVER STARTED AT", time.Now().Format(time.RFC3339))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
