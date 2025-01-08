// internal/auth/service.go
package auth

import (
    "database/sql"
    "errors"
    "time"
    "golang.org/x/crypto/bcrypt"
    "invoice-generator/internal/models" // Add this import
)

type AuthService struct {
    db *sql.DB
}

func NewAuthService(db *sql.DB) *AuthService {
    return &AuthService{db: db}
}

func (s *AuthService) Register(email, password string) error {
    // Check if user already exists
    var exists bool
    err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", email).Scan(&exists)
    if err != nil {
        return err
    }
    if exists {
        return errors.New("email already registered")
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    // Insert new user
    _, err = s.db.Exec(
        "INSERT INTO users (email, password_hash, created_at) VALUES (?, ?, ?)",
        email, string(hashedPassword), time.Now(),
    )
    return err
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
    user := &models.User{}  // Use models.User instead of User
    err := s.db.QueryRow(
        "SELECT id, email, password_hash FROM users WHERE email = ?",
        email,
    ).Scan(&user.ID, &user.Email, &user.PasswordHash)
    
    if err == sql.ErrNoRows {
        return nil, errors.New("invalid email or password")
    }
    if err != nil {
        return nil, err
    }

    // Check password
    err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
    if err != nil {
        return nil, errors.New("invalid email or password")
    }

    return user, nil
}