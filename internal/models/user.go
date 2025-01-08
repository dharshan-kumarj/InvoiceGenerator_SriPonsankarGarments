package models

import "time"

type User struct {
    ID           int64     `json:"id"`
    Email        string    `json:"email"`
    PasswordHash string    `json:"-"`     // The "-" means this field won't be included in JSON
    CreatedAt    time.Time `json:"created_at"`
}