package main

import (
	"database"
	"flag"
	"log"
	"net/http"
	"os"
)

func withAPIKey(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isValidAPIKey(r.URL.Query().Get("key")) {
			respondErr(w, r, http.StatusUnauthorized, "不正なAPIキーです")
		}
		fn(w, r)
	}
}
func isValidAPIKey(key string) bool {
	return key == "abc123"
}

func withData(d *pg.Session, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer database.DB.Close()

		// add database
		_, err := database.Init()
		if err != nil {
			log.Println("connection to DB failed, aborting...")
			log.Fatal(err)
		}
		setVar(r, "db", database.DB("ballots"))
		f(w, r)
	}
}

func main() {
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
