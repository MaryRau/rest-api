package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() {
	connStr := "postgres://postgres:322614@localhost:5432/todolistdb"
	db, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(context.Background()); err != nil {
		panic(err)
	}

	fmt.Println("Подключение к базе данных установлено")
	DB = db
}
