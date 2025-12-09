package main

import (
    "log"
    "time"
)

func main() {
    log.Println("Rating service started on :8082")

    for {
        // временный цикл (потом вставим пересчёт)
        time.Sleep(10 * time.Second)
        log.Println("rating tick")
    }
}
