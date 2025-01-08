// internal/config/config.go
package config

type Config struct {
    DatabasePath string
    ServerPort   string
    JWTSecret    string
}

func NewConfig() *Config {
    return &Config{
        DatabasePath: "./invoice.db",
        ServerPort:   "8080",
        JWTSecret:    "your-secret-key", // In production, this should be properly secured
    }
}