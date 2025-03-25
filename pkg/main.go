package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	pgURL := os.Getenv("POSTGRES_URL")
	serverPort := os.Getenv("SERVER_PORT")

	ctx := context.Background()
	db := NewDatabase(ctx, pgURL)
	defer db.Close()

	e := NewServer(db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", serverPort)))
}
