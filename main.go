package main

import (
	"context"
	"ginweb/ent"
	"ginweb/infrastructure"
	"log"
	"net/http"
	"os"
)

func main() {
	conn, err := DBconn()
	log.Print(conn)
	if err != nil {
		log.Print(err)
	}
	defer conn.Close()

	ctx := context.Background()
	if err := conn.Schema.WriteTo(ctx, os.Stdout); err != nil {
		log.Printf("failed printing schema changes: %v", err)
	}
	if err := conn.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	mux, err := infrastructure.SetUpServer(conn)
	if err != nil {
		log.Print(err)
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: mux.Router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Print(err)
	}
}

func DBconn() (client *ent.Client, err error) {
	client, err = infrastructure.Open()
	return client, err
}
