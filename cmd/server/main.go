package main

import (
	"fmt"
	"log"
	"net/http"

	pg "vamaracinggp/internal/db/postgres"
	h "vamaracinggp/internal/server/http"
)

func main() {
	db, err := pg.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := &h.Handler{
		DB: db,
	}

	fmt.Println("Server started and listening on localhost:8080")
	if err := http.ListenAndServe("localhost:8080", router.NewRouter()); err != nil {
		log.Fatal("Server failed:", err)
	}
}
