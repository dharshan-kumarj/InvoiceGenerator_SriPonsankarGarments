// internal/handlers/auth_handler.go
package handlers

import (
    "net/http"
    "html/template"
    "invoice-generator/internal/auth"     // Add this import
    "github.com/gorilla/sessions"         // Add this import
)

// Initialize the session store
var store = sessions.NewCookieStore([]byte("your-secret-key")) // Replace with a secure key in production

type AuthHandler struct {
    authService *auth.AuthService
    templates   *template.Template
}

func NewAuthHandler(authService *auth.AuthService) *AuthHandler {
    templates := template.Must(template.ParseFiles(
        "internal/templates/Login.html",
        "internal/templates/Register.html",
    ))
    return &AuthHandler{
        authService: authService,
        templates:   templates,
    }
}

func (h *AuthHandler) ShowLogin(w http.ResponseWriter, r *http.Request) {
    h.templates.ExecuteTemplate(w, "Login.html", nil)
}

func (h *AuthHandler) ShowRegister(w http.ResponseWriter, r *http.Request) {
    h.templates.ExecuteTemplate(w, "Register.html", nil)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    email := r.FormValue("email")
    password := r.FormValue("password")

    user, err := h.authService.Login(email, password)
    if err != nil {
        h.templates.ExecuteTemplate(w, "login.html", map[string]string{
            "Error": err.Error(),
        })
        return
    }

    // Create session
    session, _ := store.Get(r, "invoice-auth")
    session.Values["userID"] = user.ID
    session.Save(r, w)

    http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
    email := r.FormValue("email")
    password := r.FormValue("password")
    confirmPass := r.FormValue("confirm_password")

    if password != confirmPass {
        h.templates.ExecuteTemplate(w, "register.html", map[string]string{
            "Error": "Passwords do not match",
        })
        return
    }

    err := h.authService.Register(email, password)
    if err != nil {
        h.templates.ExecuteTemplate(w, "register.html", map[string]string{
            "Error": err.Error(),
        })
        return
    }

    http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
}

// Add logout handler
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "invoice-auth")
    session.Values = map[interface{}]interface{}{} // Clear all values
    session.Save(r, w)
    http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
}