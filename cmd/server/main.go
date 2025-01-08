// cmd/server/main.go
package main

import (
    "log"
    "net/http"
    "database/sql"
    
    "invoice-generator/internal/auth"
    "invoice-generator/internal/handlers"
    "invoice-generator/internal/database"
    
    "github.com/gorilla/mux"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Initialize database
    db, err := sql.Open("sqlite3", "./invoice.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Initialize database tables
    if err := database.InitDB(db); err != nil {
        log.Fatal("Failed to initialize database:", err)
    }

    // Initialize services and handlers
    authService := auth.NewAuthService(db)
    authHandler := handlers.NewAuthHandler(authService)
    invoiceHandler := handlers.NewInvoiceHandler()

    // Initialize router
    r := mux.NewRouter()

    // Auth routes
    r.HandleFunc("/auth/login", authHandler.ShowLogin).Methods("GET")
    r.HandleFunc("/auth/login", authHandler.Login).Methods("POST")
    r.HandleFunc("/auth/register", authHandler.ShowRegister).Methods("GET")
    r.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
    r.HandleFunc("/auth/logout", authHandler.Logout).Methods("POST")
    r.HandleFunc("/invoice", invoiceHandler.ShowInvoice).Methods("GET")
    r.HandleFunc("/invoice/generate", invoiceHandler.GenerateInvoice).Methods("POST")

    // Serve static files
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Start server
    log.Println("Server starting on port 8080...")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}