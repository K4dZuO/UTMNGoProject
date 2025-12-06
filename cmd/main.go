package main

import (
    "fmt"
    "log"
    "go_back/internal/database"
)

func main() {
    pool, err := database.NewPostgresPool()
    if err != nil {
        log.Fatalf("DB error: %v", err)
    }
    defer pool.Close()

    fmt.Println("PostgreSQL connected!")

    // Далее будет запуск сервера и обработчиков
}
