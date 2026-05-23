package routes

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"main/db"
	"main/services"
)
var client *db.PrismaClient

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from root, path=%q", r.URL.Path)
}


func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "users list")
}

func Servers(c *db.PrismaClient) {
	client = c

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/users", handleUsers)
	mux.HandleFunc("/post", services.HandlePostList)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}