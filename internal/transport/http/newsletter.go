package http

import (
	"MailBeacon/internal/types"
	"MailBeacon/internal/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func (h *Handler) NewsletterSignup(w http.ResponseWriter, r *http.Request) {
	var usrReq types.SignUpRequest

	if err := utils.ReadJson(r, &usrReq); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	validate := validator.New()
	if err := validate.Struct(usrReq); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	user := types.ConvertSignUpRequestToUser(usrReq)
	err := h.Service.SignUp(r.Context(), user)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Could not sign up user")
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "ok!"})
}
func (h *Handler) ConfirmNewsletterSignup(w http.ResponseWriter, r *http.Request) {
	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "Thank you for confirming your email address!"})
}
func (h *Handler) SendConfirmationEmail(w http.ResponseWriter, r *http.Request) {
	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "Confirmation email sent!"})
}
func (h *Handler) SendWelcomeEmail(w http.ResponseWriter, r *http.Request) {
	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "Welcome email sent!"})
}
