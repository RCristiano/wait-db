package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/go-pg/pg/v10"
)

var (
	address  string
	port     string
	user     string
	password string
	database string
	watch    bool
	timeout  int

	exitCode int
)

func init() {
	flag.StringVar(&address, "address", "localhost", "target postgres instance")
	flag.StringVar(&port, "port", "5432", "target postgres port")
	flag.StringVar(&user, "user", "postgres", "target postgres user")
	flag.StringVar(&password, "password", "postgres", "target users password")
	flag.StringVar(&database, "database", "postgres", "target database")
	flag.BoolVar(&watch, "watch", false, "keep watching command, retries connection each 2s.")
	flag.IntVar(&timeout, "timeout", 3600, "how many seconds should a watch run")

	flag.Parse()
}

func main() {

	c1 := make(chan int, 1)
	go func() {

		for {
			pgdb := pg.Connect(&pg.Options{
				Addr:     address + ":" + port,
				User:     user,
				Password: password,
				Database: database,
			})
			_, err := pgdb.Exec("SELECT 1")

			if err != nil {
				log.Println(err)
				exitCode = 1
			} else {
				log.Println("Connection sucessful!")
				exitCode = 0
				watch = false
			}

			if watch == false {
				break
			}
			time.Sleep(2 * time.Second)
		}
		c1 <- exitCode
	}()

	select {
	case res := <-c1:
		os.Exit(res)
	case <-time.After(time.Duration(timeout) * time.Second):
		log.Println("Timed out")
		os.Exit(127)
	}
}
