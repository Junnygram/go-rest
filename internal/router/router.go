package router

import (
	"net/http"

	"github.com/junnygram/go-rest/internal/handler"
)

func New() *http.ServeMux {
	r := http.NewServeMux()

	// Create user route.
	r.HandleFunc("POST /sign_in", handler.CreateUser())

	// Login user route.
	r.HandleFunc("POST /log_in", handler.LoginUser())

	// Create news route.
	r.HandleFunc("POST /news", handler.PostNews())

	// Get all news.
	r.HandleFunc("GET /news", handler.GetAllNews())

	// Get news by ID.
	r.HandleFunc("GET /news/{news_id}", handler.GetAllNews())

	// Update news by ID.
	r.HandleFunc("PUT /news/{news_id}", handler.UpdateNewsById())

	// Delete news by ID.
	r.HandleFunc("DELETE /news/{news_id}", handler.DeleteNewsById().ServeHTTP)

	return r
}
