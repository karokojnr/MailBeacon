package http

import (
	"MailBeacon/internal/types"
	"MailBeacon/internal/utils"
	"encoding/base64"
	"encoding/json"
	"log"
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
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "ok!"})
}

func (h *Handler) SendConfirmationEmail(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Message struct {
			Data string `json:"data"`
		} `json:"message"`
	}

	if err := utils.ReadJson(r, &body); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	encodedJsonObject := body.Message.Data
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedJsonObject)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return

	}

	parsedPayload := types.SendConfirmationEmailRequest{}
	if err := json.Unmarshal(decodedBytes, &parsedPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	usr := types.ConvertSendConfirmationEmailRequestToUser(parsedPayload)

	err = h.Service.SendConfirmationEmail(r.Context(), usr)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "Confirmation email sent!"})
}

func (h *Handler) ConfirmNewsletterSignup(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	email := r.URL.Query().Get("email")
	if token == "" {
		msg := ""
		if email == "" {
			msg = "Invalid email"
		} else {
			msg = "Invalid token"
		}
		utils.WriteError(w, http.StatusBadRequest, msg)
		return
	}

	usr := types.User{
		Email: email,
		Token: token,
	}

	err := h.Service.ConfirmSubscription(r.Context(), usr)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return

	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "Thank you for confirming your email address!"})
}

func (h *Handler) SendWelcomeEmail(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Message struct {
			Data string `json:"data"`
		} `json:"message"`
	}

	if err := utils.ReadJson(r, &body); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	encodedJsonObject := body.Message.Data
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedJsonObject)
	if err != nil {
		log.Fatalf("Error decoding base64 string: %v", err)
	}

	parsedPayload := types.SendWelcomeEmailRequest{}
	if err := json.Unmarshal(decodedBytes, &parsedPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	usr := types.ConvertSendWelcomeEmailRequestToUser(parsedPayload)

	err = h.Service.SendWelcomeEmail(r.Context(), usr)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{"message": "Welcome email sent!"})
}
