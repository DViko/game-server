package data_base

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type DataBase struct {
}

func (d *DataBase) ConnectDB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:root@localhost:5432/authentication")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected")

	defer conn.Close(context.Background())

	var username string
	err = conn.QueryRow(context.Background(), "SELECT username FROM users WHERE email = 'jane@gmail.com'").Scan(&username)
	if err != nil {
		log.Fatalf("Query failed: %v\n", err)
	}

	fmt.Println("User:", username)
	return conn
}
