// internal/handlers/invoice_handler.go
package handlers

import (
    "net/http"
    "html/template"
)

type InvoiceHandler struct {
    templates *template.Template
}

func NewInvoiceHandler() *InvoiceHandler {
    templates := template.Must(template.ParseFiles("internal/templates/invoice.html"))
    return &InvoiceHandler{
        templates: templates,
    }
}

func (h *InvoiceHandler) ShowInvoice(w http.ResponseWriter, r *http.Request) {
    // Later we'll add dynamic data here
    h.templates.ExecuteTemplate(w, "invoice.html", nil)
}

func (h *InvoiceHandler) GenerateInvoice(w http.ResponseWriter, r *http.Request) {
    // This will be implemented later for handling form submission
    // and generating dynamic invoices
}