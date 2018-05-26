package main

import (
	"database"
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	defer database.DB.Close()

	// add database
	_, err := database.Init()
	if err != nil {
		log.Println("connection to DB failed, aborting...")
		log.Fatal(err)
	}

	log.Println("connected to DB")

	// print env
	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev mode")
	}
	var addr = flag.String("localhost", ":8080", "/")
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
