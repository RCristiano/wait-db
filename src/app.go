package main

import (
	// "time"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
)
    
func main() {
    pgdb := pg.Connect(&pg.Options{
        Addr:     ":5432",
        User:     "postgres",
        Password: "postgres",
        Database: "postgres",
    })
    _, err := pgdb.Exec("SELECT 1")
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Connection sucessful!")
    os.Exit(0)
} 
