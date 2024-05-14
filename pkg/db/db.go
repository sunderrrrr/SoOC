package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func connect() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("postgres://postgres:562287@localhost:5432/sooc"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())
}
