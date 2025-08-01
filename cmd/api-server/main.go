package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/junnygram/go-rest/internal/router"
)

func main() {
	r := router.New()
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("failed to start server")
	}
	hello := "This is olalaye"
	fmt.Println(hello)
}
