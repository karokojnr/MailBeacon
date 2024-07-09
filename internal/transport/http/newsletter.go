package http

import (
	"MailBeacon/internal/util"
	"net/http"
)

func (h *Handler) NewsletterSignup(w http.ResponseWriter, r *http.Request) {
	util.WriteJson(w, http.StatusOK, map[string]string{"message": "Thank you for signing up for our newsletter!"})
}
func (h *Handler) ConfirmNewsletterSignup(w http.ResponseWriter, r *http.Request) {
	util.WriteJson(w, http.StatusOK, map[string]string{"message": "Thank you for confirming your email address!"})
}
func (h *Handler) SendConfirmationEmail(w http.ResponseWriter, r *http.Request) {
	util.WriteJson(w, http.StatusOK, map[string]string{"message": "Confirmation email sent!"})
}
func (h *Handler) SendWelcomeEmail(w http.ResponseWriter, r *http.Request) {
	util.WriteJson(w, http.StatusOK, map[string]string{"message": "Welcome email sent!"})
}
