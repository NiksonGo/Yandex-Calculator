package main

import (
	"log"
	"net/http"
   
	"github.com/NiksonGo/Yandex-Calculator/internal/handler"
   )
   
   func main() {
	http.HandleFunc("/api/v1/calculate", handler.CalculateHandler)
   
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
	 log.Fatalf("Server failed to start: %v", err)
	}
   }
   