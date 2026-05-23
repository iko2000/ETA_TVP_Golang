package services

import (
	"encoding/json"
	"net/http"
	"main/db"
)
var client *db.PrismaClient

func SetClient(c *db.PrismaClient) {
	client = c
}

func HandlePostList(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    posts, err := client.Post.FindMany().Exec(ctx)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(posts); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}