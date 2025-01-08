// cmd/server/main.go
package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    // Initialize router
    r := mux.NewRouter()
    
    // Serve static files
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    
    // Basic health check route
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Server is running!"))
    }).Methods("GET")
    
    // Start server
    log.Println("Server starting on port 8080...")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}