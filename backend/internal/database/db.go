package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB(user, password, dbName string) {
	connStr := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", user, password, dbName)

	var err error
	DB, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Не удалось создать пул подключений: %v\n", err)
	}

	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("Ошибка подключения к PostgreSQL: %v\n", err)
	}

	fmt.Println("Успешное подключение к PostgreSQL!")
}